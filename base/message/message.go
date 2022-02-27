package message

import (
	"encoding/xml"
	"strconv"
	"time"
)

type Replier interface {
	BuildXml(fromUserName, toUserName string) []byte
	ContentType() string
}

// Message 消息体
type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

// Reply 回复消息体
type Reply struct {
	XMLName      xml.Name  `xml:"xml"`
	ToUserName   CDATAText `xml:"ToUserName"`
	FromUserName CDATAText `xml:"FromUserName"`
	CreateTime   string    `xml:"CreateTime"`
	MsgType      CDATAText `xml:"MsgType"`

	// 文本消息
	Content *CDATAText `xml:"Content,omitempty"`
	// 图片消息
	Image *Media `xml:"Image,omitempty"`
	// 音频消息
	Voice *Media `xml:"Voice,omitempty"`
	// 视频消息
	Video *Video `xml:"Video,omitempty"`
	// 音乐消息
	Music *Music `xml:"Music,omitempty"`
}

type Media struct {
	MediaId CDATAText `xml:"MediaId,omitempty"`
}

// value2CDATA 值转换为CDATA
func value2CDATA(value string) CDATAText {
	return CDATAText{"<![CDATA[" + value + "]]>"}
}

// ptrValue2CDATA 值转换为指针型CDATA
func ptrValue2CDATA(value string) *CDATAText {
	return &CDATAText{"<![CDATA[" + value + "]]>"}
}

func (reply *Reply) BuildXml(fromUserName, toUserName string) []byte {
	reply.FromUserName = value2CDATA(fromUserName)
	reply.ToUserName = value2CDATA(toUserName)
	reply.CreateTime = strconv.FormatInt(time.Now().Unix(), 10)
	b, _ := xml.MarshalIndent(reply, "", "")
	return b
}

func (reply *Reply) ContentType() string {
	return "text/xml"
}

// CDATAText 文本域
type CDATAText struct {
	Text string `xml:",innerxml"`
}

// NewText 文本消息
func NewText(content string) *Reply {
	return &Reply{
		MsgType: value2CDATA("text"),
		Content: ptrValue2CDATA(content),
	}
}

// NewImage 图片消息
func NewImage(mediaId string) *Reply {
	return &Reply{
		MsgType: value2CDATA("image"),
		Image:   &Media{MediaId: value2CDATA(mediaId)},
	}
}

func NewVoice(mediaId string) *Reply {
	return &Reply{
		MsgType: value2CDATA("voice"),
		Voice:   &Media{MediaId: value2CDATA(mediaId)},
	}
}

// Video 视频消息
type Video struct {
	MediaId     CDATAText `xml:"MediaId"`
	Title       CDATAText `xml:"Title"`
	Description CDATAText `xml:"Description"`
}

func NewVideo(mediaId, title, description string) *Reply {
	return &Reply{
		MsgType: value2CDATA("video"),
		Video: &Video{
			MediaId:     value2CDATA(mediaId),
			Title:       value2CDATA(title),
			Description: value2CDATA(description),
		},
	}
}

// Music 音乐消息
type Music struct {
	Title        CDATAText `xml:"Title"`
	Description  CDATAText `xml:"Description"`
	MusicUrl     CDATAText `xml:"MusicUrl"`
	HQMusicUrl   CDATAText `xml:"HQMusicUrl"`
	ThumbMediaId CDATAText `xml:"ThumbMediaId"`
}

func NewMusic(title, description, musicUrl, HQMusicUrl, thumbMediaId string) *Reply {
	return &Reply{
		MsgType: value2CDATA("music"),
		Music: &Music{
			Title:        value2CDATA(title),
			Description:  value2CDATA(description),
			MusicUrl:     value2CDATA(musicUrl),
			HQMusicUrl:   value2CDATA(HQMusicUrl),
			ThumbMediaId: value2CDATA(thumbMediaId),
		},
	}
}

// News 图文
type News struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   CDATAText     `xml:"ToUserName"`
	FromUserName CDATAText     `xml:"FromUserName"`
	CreateTime   string        `xml:"CreateTime"`
	MsgType      CDATAText     `xml:"MsgType"`
	ArticleCount string        `xml:"ArticleCount"`
	Articles     []NewsArticle `xml:"Articles"`
}

// NewsArticle 图文消息文章
type NewsArticle struct {
	Item struct {
		Title       CDATAText `xml:"Title"`
		Description CDATAText `xml:"Description"`
		PicUrl      CDATAText `xml:"PicUrl"`
		Url         CDATAText `xml:"Url"`
	} `xml:"item"`
}

func NewNews(articles []map[string]string) *News {
	var list []NewsArticle
	for _, article := range articles {
		art := NewsArticle{}
		for key, value := range article {
			switch key {
			case "title":
				art.Item.Title = value2CDATA(value)
				break
			case "description":
				art.Item.Description = value2CDATA(value)
				break
			case "pic_url":
				art.Item.PicUrl = value2CDATA(value)
				break
			case "url":
				art.Item.Url = value2CDATA(value)
				break
			}
		}
	}
	return &News{
		MsgType:      value2CDATA("news"),
		ArticleCount: strconv.Itoa(len(list)),
		Articles:     list,
	}
}

func (reply *News) BuildXml(fromUserName, toUserName string) []byte {
	reply.FromUserName = value2CDATA(fromUserName)
	reply.ToUserName = value2CDATA(toUserName)
	reply.CreateTime = strconv.FormatInt(time.Now().Unix(), 10)
	b, _ := xml.MarshalIndent(reply, "", "")
	return b
}

func (reply *News) ContentType() string {
	return "text/xml"
}
