package weapp

import (
	"github.com/valyala/fasthttp"
	"testing"
)

func TestUtil_postJSONWithBody(t *testing.T) {
	params := make(map[string]interface{})
	params["b"] = "B"
	reader := new(fasthttp.Args)
	raw, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	reader.AppendBytes(raw)
	t.Logf("reader(%+v)", reader)
	reader.Add("a", "A")
	t.Logf("reader(%+v)", reader)
}

func TestUtil_fastHttpGet(t *testing.T) {
	t.Log(fasthttp.ParseByteRange([]byte(`{
			"touser": "o7wMG0cc43f87cpBvnnlJYMW0-qE",
			"template_id": "I1JXpyEHpmyeJX6hd_TtG9sQKa_C1-pNsYg3C1HVtEw",
			"data": {
				"thing1": {
					"value": "活动打卡"
				},
				"number3": {
					"value": "5"
				},
				"thing4": {
					"value": "累计达到多少次会有更多奖励噢"
				}
			}
		}`), 10000))
}

func TestUtil_fastHttpParseByteRange(t *testing.T) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI("https://api.weixin.qq.com/cgi-bin/message/subscribe/send")

	requestBody := []byte(`{data:{}}`)
	req.SetBody(requestBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源
	if err := fasthttp.Do(req, resp); err != nil {
		t.Fatal(err)
	}
	t.Log(fasthttp.ParseByteRange(resp.Body(), 0))
}
