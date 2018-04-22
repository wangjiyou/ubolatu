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
	SessionKey      string `json:"sessionKey"`
	Timestamp       string `json:"timestamp"`
}

type UserResponse struct {
	Data       string `json:"data"`
	StatusCode int    `json:"statusCode"`
}

type LoginRequest struct {
	Code string `json:"code"`
}

type FullUserInfo struct {
	UserInfo      STUserInfo
	RawData       string `json:"rawData"`
	Signature     string `json:"signature"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}

/*
type FullUserInfo struct {
	UserInfo struct {
		NickName  string `json:"nickName"`
		Gender    int    `json:"gender"`
		Language  string `json:"language"`
		City      string `json:"city"`
		Province  string `json:"province"`
		Country   string `json:"country"`
		AvatarURL string `json:"avatarUrl"`
	} `json:"userInfo"`
	RawData       string `json:"rawData"`
	Signature     string `json:"signature"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}
*/
type STUserInfo struct {
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Language  string `json:"language"`
}
