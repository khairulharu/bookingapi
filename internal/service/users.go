package service

import (
	"context"

	"github.com/khairulharu/bookingapi/domain"
	"github.com/khairulharu/bookingapi/dto"
)

type serviceUser struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &serviceUser{
		userRepository: userRepository,
	}
}

func (s serviceUser) StoreUser(ctx context.Context, req dto.ReqUser) dto.Response {

	saveUser := domain.User{
		Name:  req.Name,
		Phone: req.Phone,
	}
	valUser, _ := s.userRepository.FindByPhone(ctx, saveUser.Phone)

	if valUser != (domain.User{}) {
		return dto.Response{
			Code:    "401",
			Massage: "user has been booking",
		}
	}
	_ = s.userRepository.Insert(ctx, &saveUser)

	return dto.Response{}
}
