package entity

//Login credential
type LoginCredentials struct {
	ID       uint64 `json:"id"`
	Username string `form:"username"`
	Password string `form:"password"`
}

var LoginUser = LoginCredentials{
	ID:       1,
	Username: "masarifyuli",
	Password: "masarifyuli",
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
