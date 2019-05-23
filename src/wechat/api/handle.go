package api

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"wechat/utils"
	"net/http"
	"sort"
	"time"
	"wechat/model"
)

// Validator url is wechat send
func HandleCheckSignature(c *gin.Context) {
	signature, _ := c.GetQuery("signature")
	timeStamp, _ := c.GetQuery("timestamp")
	nonce, _ := c.GetQuery("nonce")
	echoStr, _ := c.GetQuery("echostr")

	tmpStrings := []string{utils.Conf.App.Token, timeStamp, nonce}
	sort.Strings(tmpStrings)
	tmpStr := tmpStrings[0] + tmpStrings[1] + tmpStrings[2]
	tmp := utils.StrToSha1(tmpStr)
	if tmp == signature {
		c.String(200, echoStr)
	} else {
		c.String(401, "Unauthorized")
	}
}

// Handle request
func HandleRequest(c *gin.Context) {
	var base model.BaseMsg
	contentType := c.Request.Header.Get("Content-Type")
	switch contentType {
	case "text/xml":
		body, _ := c.GetRawData()
		err := xml.Unmarshal(body, &base)
		if err != nil {
			c.XML(http.StatusBadGateway, gin.H{
				"status":  false,
				"message": "Unmarshal body err",
			})
			return
		}
		result, err := handleReply(base, body)
		if err != nil {
			c.XML(http.StatusBadGateway, gin.H{
				"status":  false,
				"message": "Unmarshal body err",
			})
			return
		}
		result = utils.ReplaceEscapeStr(result)
		c.String(http.StatusOK, string(result))
	}
}

// Handle reply
func handleReply(base model.BaseMsg, body []byte) ([]byte, error) {
	// text msg
	if base.MsgType == "text" {
		var xmlContent model.TextMsg
		err := xml.Unmarshal(body, &xmlContent)
		if err != nil {
			return nil, err
		}
		result := model.TextMsg{}
		result.ToUserName = model.CDataString{Value: xmlContent.FromUserName.Value}
		result.FromUserName = model.CDataString{Value: xmlContent.ToUserName.Value}
		result.CreateTime = time.Now().Unix()
		result.MsgType = model.CDataString{Value: xmlContent.MsgType.Value}
		result.Content = model.CDataString{Value: xmlContent.Content.Value}
		reply, _ := xml.Marshal(result)
		return reply, nil
	}

	// Image msg
	if base.MsgType == "Image" {
		var xmlContent model.ReceiveImageMsg
		err := xml.Unmarshal(body, &xmlContent)
		if err != nil {
			return nil, err
		}
		result := model.ReplyImageMsg{}
		result.ToUserName = model.CDataString{Value: xmlContent.FromUserName.Value}
		result.FromUserName = model.CDataString{Value: xmlContent.ToUserName.Value}
		result.CreateTime = time.Now().Unix()
		result.MsgType = model.CDataString{Value: xmlContent.MsgType.Value}
		result.Image = model.Image{MediaId: model.CDataString{Value: xmlContent.MediaId.Value}}
		reply, _ := xml.Marshal(result)
		return reply, nil
	}
	return nil, nil
}
