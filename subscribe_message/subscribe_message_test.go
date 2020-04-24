package subscribe_message

import (
	"net/http"
	"net/http/httptest"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestSendSubscribeMessage2(t *testing.T) {
	message := SubscribeMessage{
		ToUser:     "o7wMG0cc43f87cpBvnnlJYMW0-qE",
		TemplateID: "I1JXpyEHpmyeJX6hd_TtG2YZEJHP9hAZDVmZKyWXePg",
		Data: SubscribeMessageData{
			"thing1":  SubscribeMessageDataValue{Value: "活动打卡"},
			"number3": SubscribeMessageDataValue{Value: "5"},
			"thing4":  SubscribeMessageDataValue{Value: "累计达到多少次会有更多奖励噢"},
		},
	}
	send, err := message.Send("29_YWpzBygjFlmQbADwDJgj2fJXjG0_ODUKPBnjEy8q-4_KFUOyVoRduQxpQW9kP1TDC_DJ_3mm6CB4r4lAzN9n3vitsBBWkVRSXwitf7zyxJu1PJHCgmJAYLwRtpw_xrdyJ_JtAhDmORNFldy8WJAfABAFSJ")
	t.Logf("send(%+v) err(%+v)", send, err)
}

func TestSendSubscribeMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSubscribeMessage {
			t.Fatalf("Except to path '%s',get '%s'", apiSubscribeMessage, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ToUser string `json:"touser"` // 用户 openid

			TemplateID string `json:"template_id"`
			Page       string `json:"page,omitempty"`
			Data       map[string]struct {
				Value string `json:"value"`
			} `json:"data"`
		}{}
		if err := jsoniter.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ToUser == "" {
			t.Fatal("param touser can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok"
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	sender := SubscribeMessage{
		ToUser:     "mock-open-id",
		TemplateID: "mock-template-id",
		Page:       "mock-page",
		Data: SubscribeMessageData{
			"mock01.DATA": {
				Value: "mock-value",
			},
		},
	}

	_, err := sender.send(ts.URL+apiSubscribeMessage, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscribeMessageSerialization(t *testing.T) {
	str := `{"keyword1":{"value":"哈根达斯情黏中秋月饼冰淇淋，陪你爱在浓浓秋意里。"},"keyword2":{"value":"9月15日"}}`
	var value SubscribeMessageData
	if err := jsoniter.UnmarshalFromString(str, &value); err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", value)
}
