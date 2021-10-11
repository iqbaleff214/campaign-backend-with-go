package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(request RegisterUserRequest) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterUser(r RegisterUserRequest) (User, error)  {
	user := User{}
	password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user = User{
		Name: r.Name,
		Email: r.Email,
		Occupation: r.Occupation,
		Password: string(password),
		Role: "user",
	}

	return s.repository.Store(user)
}
