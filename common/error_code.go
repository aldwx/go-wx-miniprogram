package common

// 全局返回码
// API https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Global_Return_Code.html
const (
	ErrCodeCommonSystemIsBusy      = -1    // 系统繁忙，此时请开发者稍候再试
	ErrCodeCommonSuccess           = 0     // 请求成功
	ErrCodeCommonAPIfReqOutOfLimit = 45009 // api频率超出限制
	ErrCodeCommonInvalidCredential = 40001 // 无效的凭证类型
	ErrCodeCommonIllegalCredential = 40002 // 不合法的凭证类型
	ErrCodeCommonTokenInvalid      = 4     // 自定义异常。Token无效-客户提供token获取方式有问题
)

// 订阅消息
// 发送订阅消息
// API  https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
// POST https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN
const (
	ErrCodeSubMsgOpenIdIsEmpty        = 40003 // touser字段openid为空或者不正确
	ErrCodeSubMsgTemplateIdIsEmpty    = 40037 // 订阅模板id为空不正确
	ErrCodeSubMsgUserCancelled        = 43101 // 用户拒绝接受消息，如果用户之前曾经订阅过，则表示用户取消了订阅关系
	ErrCodeSubMsgTemplateParamIsEmpty = 47003 // 模板参数不准确，可能为空或者不满足规则，errmsg会提示具体是哪个字段出错
	ErrCodeSubMsgPagePathIncorrect    = 41030 // page路径不正确，需要保证在现网版本小程序中存在，与app.json保持一致
	ErrCodeSubMsgUserRefuseToAccept   = 43101 // 用户拒绝接受消息
)
