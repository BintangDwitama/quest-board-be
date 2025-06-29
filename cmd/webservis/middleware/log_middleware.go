package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m *middleware) Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		RequestTime := time.Now()
		UUID := uuid.New()

		ctx.Next()

		requestUUID := UUID
		requestOrigin := ctx.ClientIP()
		responseTime := time.Since(RequestTime)
		status := ctx.Writer.Status()

		log.Printf("requestID : %s origin : %s response time : %s status : %d\n", requestUUID, requestOrigin, responseTime, status)
	}
}
