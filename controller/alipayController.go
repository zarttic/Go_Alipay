/**
 *@filename       alipayNotify.go
 *@Description
 *@author          liyajun
 *@create          2022-12-02 20:05
 */

package controller

import (
	"alipay/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"net/http"
)

func AlipayReturn(ctx *gin.Context) {
	notifyReq, err := alipay.ParseNotifyToBodyMap(ctx.Request)
	if err != nil {
		xlog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误上",
		})
		return
	}
	//验签
	ok, err := alipay.VerifySign(settings.Conf.AlipayPublicKey, notifyReq)
	fmt.Println(settings.Conf.AlipayPublicKey)
	if err != nil {
		xlog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误下",
		})
		return
	}
	msg := ""
	if ok {
		msg = "验签成功"
	} else {
		msg = "验签失败"
	}
	//todo 开始自己的业务
	ctx.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func AlipayNotify(ctx *gin.Context) {
	fmt.Println("回调函数 notify")
	tradeStatus := ctx.PostForm("trade_status")
	//if already closed
	if tradeStatus == "TRADE_CLOSED" {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "trade already closed",
		})
	}
	// trade success we need to sign
	if tradeStatus == "TRADE_SUCCESS" {
		notifyReq, err := alipay.ParseNotifyToBodyMap(ctx.Request)
		if err != nil {
			xlog.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数错误上",
			})
			return
		}
		//验签
		ok, err := alipay.VerifySign(settings.Conf.AlipayPublicKey, notifyReq)
		fmt.Println(settings.Conf.AlipayPublicKey)
		if err != nil {
			xlog.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数错误下",
			})
			return
		}
		msg := ""
		if ok {
			msg = "success"
		} else {
			msg = "验签失败"
		}
		// 验签
		//todo 操作自己的业务
		ctx.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	}
}
