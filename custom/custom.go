package custom

import (
	"encoding/json"
	"fmt"

	"github.com/keyunq/gowechat/context"
	"github.com/keyunq/gowechat/util"
)

const (
	customSendURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

//Custom 客服消息
type Custom struct {
	*context.Context
}

//NewCustom 实例化
func NewCustom(context *context.Context) *Custom {
	ct := new(Custom)
	ct.Context = context
	return ct
}

//图片消息
type CustomImageData struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

type resCustomSend struct {
	util.CommonError
}

//Send 发送客服消息
func (custom *Custom) Send(customImage *CustomImageData) (err error) {
	accessToken, err := custom.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customSendURL, accessToken)
	response, err := util.PostJSON(uri, customImage)

	var result resCustomSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = fmt.Errorf("create qrcode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//SetCustomImageData
func (customImage *CustomImageData) SetCustomImageData(touser string, msgtype string, media_id string) {
	customImage.ToUser = touser
	customImage.MsgType = msgtype
	customImage.Image.MediaId = media_id
}
