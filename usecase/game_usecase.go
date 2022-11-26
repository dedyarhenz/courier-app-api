package usecase

import (
	"final-project-backend/dto"
)

type GameUsecase interface {
	Play(request dto.GamePlayRequest) (*dto.GameResponse, error)
}
