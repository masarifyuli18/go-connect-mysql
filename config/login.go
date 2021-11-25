package config

type LoginService interface {
	LoginUser(username string, password string) bool
}
type loginInformation struct {
	username string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		username: "masarifyuli",
		password: "masarifyuli",
	}
}
func (info *loginInformation) LoginUser(username string, password string) bool {
	return info.username == username && info.password == password
}
