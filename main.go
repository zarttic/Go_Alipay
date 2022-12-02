/**
 *@filename       main.go
 *@Description
 *@author          liyajun
 *@create          2022-12-02 19:08
 */

package main

import (
	"alipay/controller"
	"alipay/settings"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"time"
)

func main() {

	//初始化 读取配置
	err := settings.Init()
	if err != nil {
		xlog.Error(err)
	}
	//初始化支付  注入配置
	client, err := alipay.NewClient(settings.Conf.Appid, settings.Conf.AlipayPrivateKey, settings.Conf.IsProduction)
	//获取url进行支付
	//设置编码
	client.SetCharset(alipay.UTF8).SetSignType(alipay.RSA2).SetNotifyUrl(settings.Conf.NotifyURL).SetReturnUrl(settings.Conf.ReturnURL)
	//获取时间戳 作为订单编号
	ts := time.Now().UnixMilli()
	tradeNo := fmt.Sprintf("%d", ts)
	bm := make(gopay.BodyMap)
	//交易名称
	bm.Set("subject", "test")
	//总价
	bm.Set("total_amount", "6666")
	//订单号
	bm.Set("out_trade_no", tradeNo)
	//商品编码
	bm.Set("product_code", settings.Conf.ProductCode)

	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Errorf("payUrl[(%+v)]", payUrl)
	//支付成功 回调
	r := gin.Default()
	r.POST("/pay/alipay/notify", controller.AlipayNotify)
	r.GET("/pay/alipay/return", controller.AlipayReturn)
	r.Run(":" + settings.Conf.AppConfig.Port)

}
