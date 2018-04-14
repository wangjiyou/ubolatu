package pub

type UserInfoRequest struct {
	OpenID          string `json:"openId"`
	NickName        string `json:"nickName"`
	Gender          string `json:"gender"`
	City            string `json:"city"`
	Province        string `json:"province"`
	Country         string `json:"country"`
	AvatarURL       string `json:"avatarUrl"`
	UnionID         string `json:"unionId"`
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	Timestamp       string `json:"timestamp"`
}

type UserResponse struct {
	Data       string `json:"data"`
	StatusCode int    `json:"statusCode"`
}

type LoginRequest struct {
	Code string `json:"code"`
}
