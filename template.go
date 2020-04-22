package weapp

import "strconv"

const (
	apiGetPubTemplateTitleList    = "/wxaapi/newtmpl/getpubtemplatetitles"
	apiGetCategory                = "/wxaapi/newtmpl/getcategory"
	apiGetPubTemplateKeyWordsById = "/wxaapi/newtmpl/getpubtemplatekeywords"
	apiGetTemplateList            = "/wxaapi/newtmpl/gettemplate"
)

// VerifySignatureResponse 生物认证秘钥签名验证请求返回数据
type PubTemplateTitleListResponse struct {
	CommonError
	Count uint                `json:"count"` // 模版标题列表总数
	Data  []*PubTemplateTitle `json:"data"`  // 模板标题列表
}

type PubTemplateTitle struct {
	Tid        int32  `json:"tid"`        // 模版标题 id
	Title      string `json:"title"`      // 模版标题
	Typ        int32  `json:"type"`       // 模版类型，2 为一次性订阅，3 为长期订阅
	CategoryId string `json:"categoryId"` // 模版所属类目 id
}

// GetPubTemplateTitleList 获取帐号所属类目下的公共模板标题
// accessToken 接口调用凭证
func GetPubTemplateTitleList(token, ids string, start, limit uint) (*PubTemplateTitleListResponse, error) {
	api := baseURL + apiGetPubTemplateTitleList
	return getPubTemplateTitleList(api, token, ids, start, limit)
}

func getPubTemplateTitleList(api, token, ids string, start, limit uint) (*PubTemplateTitleListResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}
	url, err = encodeURL(url, requestQueries{
		"ids":   ids,
		"start": strconv.Itoa(int(start)),
		"limit": strconv.Itoa(int(limit)),
	})
	if err != nil {
		return nil, err
	}

	res := new(PubTemplateTitleListResponse)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

type Category struct {
	CommonError

	Data []*CategoryData `json:"data"`
}

type CategoryData struct {
	Id   int    `json:"id"`   // 类目id，查询公共库模版时需要
	Name string `json:"name"` // 类目的中文名
}

// GetCategory 获取帐号所属类目下的公共模板标题
// accessToken 接口调用凭证
func GetCategory(token string) (*Category, error) {
	api := baseURL + apiGetCategory
	return getCategory(api, token)
}

func getCategory(api, token string) (*Category, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(Category)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

type PubTemplateKeyword struct {
	CommonError

	Data []*PubTemplateKeywordData `json:"data"`
}

type PubTemplateKeywordData struct {
	CommonError

	Kid     uint   `json:"kid"`     // 关键词 id，选用模板时需要
	Name    string `json:"name"`    // 关键词内容
	Example string `json:"example"` // 关键词内容对应的示例
	Rule    string `json:"rule"`    // 参数类型

}

// GetPubTemplateKeyWordsById 获取帐号所属类目下的公共模板标题
// accessToken 接口调用凭证
func GetPubTemplateKeyWordsById(token string, templateId uint32) (*PubTemplateKeyword, error) {
	api := baseURL + apiGetPubTemplateKeyWordsById
	return getPubTemplateKeyWordsById(api, token, templateId)
}

func getPubTemplateKeyWordsById(api, token string, templateId uint32) (*PubTemplateKeyword, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}
	url, err = encodeURL(url, requestQueries{
		"tid": strconv.Itoa(int(templateId)),
	})
	if url, err = encodeURL(url, requestQueries{
		"tid": strconv.Itoa(int(templateId)),
	}); err != nil {
		return nil, err
	}

	res := new(PubTemplateKeyword)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

type Template struct {
	CommonError

	Data []*TemplateData `json:"data"`
}

type TemplateData struct {
	CommonError

	PriTmplId string `json:"priTmplId"` // 添加至帐号下的模板 id，发送小程序订阅消息时所需
	Title     string `json:"title"`     // 模版标题
	Content   string `json:"content"`   // 模版内容
	Example   string `json:"example"`   // 模板内容示例
	Typ       uint32 `json:"type"`      // 模板内容示例

}

// getTemplateList 获取当前帐号下的个人模板列表
// accessToken 接口调用凭证
func GetTemplateList(token string) (*Template, error) {
	api := baseURL + apiGetTemplateList
	return getTemplateList(api, token)
}

func getTemplateList(api, token string) (*Template, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(Template)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}
