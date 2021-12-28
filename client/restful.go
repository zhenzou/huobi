package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/zhenzou/huobi/config"
)

func NewRestfulClient(config config.Config) *RestfulClient {
	return &RestfulClient{
		baseURL:   config.Host,
		accessKey: config.AccessKey,
		secretKey: config.SecretKey,
		cli: resty.New().
			SetBaseURL(buildBaseURL(config.Host)).
			SetPreRequestHook(sign(config.AccessKey, config.SecretKey)).
			SetDebug(config.Debug),
	}
}

func buildBaseURL(host string) string {
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https") {
		return host
	}
	return "https://" + host
}

func sign(accessKey, secretKey string) resty.PreRequestHook {
	const RFC3399WithoutZone = "2006-01-02T15:04:05"
	signer := NewSigner(secretKey)

	return func(client *resty.Client, request *http.Request) error {
		query := request.URL.Query()
		query.Add("AccessKeyId", accessKey)
		query.Add("SignatureMethod", "HmacSHA256")
		query.Add("SignatureVersion", "2")
		query.Add("Timestamp", time.Now().UTC().Format(RFC3399WithoutZone))

		signature := signer.Sign(request.Method, request.Host, request.URL.Path, query.Encode())

		query.Add("Signature", signature)
		request.URL.RawQuery = query.Encode()

		return nil
	}
}

type RawApiResponse struct {
	Data interface{} `json:"data"`
}

type defaultApiResponse struct {
	code int64
	msg  string
	Data interface{} `json:"data"`
}

func (r *defaultApiResponse) UnmarshalJSON(data []byte) error {
	a := &apiResponseForJSON{
		Data: r.Data,
	}
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}
	if a.Status != "" {
		if a.Status != "ok" {
			r.code = a.ErrCode
			r.msg = a.ErrMsg
		}
	} else if a.Code != 200 {
		r.code = a.Code
		r.msg = a.Message
	} else {
		r.Data = a.Data
	}
	return nil
}

func (r defaultApiResponse) Error() error {
	if r.msg != "" {
		return fmt.Errorf("huobi api error. code %d, msg %s", r.code, r.msg)
	}
	return nil
}

type apiResponseForJSON struct {
	// for futures
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	// for spot
	Code    int64  `json:"code"`
	Message string `json:"message"`
	// for v1
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type RestfulClient struct {
	baseURL   string
	accessKey string
	secretKey string
	cli       *resty.Client
}

func (cli *RestfulClient) Get(ctx context.Context, path string, params url.Values, body interface{}) error {

	resp := cli.buildResponseBody(body)

	_, err := cli.cli.NewRequest().
		ForceContentType("application/json").
		SetContext(ctx).
		SetQueryParamsFromValues(params).
		SetResult(resp).
		SetError(resp).
		Get(path)
	if err != nil {
		return err
	}
	if resp, ok := resp.(*defaultApiResponse); ok {
		return resp.Error()
	}
	return nil
}

func (cli *RestfulClient) Post(ctx context.Context, path string, req, body interface{}) error {
	resp := cli.buildResponseBody(body)
	r := cli.cli.NewRequest().
		ForceContentType("application/json").
		SetContext(ctx).
		SetResult(resp).
		SetError(resp)

	if form, ok := req.(url.Values); ok {
		r.SetFormDataFromValues(form)
	} else {
		r.SetBody(req)
	}
	_, err := r.Post(path)
	if err != nil {
		return err
	}
	if resp, ok := resp.(*defaultApiResponse); ok {
		return resp.Error()
	}
	return nil
}

func (cli *RestfulClient) buildResponseBody(body interface{}) interface{} {
	var resp interface{}
	if raw, ok := body.(RawApiResponse); ok {
		resp = raw.Data
	} else {
		resp = &defaultApiResponse{
			Data: body,
		}
	}
	return resp
}

type PageRequest struct {
	PageIndex int64  `json:"page_index,omitempty"`
	PageSize  int64  `json:"page_size,omitempty"`
	SortKey   string `json:"sort_key,omitempty"`
}

type PageResponse struct {
	TotalSize   int64 `json:"total_size"`
	TotalPage   int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
}
