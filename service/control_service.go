package service

import (
	dto "github.com/DnullP/se_work/model/DTO"
)

type ControlService interface {
	HandleAdminCommand(dto.Command)
	HandleUserCommand(dto.Command)
}
