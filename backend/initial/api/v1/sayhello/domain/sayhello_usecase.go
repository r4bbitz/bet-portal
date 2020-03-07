package domain

import (
	"github.com/r4bbitz/SayHello/api/v1/sayhello/data/repository"
	"github.com/r4bbitz/SayHello/api/v1/sayhello/model"
)

// PromotionUseCase is the interface
type SayHelloUseCase interface {
	Get(timeInput string) (model.Message, error)
}

// Encapsulated Implementation of UseCase Interface
type sayHelloUseCase struct {
	Repo repository.SayHelloRepository
}
func NewSayHelloUseCase(repo repository.SayHelloRepository) SayHelloUseCase {
	return &sayHelloUseCase{
		Repo: repo,
	}
}
func (uc *sayHelloUseCase) Get(timeInput string) (model.Message, error) {

	return uc.Repo.Get(timeInput)
}
