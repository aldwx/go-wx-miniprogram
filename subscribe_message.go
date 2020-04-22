package weapp

import (
	"time"
)

const (
	apiSubscribeMessage = "/cgi-bin/message/subscribe/send"
)

// SubscribeMessage 订阅消息
type SubscribeMessage struct {
	ToUser     string               `json:"touser"`
	TemplateID string               `json:"template_id"`
	Page       string               `json:"page,omitempty"`
	Data       SubscribeMessageData `json:"data"`
}

// SubscribeMessageData 订阅消息模板数据
type SubscribeMessageData map[string]SubscribeMessageDataValue

type SubscribeMessageDataValue struct {
	Value string `json:"value"`
}

// Send 发送订阅消息
//
// token access_token
func (sm *SubscribeMessage) Send(token string) (*CommonError, error) {
	api := baseURL + apiSubscribeMessage
	return sm.send(api, token)
}

// Send 发送订阅消息 模拟测试
//
// token access_token
func (sm *SubscribeMessage) SendMock(token string) (*CommonError, error) {
	time.Sleep(50 * time.Millisecond)
	return &CommonError{ErrMSG: "ErrCodeCommonSuccess", ErrCode: ErrCodeCommonSuccess}, nil
}

func (sm *SubscribeMessage) send(api, token string) (*CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := &CommonError{}
	if err := postJSON(api, sm, res); err != nil {
		return nil, err
	}

	return res, nil
}
