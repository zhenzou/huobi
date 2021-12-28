package spot

import (
	"encoding/json"
)

type WebSocketResponseBase struct {
	Status string `json:"status"`
	Ch     string `json:"ch"`
	Ts     int64  `json:"ts"`
}

type WebSocketV2ResponseBase struct {
	Action  string `json:"action"`
	Code    int32  `json:"code"`
	Ch      string `json:"ch"`
	Message string `json:"message"`
}

func (p *WebSocketV2ResponseBase) IsSuccess() bool {
	return p.Code == 200
}

func ParseWSV2Resp(message string) *WebSocketV2ResponseBase {
	result := &WebSocketV2ResponseBase{}
	err := json.Unmarshal([]byte(message), result)
	if err != nil {
		return nil
	}

	return result
}
