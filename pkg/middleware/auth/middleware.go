package auth

import (
	"fmt"
	"net/http"

	"github.com/NaoNaoOnline/apiserver/pkg/context/token"
	"github.com/golang-jwt/jwt/request"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type MiddlewareConfig struct {
	Log logger.Interface
}

type Middleware struct {
	log logger.Interface
}

func NewMiddleware(c MiddlewareConfig) *Middleware {
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}

	return &Middleware{
		log: c.Log,
	}
}

func (m *Middleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tok, err := request.AuthorizationHeaderExtractor.ExtractToken(r)
		if err != nil {
			tracer.Panic(err)
		}

		h.ServeHTTP(w, r.WithContext(token.NewContext(r.Context(), tok)))
	})
}
