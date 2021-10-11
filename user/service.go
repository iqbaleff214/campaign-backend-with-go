package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(request RegisterUserRequest) (User, error)
	Login(request LoginRequest) (User, error)
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

func (s *service) Login(request LoginRequest) (User, error) {
	email := request.Email
	password := request.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("the email is not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
