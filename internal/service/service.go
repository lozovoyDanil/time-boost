package service

//TODO: create a separated service for authorization and authentication
type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService() *Service {
	return &Service{
		Authorization: NewAuthService(),
	}
}
