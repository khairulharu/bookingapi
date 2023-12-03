package service

import (
	"context"

	"github.com/khairulharu/bookingapi/domain"
	"github.com/khairulharu/bookingapi/dto"
)

type serviceBooking struct {
	chairRepository domain.ChairRepository
	userRepository  domain.UserRepository
	userService     domain.UserService
}

func NewBooking(chairRepository domain.ChairRepository, userRepository domain.UserRepository, userService domain.UserService) domain.BookingService {
	return &serviceBooking{
		chairRepository: chairRepository,
		userRepository:  userRepository,
		userService:     userService,
	}
}

func (s serviceBooking) GetBookingChairs(ctx context.Context) dto.Response {
	chairs, err := s.chairRepository.GetChairs(ctx)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "INVALID",
			Error:   err.Error(),
		}
	}

	var resChairs []dto.ResChair

	for _, v := range chairs {
		resChairs = append(resChairs, dto.ResChair{
			Id:        v.Id,
			CodeRef:   v.Code,
			IsBook:    v.IsBook,
			UserBook:  v.UserBook,
			UserPhone: v.UserPhone,
			Pay:       v.Pay,
		})
	}

	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
		Data:    resChairs,
	}
}

func (s serviceBooking) SaveBookingChair(ctx context.Context, chair dto.ReqChair) dto.Response {
	valChair, err := s.chairRepository.GetChairByID(ctx, chair.Id)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "INVALID",
			Error:   "erorr chairs is not found",
		}
	}

	if valChair.IsBook == 1 {
		return dto.Response{
			Code:    "401",
			Massage: "INVALID",
			Error:   "chairs is booking",
		}
	}

	valChair.IsBook = 1
	valChair.UserBook = chair.UserBook
	valChair.UserPhone = chair.UserPhone
	valChair.Pay = 1

	var reqUser = dto.ReqUser{
		Name:  chair.UserBook,
		Phone: chair.UserPhone,
	}

	if res := s.userService.StoreUser(ctx, reqUser); res != (dto.Response{}) {
		return res
	}

	err = s.chairRepository.Update(ctx, &valChair)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "INVALID",
			Error:   err.Error(),
		}
	}
	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
	}
}
