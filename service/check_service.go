package service

import dto "github.com/DnullP/se_work/model/DTO"

type CheckService interface {
	CheckIn(dto.CheckInfo)
	CheckOut(dto.CheckInfo)
}
