package di

import (
	"github.com/r4bbitz/SayHello/api/v1/sayhello/data/repository"
	"github.com/r4bbitz/SayHello/api/v1/sayhello/domain"
	"github.com/r4bbitz/SayHello/api/v1/sayhello/gateway/handler"
)


func ProvideSayHelloRepository() repository.SayHelloRepository {

	return repository.NewSayHelloRepository()
}

// ProvidePromotionUseCase is the Provide UseCase
func ProvideSayHelloUseCase() domain.SayHelloUseCase {

	return domain.NewSayHelloUseCase(ProvideSayHelloRepository())
}

// ProvidePromotionHandler is the Provide Handler
func ProvideSayHelloHandler() handler.SayHelloHandler {

	return handler.NewSayHellonHandler(ProvideSayHelloUseCase())
}
