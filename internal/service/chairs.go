package service

import (
	"context"

	"github.com/khairulharu/bookingapi/domain"
	"github.com/khairulharu/bookingapi/dto"
	"github.com/khairulharu/bookingapi/internal/util"
)

type serviceChair struct {
	chairRepository domain.ChairRepository
	userRepository  domain.UserRepository
}

func NewChair(chairRepository domain.ChairRepository, userRepository domain.UserRepository) domain.ChairService {
	return &serviceChair{
		chairRepository: chairRepository,
		userRepository:  userRepository,
	}
}

func (s serviceChair) StoreChairs(ctx context.Context, value int) dto.Response {
	var chair []domain.Chair

	for i := 0; i < value; i++ {
		chair = append(chair, domain.Chair{
			IsBook:    0,
			Code:      util.GenerateRandomString(6),
			UserBook:  "not_booking",
			UserPhone: "not_booking",
			Pay:       0,
		})
	}
	if err := s.chairRepository.Insert(ctx, &chair, value); err != nil {
		return dto.Response{
			Code:    "401",
			Massage: "Error when insert chairs data",
			Error:   err.Error(),
		}
	}
	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
	}
}

func (s serviceChair) DeleteChairs(ctx context.Context) dto.Response {
	chairs, err := s.chairRepository.GetChairs(ctx)
	if err != nil {
		return dto.Response{
			Code:    "401",
			Massage: "chair not found",
			Error:   err.Error(),
		}
	}
	if err := s.chairRepository.Delete(ctx, &chairs); err != nil {
		return dto.Response{
			Code:    "401",
			Massage: "error delete chairs",
			Error:   err.Error(),
		}
	}
	users, err := s.userRepository.GetUsers(ctx)
	if err != nil {
		return dto.Response{
			Code:    "401",
			Massage: "users not found",
			Error:   err.Error(),
		}
	}
	if err := s.userRepository.Delete(ctx, &users); err != nil {
		return dto.Response{
			Code:    "401",
			Massage: "error delete users",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
	}
}
