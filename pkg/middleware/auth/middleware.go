package auth

import (
	"fmt"
	"net/http"
	"strings"

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
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Middleware{
		log: c.Log,
	}
}

func (m *Middleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		var tok string
		{
			tok, err = request.AuthorizationHeaderExtractor.ExtractToken(r)
			if isNoToken(err) {
				tok = "" // no access token is valid within this middleware
			} else if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		h.ServeHTTP(w, r.WithContext(token.NewContext(r.Context(), tok)))
	})
}

func isNoToken(err error) bool {
	return strings.Contains(err.Error(), "no token")
}
