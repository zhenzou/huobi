package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/zhenzou/huobi/common"
	"github.com/zhenzou/huobi/config"
)

// TODO 自定义 decoder，暂时默认是 json
func NewWebSocketClient(cfg config.Config, path string) *WebSocketClient {
	client := WebSocketClient{
		host:             cfg.Host,
		accessKey:        cfg.AccessKey,
		secretKey:        cfg.SecretKey,
		autoConnect:      true,
		path:             path,
		subscribedTopics: map[string]interface{}{},
		subscribers:      map[string]*subscriber{},
		requests:         map[string]*common.Future{},
		writeCh:          make(chan interface{}, 128),
		readCh:           make(chan []byte, 128),
		stopper:          common.NewStopper(context.Background()),
	}
	return &client
}

type WebSocketClient struct {
	host string
	path string

	accessKey string
	secretKey string

	writeCh chan interface{}
	readCh  chan []byte

	autoConnect      bool
	subscribedTopics map[string]interface{}
	subscribers      map[string]*subscriber
	requests         map[string]*common.Future

	lock    sync.RWMutex
	once    sync.Once
	stopper *common.Stopper
}

func (client *WebSocketClient) Open(ctx context.Context) error {
	var err error

	client.once.Do(func() {
		err = client.connect(ctx)
	})

	return err
}

func (client *WebSocketClient) Close(ctx context.Context) {
	client.stopper.Stop()

	client.lock.Lock()
	defer client.lock.Unlock()

	for topic, subscriber := range client.subscribers {
		subscriber.close()
		delete(client.subscribers, topic)
		delete(client.subscribedTopics, topic)
	}
}

func (client *WebSocketClient) Sub(ctx context.Context, topic string, ch interface{}) error {
	return client.sub(ctx, topic, WSSubRequest{Sub: topic, Id: common.UUID()}, ch)
}

func (client *WebSocketClient) SubWithOp(ctx context.Context, topic string, ch interface{}) error {
	return client.sub(ctx, topic, WSOpRequest{Op: "sub", Topic: topic, Cid: common.UUID()}, ch)
}

func (client *WebSocketClient) SubWithAction(ctx context.Context, topic string, ch interface{}) error {
	return client.sub(ctx, topic, WSActionRequest{Action: "sub", Ch: topic}, ch)
}

func (client *WebSocketClient) sub(ctx context.Context, topic string, payload, ch interface{}) error {

	err := client.Open(ctx)
	if err != nil {
		return err
	}

	receiver := reflect.ValueOf(ch)
	if receiver.Kind() != reflect.Chan {
		panic("chan type required")
	}

	topic = strings.ToLower(topic)

	// 这把锁很大，但是一般只是在启动的时候调用，影响应该不大
	client.lock.Lock()
	defer client.lock.Unlock()

	if subscriber, found := client.subscribers[topic]; found {
		log.Printf("topic %s already subscribed\n", topic)
		subscriber.addReceiver(receiver)
		return nil
	}

	future := common.NewFuture(payload)
	client.writeCh <- future

	err = future.Await(ctx)
	if err != nil {
		return err
	}
	log.Printf("request to subscribe topic %s\n", topic)

	// TODO 应该接收到响应以后才放到 topic 里面，但是请求体没有保存，得看怎么实现
	client.subscribedTopics[topic] = payload
	client.subscribers[topic] = newSubscriber(receiver, receiver.Type().Elem())

	return nil
}

// TODO unsubscribe的时候只删除对应的，而不是全部的 subscriber
func (client *WebSocketClient) Unsub(ctx context.Context, topic string) error {
	payload := WSUnsubRequest{Unsub: topic, Id: common.UUID()}
	return client.unsub(ctx, topic, payload)
}

func (client *WebSocketClient) UnsubWithOP(ctx context.Context, topic string) error {
	payload := WSOpRequest{Op: "unsub", Topic: topic, Cid: common.UUID()}
	return client.unsub(ctx, topic, payload)
}

func (client *WebSocketClient) unsub(ctx context.Context, topic string, payload interface{}) error {

	client.lock.Lock()
	defer client.lock.Unlock()

	if _, found := client.subscribers[topic]; !found {
		return nil
	}

	future := common.NewFuture(payload)
	client.writeCh <- future

	err := future.Await(ctx)
	if err != nil {
		return err
	}

	log.Printf("request to unsubscribe %s\n", topic)

	delete(client.subscribedTopics, topic)

	return nil
}

func (client *WebSocketClient) Req(ctx context.Context, topic string, from, to time.Time, ptr interface{}) error {
	err := client.Open(ctx)
	if err != nil {
		return err
	}

	future := common.NewFuture(ptr)

	id := common.UUID()
	client.addRequest(id, future)
	defer client.removeRequest(id)

	client.writeCh <- WSReqRequest{Req: topic, Id: id, From: timeToTimestamp(from), To: timeToTimestamp(to)}

	log.Printf("request to req %s with id %s\n", topic, id)
	return future.Await(ctx)
}

func (client *WebSocketClient) addRequest(id string, future *common.Future) {
	client.lock.Lock()
	defer client.lock.Unlock()
	client.requests[id] = future
}

func (client *WebSocketClient) removeRequest(id string) {
	client.lock.Lock()
	defer client.lock.Unlock()
	delete(client.requests, id)
}

func (client *WebSocketClient) closeConn(conn *websocket.Conn) {
	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Printf("close websocket connection error %s\n", err.Error())
		}
	}
}

func (client *WebSocketClient) reconnect() {
	for {
		// 错误，一直重试
		err := client.connect(context.Background())
		if err == nil {
			break
		}
	}
	client.resubscribe()
}

func (client *WebSocketClient) resubscribe() {
	log.Println("request to resubscribe all topic")

	client.lock.RLock()
	defer client.lock.RUnlock()

	for _, payload := range client.subscribedTopics {
		client.writeCh <- payload
	}
}

func (client *WebSocketClient) connect(ctx context.Context) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	urlStr := fmt.Sprintf("wss://%s%s", client.host, client.path)
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, urlStr, nil)
	if err != nil {
		log.Println("WebSocket connected error", err)
		return err
	}

	log.Println("websocket connected:", urlStr)

	conn.SetCloseHandler(client.onClose)

	go client.readLoop(conn)
	go client.writeLoop(conn)

	for _, path := range config.PrivatePaths {
		if path == client.path {
			return client.requestAuth(ctx, conn)
		}
	}

	return nil
}

func (client *WebSocketClient) onClose(code int, msg string) error {
	log.Printf("websocket closed by server. code:%d, msg:%s\n", code, msg)
	return nil
}

// 同步请求
func (client *WebSocketClient) requestAuth(ctx context.Context, conn *websocket.Conn) error {
	if conn == nil {
		panic(errors.New("websocket conn is null"))
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")

	params := url.Values{}
	params.Set("AccessKeyId", client.accessKey)
	params.Set("SignatureMethod", "HmacSHA256")
	params.Set("SignatureVersion", "2")
	params.Set("Timestamp", timestamp)

	sign := NewSigner(client.secretKey)
	signature := sign.Sign("GET", client.host, client.path, params.Encode())

	id := common.UUID()

	request := WSAuthRequest{
		Op:               "auth",
		Type:             "api",
		Cid:              id,
		AccessKeyId:      client.accessKey,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
		Timestamp:        timestamp,
		Signature:        signature,
	}
	log.Printf("request to auth")

	client.writeCh <- request

	resp := &WSAuthResponse{}
	future := common.NewFuture(resp)
	client.addRequest("auth", future)
	defer client.removeRequest("auth")

	err := future.Await(ctx)
	if err != nil {
		return err
	}
	if resp.ErrCode != 0 {
		return fmt.Errorf("auth error. Code:%d, Mesage:%s", resp.ErrCode, resp.ErrMsg)
	}
	return nil

}

func (client *WebSocketClient) readLoop(conn *websocket.Conn) {
	client.stopper.Begin()
	defer client.stopper.End()

	for {
		select {
		case <-client.stopper.Chan():
			client.closeConn(conn)
			log.Println("read loop stopped")
			return

		default:
			msgType, buf, err := conn.ReadMessage()
			if err != nil {
				log.Println("read message error", err)
				client.closeConn(conn)
				client.reconnect()
				return
			}

			var message []byte
			if msgType == websocket.BinaryMessage {
				message, err = common.Ungzip(buf)
				if err != nil {
					log.Println("decompress data error", err)
					continue
				}
			} else if msgType == websocket.TextMessage {
				message = buf
			}

			var jdata map[string]interface{}

			err = json.Unmarshal(message, &jdata)
			if err != nil {
				log.Println("msg to map[string]json.RawMessage error.", err)
				continue
			}

			if op, found := jdata["op"]; found {
				// order 相关处理
				switch op {
				case "ping":
					ts := jdata["ts"]
					client.writeCh <- WSOpRequest{Op: "pong", Ts: ts.(string)}
				case "pong":
					// pong 响应异常
					log.Println("pong error ", jdata["err-msg"].(string))

				case "close":
					log.Println("error occurs when authentication in server side.")

				case "error":
					log.Println("illegal op or internal error, but websocket is still connected.")

				case "auth":
					// 火币接口上说有 cid 返回，实际没有，这里固定使用auth作为 id
					client.handleReqResponse("auth", message)
				case "notify":
					topic := jdata["topic"].(string)
					client.handleSubResponse(topic, message, jdata)

				case "sub":
					topic := jdata["topic"]
					log.Printf("subscribed order %s \n", topic)

				case "unsub":
					topic := jdata["topic"].(string)
					log.Println("unsubscribed order topic: ", topic)
					delete(client.subscribers, topic)

				default:
					log.Println("websocket received unknown data", jdata)
				}

			} else {
				// market
				if ts, found := jdata["ping"]; found {
					ts := int64(ts.(float64))
					client.writeCh <- WSPongRequest{Pong: ts}
				} else if topic, found := jdata["subbed"]; found {
					// sub success reply
					log.Println("subscribed ", topic)

				} else if topic, found := jdata["unsubbed"]; found {
					// Unsub success reply
					log.Println("unsubscribed ", topic)
					delete(client.subscribers, topic.(string))

				} else if topic, found := jdata["ch"]; found {
					// market sub reply data
					client.handleSubResponse(topic.(string), message, jdata)
				} else if topic, found := jdata["rep"]; found {
					// market request reply data
					client.handleReqResponse(topic.(string), message)
				} else if code, found := jdata["err-code"]; found {
					// market request reply data
					msg := jdata["err-msg"]
					log.Printf("error %d %s \n", code, msg)
				} else {
					log.Println("websocket received unknown data ", jdata)
				}
			}
		}
	}
}

func (client *WebSocketClient) writeLoop(conn *websocket.Conn) {
	client.stopper.Begin()
	defer client.stopper.End()

	ch := client.writeCh

	for payload := range ch {
		select {
		case <-client.stopper.Chan():
			log.Println("write loop stopped")
			return

		default:
			switch raw := payload.(type) {
			case *common.Future:
				// future 写入失败，返回失败，不再重试
				err := conn.WriteJSON(raw.Data)
				if err != nil {
					log.Println("write error ", err)
					// 只关闭，不自动重新连接，出发 read loop 去 reconnect，防止无限重连
					raw.Reject(err)
					client.closeConn(conn)
					return
				}
				raw.Resolve()

			default:
				// 写入失败，重新开始连接，并且重试写入
				err := conn.WriteJSON(payload)
				if err != nil {
					log.Println("write error ", err)
					ch <- payload
					// 只关闭，不自动重新连接，触发 read loop 去 reconnect，防止无限重连
					client.closeConn(conn)
					return
				}
			}
		}

	}
}

func (client *WebSocketClient) handleSubResponse(topic string, data []byte, jdata map[string]interface{}) {

	var subscriber *subscriber = nil
	topic = strings.ToLower(topic)

	client.lock.RLock()

	if _, found := client.subscribers[topic]; found {
		subscriber = client.subscribers[topic]
	} else if topic == "accounts" || topic == "positions" {
		//contract_code := jdata["data"][0]["contract_code"], &contract_code)
		contractCode := jdata["data"]

		topic := fmt.Sprintf("%s.%s", topic, contractCode)
		if _, found := client.subscribers[topic]; found {
			subscriber = client.subscribers[topic]
		} else if _, found := client.subscribers[fmt.Sprintf("%s.*", topic)]; found {
			subscriber = client.subscribers[fmt.Sprintf("%s.*", topic)]
		}
	} else if topic[:7] == "orders." {
		subscriber, _ = client.subscribers["orders.*"]
	} else if topic[:12] == "matchorders." {
		subscriber, _ = client.subscribers["matchorders.*"]
	} else if topic[:14] == "trigger_order." {
		subscriber, _ = client.subscribers["trigger_order.*"]
	} else if topic[len(topic)-19:] == ".liquidation_orders" {
		subscriber, _ = client.subscribers["public.*.liquidation_orders"]
	}

	client.lock.RUnlock()

	if subscriber == nil {
		log.Println("no subscribe for topic ", topic)
		return
	}

	if err := subscriber.handle(topic, data); err != nil {
		log.Println("handle message error", err)
	}
}

func (client *WebSocketClient) handleReqResponse(id string, body []byte) {

	client.lock.RLock()
	future, ok := client.requests[id]
	client.lock.RUnlock()

	if !ok {
		log.Println("no future for req ", id)
		return
	}
	err := json.Unmarshal(body, future.Data)
	if err != nil {
		future.Reject(err)
	} else {
		future.Resolve()
	}
}

func newSubscriber(receiver reflect.Value, typ reflect.Type) *subscriber {
	return &subscriber{
		receivers: []reflect.Value{receiver},
		typ:       typ,
	}
}

type subscriber struct {
	receivers []reflect.Value
	typ       reflect.Type
}

func (s *subscriber) handle(topic string, data []byte) error {
	value := reflect.New(s.typ)
	err := json.Unmarshal(data, value.Interface())
	if err != nil {
		return err
	}
	for _, ch := range s.receivers {
		ch.Send(value.Elem())
	}
	return nil
}

func (s *subscriber) addReceiver(receiver reflect.Value) {
	s.receivers = append(s.receivers, receiver)
}

func (s *subscriber) close() {
	for _, receiver := range s.receivers {
		receiver.Close()
	}
}

func timeToTimestamp(time time.Time) int64 {
	if time.IsZero() {
		return 0
	}
	return time.Unix()
}

type WSActionRequest struct {
	Action string `json:"action"`
	Ch     string `json:"ch"`
}

type WSAuthRequest struct {
	Op               string `json:"op"`
	Type             string `json:"type"`
	Cid              string `json:"cid,omitempty"`
	AccessKeyId      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
	Ticket           string `json:"ticket,omitempty"`
}

type WSAuthResponse struct {
	Op      string `json:"op"`
	Type    string `json:"type"`
	Cid     string `json:"cid,omitempty"`
	ErrCode int64  `json:"err-code"`
	ErrMsg  string `json:"err-msg"`
	UserID  int64  `json:"user_id"`
}

type WSOpRequest struct {
	Op    string `json:"op"`
	Cid   string `json:"cid,omitempty"`
	Topic string `json:"topic"`
	Ts    string `json:"ts,omitempty"`
}

type WSReqRequest struct {
	Req  string `json:"req"`
	Id   string `json:"id"`
	From int64  `json:"from,omitempty"`
	To   int64  `json:"to,omitempty"`
}

type WSSubRequest struct {
	Sub      string `json:"sub"`
	Id       string `json:"id"`
	DataType string `json:"data_type,omitempty"`
}

type WSUnsubRequest struct {
	Unsub    string `json:"unsub"`
	Id       string `json:"id"`
	DataType string `json:"data_type,omitempty"`
}

type WSPongRequest struct {
	Pong int64 `json:"pong"`
}
