package weapp

import (
	"github.com/aldwx/go-wx-miniprogram/auth"
	"github.com/aldwx/go-wx-miniprogram/customer_message"
	"github.com/aldwx/go-wx-miniprogram/data_analysis"
	"github.com/aldwx/go-wx-miniprogram/express"
	"github.com/aldwx/go-wx-miniprogram/img"
	"github.com/aldwx/go-wx-miniprogram/immediate_delivery"
	"github.com/aldwx/go-wx-miniprogram/nearby_poi"
	"github.com/aldwx/go-wx-miniprogram/ocr"
	"github.com/aldwx/go-wx-miniprogram/plugin"
	"github.com/aldwx/go-wx-miniprogram/qrcode"
	"github.com/aldwx/go-wx-miniprogram/sec_check"
	"github.com/aldwx/go-wx-miniprogram/soter"
	"github.com/aldwx/go-wx-miniprogram/subscribe_message"
	"github.com/aldwx/go-wx-miniprogram/uniform_message"
	"github.com/aldwx/go-wx-miniprogram/updatable_message"
)

var (
	Auth              = new(auth.Auth)                            // 用户登录授权
	DataAnalysis      = new(data_analysis.DataAnalysis)           // 数据分析
	CustomerMessage   = new(customer_message.CustomerMessage)     // 客服消息
	UniformMsgSender  = new(uniform_message.UniformMsgSender)     // 统一服务消息
	UpdatableMessage  = new(updatable_message.UpdatableMessage)   // 动态消息
	Plugin            = new(plugin.Plugin)                        // 插件管理
	NearbyPoi         = new(nearby_poi.NearbyPoi)                 // 附近小程序
	QRCode            = new(qrcode.QRCode)                        // 小程序码
	SecCheck          = new(sec_check.SecCheck)                   // 内容安全
	_                 = 0                                         // todo 广告
	Img               = new(img.Img)                              // 图像处理
	ImmediateDelivery = new(immediate_delivery.ImmediateDelivery) // 即时配送
	Express           = new(express.Express)                      // 物流助手
	Ocr               = new(ocr.Ocr)                              // Ocr
	_                 = 0                                         // todo 运维中心
	_                 = 0                                         // todo 小程序搜索
	_                 = 0                                         // todo 服务市场
	Soter             = new(soter.Soter)                          // 生物认证
	SubscribeMessage  = new(subscribe_message.SubscribeMessage)   // 订阅消息
)
