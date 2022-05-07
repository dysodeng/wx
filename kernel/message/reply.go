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

// Reply 消息回复体
type Reply struct {
	replier Replier
}

func NewReply(replier Replier) *Reply {
	return &Reply{
		replier: replier,
	}
}
func (reply *Reply) Replier() Replier {
	return reply.replier
}

// ReplyBody 回复消息体
type ReplyBody struct {
	XMLName      xml.Name  `xml:"xml"`
	ToUserName   CDATAText `xml:"ToUserName"`
	FromUserName CDATAText `xml:"FromUserName"`
	CreateTime   string    `xml:"CreateTime"`
	MsgType      CDATAText `xml:"MsgType"`

	Content *CDATAText `xml:"Content,omitempty"` // 文本消息
	Image   *Media     `xml:"Image,omitempty"`   // 图片消息
	Voice   *Media     `xml:"Voice,omitempty"`   // 音频消息
	Video   *Video     `xml:"Video,omitempty"`   // 视频消息
	Music   *Music     `xml:"Music,omitempty"`   // 音乐消息
}

func (reply *ReplyBody) BuildXml(fromUserName, toUserName string) []byte {
	reply.FromUserName = value2CDATA(fromUserName)
	reply.ToUserName = value2CDATA(toUserName)
	reply.CreateTime = strconv.FormatInt(time.Now().Unix(), 10)
	b, _ := xml.MarshalIndent(reply, "", "")
	return b
}

func (reply *ReplyBody) ContentType() string {
	return "text/xml"
}

// Media 媒体消息
type Media struct {
	MediaId CDATAText `xml:"MediaId,omitempty"`
}

// Music 音乐消息
type Music struct {
	Title        CDATAText `xml:"Title"`
	Description  CDATAText `xml:"Description"`
	MusicUrl     CDATAText `xml:"MusicUrl"`
	HQMusicUrl   CDATAText `xml:"HQMusicUrl"`
	ThumbMediaId CDATAText `xml:"ThumbMediaId"`
}

// Video 视频消息
type Video struct {
	MediaId     CDATAText `xml:"MediaId"`
	Title       CDATAText `xml:"Title"`
	Description CDATAText `xml:"Description"`
}

// News 图文消息
type News struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   CDATAText     `xml:"ToUserName"`
	FromUserName CDATAText     `xml:"FromUserName"`
	CreateTime   string        `xml:"CreateTime"`
	MsgType      CDATAText     `xml:"MsgType"`
	ArticleCount int           `xml:"ArticleCount"`
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

// NewText 文本消息
func NewText(content string) *ReplyBody {
	return &ReplyBody{
		MsgType: value2CDATA("text"),
		Content: ptrValue2CDATA(content),
	}
}

// NewImage 图片消息
func NewImage(mediaId string) *ReplyBody {
	return &ReplyBody{
		MsgType: value2CDATA("image"),
		Image:   &Media{MediaId: value2CDATA(mediaId)},
	}
}

// NewVoice 音频消息
func NewVoice(mediaId string) *ReplyBody {
	return &ReplyBody{
		MsgType: value2CDATA("voice"),
		Voice:   &Media{MediaId: value2CDATA(mediaId)},
	}
}

// NewVideo 视频消息
func NewVideo(mediaId, title, description string) *ReplyBody {
	return &ReplyBody{
		MsgType: value2CDATA("video"),
		Video: &Video{
			MediaId:     value2CDATA(mediaId),
			Title:       value2CDATA(title),
			Description: value2CDATA(description),
		},
	}
}

// NewMusic 音乐消息
func NewMusic(title, description, musicUrl, HQMusicUrl, thumbMediaId string) *ReplyBody {
	return &ReplyBody{
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

// NewNews 图文消息
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
		ArticleCount: len(list),
		Articles:     list,
	}
}
