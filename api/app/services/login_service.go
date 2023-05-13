package services

type LoginService interface {
	Login(email string, password string) bool
}

type loginService struct {
	authorizedEmail    string
	authorizedPassword string
}

var loginMock = loginService{authorizedEmail: "test1@gmail.com", authorizedPassword: "pass"}

func NewLoginService() LoginService {
	return &loginService{
		authorizedEmail:    loginMock.authorizedEmail,
		authorizedPassword: loginMock.authorizedPassword,
	}
}

func (service *loginService) Login(email string, password string) bool {
	// TODO: once connected to DB, search given email in DB and if found,
	//       compare stored passhash and received passhash (hashed in controller)
	//       using bcrypt
	return service.authorizedEmail == email && service.authorizedPassword == password
}
