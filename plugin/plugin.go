package plugin

import (
	httpclient "github.com/aldwx/go-http-client"
	"github.com/aldwx/go-wx-miniprogram/common"
)

const (
	apiPlugin    = "/wxa/plugin"
	apiDevPlugin = "/wxa/devplugin"
)

type Plugin struct{}

// ApplyPlugin 向插件开发者发起使用插件的申请
// accessToken 接口调用凭证
// action	string		是	此接口下填写 "apply"
// appID	string		是	插件 appId
// reason	string		否	申请使用理由
func (p *Plugin) ApplyPlugin(token, appID, reason string) (*common.CommonError, error) {
	api := common.BaseURL + apiPlugin
	return applyPlugin(api, token, appID, reason)
}

func applyPlugin(api, token, appID, reason string) (*common.CommonError, error) {
	url, err := httpclient.TokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := httpclient.RequestParams{
		"action":       "apply",
		"plugin_appid": appID,
	}

	if reason != "" {
		params["reason"] = reason
	}

	res := new(common.CommonError)
	if err := httpclient.PostJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPluginDevApplyListResponse 查询已添加的插件返回数据
type GetPluginDevApplyListResponse struct {
	common.CommonError
	ApplyList []struct {
		AppID      string `json:"appid"`      // 插件 appId
		Status     uint8  `json:"status"`     // 插件状态
		Nickname   string `json:"nickname"`   // 插件昵称
		HeadImgURL string `json:"headimgurl"` // 插件头像
		Categories []struct {
			First  string `json:"first"`
			Second string `json:"second"`
		} `json:"categories"` // 使用者的类目
		CreateTime string `json:"create_time"` // 使用者的申请时间
		ApplyURL   string `json:"apply_url"`   // 使用者的小程序码
		Reason     string `json:"reason"`      // 使用者的申请说明
	} `json:"apply_list"` // 申请或使用中的插件列表
}

// GetPluginDevApplyList 获取当前所有插件使用方
// accessToken 接口调用凭证
// page	number		是	要拉取第几页的数据
// num		是	每页的记录数
func (p *Plugin) GetPluginDevApplyList(token string, page, num uint) (*GetPluginDevApplyListResponse, error) {
	api := common.BaseURL + apiDevPlugin
	return getPluginDevApplyList(api, token, page, num)
}

func getPluginDevApplyList(api, token string, page, num uint) (*GetPluginDevApplyListResponse, error) {
	url, err := httpclient.TokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := httpclient.RequestParams{
		"num":    num,
		"page":   page,
		"action": "dev_apply_list",
	}

	res := new(GetPluginDevApplyListResponse)
	if err := httpclient.PostJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPluginListResponse 查询已添加的插件返回数据
type GetPluginListResponse struct {
	common.CommonError
	PluginList []struct {
		AppID      string `json:"appid"`      // 插件 appId
		Status     int8   `json:"status"`     // 插件状态
		Nickname   string `json:"nickname"`   // 插件昵称
		HeadImgURL string `json:"headimgurl"` // 插件头像
	} `json:"plugin_list"` // 申请或使用中的插件列表
}

// GetPluginList 查询已添加的插件
// accessToken 接口调用凭证
func (p *Plugin) GetPluginList(token string) (*GetPluginListResponse, error) {
	api := common.BaseURL + apiPlugin
	return getPluginList(api, token)
}

func getPluginList(api, token string) (*GetPluginListResponse, error) {
	url, err := httpclient.TokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := httpclient.RequestParams{
		"action": "list",
	}

	res := new(GetPluginListResponse)
	if err := httpclient.PostJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DevAction 修改操作
type DevAction string

// 所有修改操作
const (
	DevAgree  DevAction = "dev_agree"  // 同意申请
	DevRefuse DevAction = "dev_refuse" // 拒绝申请
	DevDelete DevAction = "dev_refuse" // 删除已拒绝的申请者
)

// SetDevPluginApplyStatus 修改插件使用申请的状态
// accessToken 接口调用凭证
// appID 使用者的 appid。同意申请时填写。
// reason 拒绝理由。拒绝申请时填写。
// action 修改操作
func (p *Plugin) SetDevPluginApplyStatus(token, appID, reason string, action DevAction) (*common.CommonError, error) {
	api := common.BaseURL + apiDevPlugin
	return setDevPluginApplyStatus(api, token, appID, reason, action)
}

func setDevPluginApplyStatus(api, token, appID, reason string, action DevAction) (*common.CommonError, error) {
	url, err := httpclient.TokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := httpclient.RequestParams{
		"action": action,
		"appid":  appID,
		"reason": reason,
	}

	res := new(common.CommonError)
	if err := httpclient.PostJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UnbindPlugin 查询已添加的插件
// accessToken 接口调用凭证
// appID 插件 appId
func (p *Plugin) UnbindPlugin(token, appID string) (*common.CommonError, error) {
	api := common.BaseURL + apiPlugin
	return unbindPlugin(api, token, appID)
}

func unbindPlugin(api, token, appID string) (*common.CommonError, error) {
	url, err := httpclient.TokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := httpclient.RequestParams{
		"action":       "unbind",
		"plugin_appid": appID,
	}

	res := new(common.CommonError)
	if err := httpclient.PostJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}