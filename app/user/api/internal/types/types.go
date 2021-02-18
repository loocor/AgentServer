// Code generated by goctl. DO NOT EDIT.
package types

type Profile struct {
	Id         int64  `json:"id, optional"`
	Kind       int64  `json:"kind, options=0|1|2|3|4"`
	State      int64  `json:"state, options=0|1|2|3"`
	Role       string `json:"role, optional"`
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	Nickname   string `json:"nickname, optional"`
	Gender     string `json:"gender"`
	OpenId     string `json:"openId, optional"`
	From       string `json:"from"`
	Password   string `json:"password"`
	IdNumber   string `json:"idNumber"`
	Organize   string `json:"organize, optional"`
	Department string `json:"department, optional"`
	JobTitle   string `json:"jobTitle, optional"`
	Avatar     string `json:"avatar, optional"`
	Address    string `json:"address, optional"`
	CreateTime int64  `json:"createTime, optional"`
	UpdateTime int64  `json:"updateTime, optional"`
	DeleteTime int64  `json:"deleteTime, optional"`
}

type AuthToken struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegReq struct {
	Captcha string  `json:"captcha"`
	Profile Profile `json:"profile"`
}

type RegResp struct {
	Code    int32     `json:"code"`
	Message string    `json:"message"`
	Token   AuthToken `json:"token"`
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

type LoginResp struct {
	Code    int32     `json:"code"`
	Message string    `json:"message"`
	Token   AuthToken `json:"token"`
}

type LogoutReq struct {
	Phone string `json:"phone"`
}

type LogoutResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type ProfileViewReq struct {
	Phone string `json:"phone"`
}

type ProfileViewResp struct {
	Code    int32   `json:"code"`
	Message string  `json:"message"`
	Profile Profile `json:"profile"`
}

type ProfileUpdateReq struct {
	Profile Profile `json:"profile"`
}

type ProfileUpdateResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
