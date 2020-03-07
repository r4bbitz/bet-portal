package handler

import (
	"github.com/labstack/echo"
	"github.com/r4bbitz/SayHello/api/v1/sayhello/domain"
	"net/http"
)

// PromotionHandler is the interface
type SayHelloHandler interface {
	Get(c echo.Context) error
}

type sayHelloHandler struct {
	UseCase domain.SayHelloUseCase
}
func NewSayHellonHandler(useCase domain.SayHelloUseCase) SayHelloHandler {

	return &sayHelloHandler{
		UseCase: useCase,
	}
}
func SayHelloMessageHandler(useCase domain.SayHelloUseCase) SayHelloHandler {

	return &sayHelloHandler{
		UseCase: useCase,
	}
}
func (h *sayHelloHandler) Get(c echo.Context) error {
	timeInput := c.QueryParam("timeInput")
	message, err := h.UseCase.Get(timeInput)
	if err == nil {
		return c.JSON(http.StatusOK, message)
	} else {
		return c.JSON(http.StatusNotFound, err)
	}

}
