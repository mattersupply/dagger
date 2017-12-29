package dagger

import (
	"net/http"

	"github.com/rs/cors"
)

// CORSConfig allows to customize the allowed headers
type CORSConfig struct {
	AllowedHeaders []string
	Debug          bool
}

// CORS returns an adapter that wrapps the request with CORS Headers
func CORS(options cors.Options) Adapter {
	return func(h http.Handler) http.Handler {
		c := cors.New(options)

		return c.Handler(h)
	}
}
