package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/BintangDwitama/quest-board-be/internal/domain/usecase"
)

type middleware struct {
	usecase       usecase.Usecase
	middlewareCfg middlewareConfig
}

type middlewareConfig struct {
	jwtSecret   string
	env         string
	serviceName string
}

type Middleware interface {
	Logger() gin.HandlerFunc
}

func NewMiddleware(u usecase.Usecase) Middleware {
	return &middleware{}
}

type MiddlewareFactory interface {
	Use(...gin.HandlerFunc)
	Middlewares() []gin.HandlerFunc
}

type middlewareFactory struct {
	middlewares []gin.HandlerFunc
}

func NewMiddlewareFactory() MiddlewareFactory {
	return &middlewareFactory{}
}

func (factory *middlewareFactory) Use(middlewares ...gin.HandlerFunc) {
	factory.middlewares = append(factory.middlewares, middlewares...)
}

func (factory *middlewareFactory) Middlewares() []gin.HandlerFunc {
	return factory.middlewares
}
