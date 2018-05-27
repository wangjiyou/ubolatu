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
	OpenID        string     `json:"openId"`
	RawData       string     `json:"rawData"`
	Signature     string     `json:"signature"`
	EncryptedData string     `json:"encryptedData"`
	Iv            string     `json:"iv"`
	UserInfo      STUserInfo `json:"userInfo"`
}

type STUserInfo struct {
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Language  string `json:"language"`
}

type FriendShipRequest struct {
	OwnerID    string `json:"ownerId"`
	FriendID   string `json:"friendId"`
	FriendName string `json:"friendName"`
	AddType    string `json:"addType"`
	CreateAt   string `json:"createAt"`
}
