package service

import (
	"context"

	"github.com/khairulharu/bookingapi/domain"
	"github.com/khairulharu/bookingapi/dto"
)

type serviceBooking struct {
	chairRepository domain.ChairRepository
	userRepository  domain.UserRepository
}

func NewBooking(chairRepository domain.ChairRepository, userRepository domain.UserRepository) domain.BookingService {
	return &serviceBooking{
		chairRepository: chairRepository,
		userRepository:  userRepository,
	}
}

func (s serviceBooking) GetBookingChairs(ctx context.Context) dto.Response {
	chairs, err := s.chairRepository.GetChairs(ctx)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "error when get chairs",
			Error:   err.Error(),
		}
	}

	var resChairs []dto.ResChair

	for _, v := range chairs {
		resChairs = append(resChairs, dto.ResChair{
			Id:        v.Id,
			IsBook:    v.IsBook,
			UserBook:  v.UserBook,
			UserPhone: v.UserPhone,
			Pay:       v.Pay,
		})
	}

	return dto.Response{
		Code:    "00",
		Massage: "APPROVE",
		Data:    resChairs,
	}
}

func (s serviceBooking) SaveBookingChair(ctx context.Context, chair dto.ReqChair) dto.Response {
	valChair, err := s.chairRepository.GetChairByID(ctx, chair.Id)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "error chair not found",
			Error:   err.Error(),
		}
	}

	if valChair.IsBook == 1 {
		return dto.Response{
			Code:    "401",
			Massage: "chair is booking",
		}
	}

	saveChair := domain.Chair{
		Id:        chair.Id,
		IsBook:    1,
		UserBook:  chair.UserBook,
		UserPhone: chair.UserPhone,
		Pay:       1,
	}

	saveUser := domain.User{
		Name:  chair.UserBook,
		Phone: chair.UserPhone,
	}
	valUser, _ := s.userRepository.FindByPhone(ctx, saveUser.Phone)

	if valUser != (domain.User{}) {
		return dto.Response{
			Code:    "401",
			Massage: "user has been booking",
		}
	}
	_ = s.userRepository.Insert(ctx, &saveUser)

	err = s.chairRepository.UpdateChair(ctx, &saveChair)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "error",
			Error:   err.Error(),
		}
	}
	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
	}
}
