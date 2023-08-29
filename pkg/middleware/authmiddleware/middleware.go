package authmiddleware

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/context/subjectclaim"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type MiddlewareConfig struct {
	Aud string
	Iss string
	Log logger.Interface
}

type Middleware struct {
	log logger.Interface
	jwt *jwtmiddleware.JWTMiddleware
}

func NewMiddleware(c MiddlewareConfig) *Middleware {
	if c.Aud == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Aud must not be empty", c)))
	}
	if c.Iss == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Iss must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	val, err := validator.New(
		jwks.NewCachingProvider(musUrl(c.Iss), 5*time.Minute).KeyFunc,
		validator.RS256,
		c.Iss,
		[]string{c.Aud},
	)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return &Middleware{
		log: c.Log,
		jwt: jwtmiddleware.New(
			val.ValidateToken,
			jwtmiddleware.WithCredentialsOptional(true),
			jwtmiddleware.WithValidateOnOptions(false),
			jwtmiddleware.WithErrorHandler(errHan),
		),
	}
}

func (m *Middleware) Handler(h http.Handler) http.Handler {
	// CheckJWT extracts and validates the bearer access token provided with the
	// request's authorization header, if any. Any valid claims are put into the
	// request's context and can be accessed like shown below.
	return m.jwt.CheckJWT(http.HandlerFunc(func(wri http.ResponseWriter, req *http.Request) {
		var ctx context.Context
		{
			ctx = req.Context()
		}

		{
			cla, _ := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
			if cla != nil {
				req = req.Clone(subjectclaim.NewContext(ctx, cla.RegisteredClaims.Subject))
			}
		}

		{
			h.ServeHTTP(wri, req)
		}
	}))
}

func errHan(wri http.ResponseWriter, req *http.Request, err error) {
	twirp.WriteError(wri, err)
}

func musUrl(str string) *url.URL {
	u, e := url.Parse(str)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return u
}
