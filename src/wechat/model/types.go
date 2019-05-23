package model

import (
	"encoding/xml"
)

type BaseMsg struct {
	XMLName xml.Name `xml:"xml"`
	MsgType string   `xml:"MsgType"`
}

type TextMsg struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   CDataString `xml:"ToUserName"`
	FromUserName CDataString `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      CDataString `xml:"MsgType"`
	Content      CDataString `xml:"Content"`
}

type ReceiveImageMsg struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   CDataString `xml:"ToUserName"`
	FromUserName CDataString `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      CDataString `xml:"MsgType"`
	PicUrl       CDataString `xml:"PicUrl"`
	MediaId      CDataString `xml:"MediaId"`
}

type ReplyImageMsg struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   CDataString `xml:"ToUserName"`
	FromUserName CDataString `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      CDataString `xml:"MsgType"`
	Image        Image       `xml:"Image"`
}

type Image struct {
	MediaId CDataString `xml:"MediaId"`
}

type CDataString struct {
	Value string `xml:",cdata"`
}
