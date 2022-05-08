package user

// User 微信用户
type User struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        uint8    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionId    string   `json:"unionid"`
}
