package middleware

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// NewCorsMiddleware cria e retorna um novo middleware para lidar com CORS.
func NewCorsMiddleware() gin.HandlerFunc {
	origins := []string{
		"http://localhost:7788",
	}

	corsOptions := cors.Options{
		AllowedOrigins:   origins,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	return cors.New(corsOptions)
}
