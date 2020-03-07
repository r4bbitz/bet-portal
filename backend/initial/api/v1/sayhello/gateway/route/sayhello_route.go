package route

import (
	"github.com/labstack/echo"
	"github.com/r4bbitz/SayHello/api/v1/sayhello/gateway/handler"
)

// PromotionRoute is the interface
type SayHelloRoute interface {
	Initial(e *echo.Echo)
}

type sayHelloRoute struct {
	Handle handler.SayHelloHandler
}
func NewSayHelloRoute(handler handler.SayHelloHandler) SayHelloRoute {
	return &sayHelloRoute{
		Handle: handler,
	}
}

func (r *sayHelloRoute) Initial(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("/sayHello/", r.Handle.Get)
}

