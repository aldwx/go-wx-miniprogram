package soter

import (
	httpclient "github.com/aldwx/go-http-client"
	"github.com/aldwx/go-wx-miniprogram/common"
)

const (
	apiVerifySignature = "/cgi-bin/soter/verify_signature"
)

type Soter struct{}

// VerifySignatureResponse 生物认证秘钥签名验证请求返回数据
type VerifySignatureResponse struct {
	common.CommonError
	IsOk bool `json:"is_ok"`
}

// VerifySignature 生物认证秘钥签名验证
// accessToken 接口调用凭证
// openID 用户 openid
// data 通过 wx.startSoterAuthentication 成功回调获得的 resultJSON 字段
// signature 通过 wx.startSoterAuthentication 成功回调获得的 resultJSONSignature 字段
func (s *Soter) VerifySignature(token, openID, data, signature string) (*VerifySignatureResponse, error) {
	api := common.BaseURL + apiVerifySignature
	return verifySignature(api, token, openID, data, signature)
}

func verifySignature(api, token, openID, data, signature string) (*VerifySignatureResponse, error) {
	url, err := httpclient.TokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := httpclient.RequestParams{
		"openid":         openID,
		"json_string":    data,
		"json_signature": signature,
	}

	res := new(VerifySignatureResponse)
	if err := httpclient.PostJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
