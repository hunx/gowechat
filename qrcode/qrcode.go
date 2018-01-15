package qrcode

import (
	"encoding/json"
	"fmt"

	"github.com/keyunq/gowechat/context"
	"github.com/keyunq/gowechat/util"
)

const (
	qrcodeSendURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
)

//Qrcode 带参数二维码
type Qrcode struct {
	*context.Context
}

//NewQrcode 实例化
func NewQrcode(context *context.Context) *Qrcode {
	qr := new(Qrcode)
	qr.Context = context
	return qr
}

//整型参数
type IDScene struct {
	Scene_id int32 `json:"scene_id"`
}

//字符串参数
type StrSecne struct {
	Scene_str string `json:"scene_str"`
}

//临时IDqrcode 数据
type IDQrcodeTemp struct {
	ExpireSeconds int32  `json:"expire_seconds"` // 必须, 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）
	ActionName    string `json:"action_name"`    // 必须, 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值

	ActionInfo struct {
		Scene IDScene
	} `json:"action_info"` //必须，
}

//临时Strqrcode 数据
type StrQrcodeTemp struct {
	ExpireSeconds int32  `json:"expire_seconds"` // 必须, 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）
	ActionName    string `json:"action_name"`    // 必须, 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值

	ActionInfo struct {
		Scene StrSecne `json:"scene"`
	} `json:"action_info"` //必须，
}

//永久IDqrcode 数据
type IDQrcode struct {
	ActionName string `json:"action_name"` // 必须, 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值

	ActionInfo struct {
		Scene IDScene
	} `json:"action_info"` //必须，
}

//永久Strqrcode 数据
type StrQrcode struct {
	ActionName string `json:"action_name"` // 必须, 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	ActionInfo struct {
		Scene StrSecne `json:"scene"`
	} `json:"action_info"` //必须，
}

type resQrcodeSend struct {
	util.CommonError
	Ticket        string `json:"ticket"`
	ExpireSeconds int32  `json:"expire_seconds"`
	URL           string `json:"url"`
}

//Send 发送生成临时二维码请求
func (qrcode *Qrcode) Send(qrcodeTemp *StrQrcodeTemp) (ticket string, err error) {
	accessToken, err := qrcode.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", qrcodeSendURL, accessToken)
	response, err := util.PostJSON(uri, qrcodeTemp)

	var result resQrcodeSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = fmt.Errorf("create qrcode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}

	ticket = result.Ticket
	return
}

//SetQrData
func (qrcode *StrQrcodeTemp) SetQrData(expireSeconds int32, actionName string, actionStr string) {
	qrcode.ExpireSeconds = expireSeconds
	qrcode.ActionName = actionName
	qrcode.ActionInfo.Scene.Scene_str = actionStr
}

//Send 发送生成永久二维码请求
func (qrcode *Qrcode) SendForever(strQrcode *StrQrcode) (ticket string, err error) {
	accessToken, err := qrcode.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", qrcodeSendURL, accessToken)
	response, err := util.PostJSON(uri, strQrcode)

	var result resQrcodeSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = fmt.Errorf("create qrcode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}

	ticket = result.Ticket
	return
}

//SetForeverQrData
func (qrcode *StrQrcode) SetForeverQrData(actionName string, actionStr string) {
	qrcode.ActionName = actionName
	qrcode.ActionInfo.Scene.Scene_str = actionStr
}
