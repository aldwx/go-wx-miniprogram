package subscribe_message

import (
	"time"

	httpclient "github.com/aldwx/go-http-client"
	"github.com/aldwx/go-wx-miniprogram/common"
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
func (sm *SubscribeMessage) Send(token string) (*common.CommonError, error) {
	api := common.BaseURL + apiSubscribeMessage
	return sm.send(api, token)
}

// Send 发送订阅消息 模拟测试
//
// token access_token
func (sm *SubscribeMessage) SendMock(token string) (*common.CommonError, error) {
	time.Sleep(50 * time.Millisecond)
	return &common.CommonError{ErrMSG: "ErrCodeCommonSuccess", ErrCode: common.ErrCodeCommonSuccess}, nil
}

func (sm *SubscribeMessage) send(api, token string) (*common.CommonError, error) {
	api, err := httpclient.TokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := &common.CommonError{}
	if err := httpclient.PostJSON(api, sm, res); err != nil {
		return nil, err
	}

	return res, nil
}
