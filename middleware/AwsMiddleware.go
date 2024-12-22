package middleware

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
	"net/http"
)

func AwsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg, err := config.LoadDefaultConfig(r.Context())
		if err != nil {
			log.Printf("Error loading AWS config: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "awsConfig", cfg)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
