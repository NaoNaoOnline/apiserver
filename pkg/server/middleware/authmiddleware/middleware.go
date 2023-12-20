package authmiddleware

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/server/context/subjectclaim"
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

	var err error

	var fnc func(ctx context.Context) (interface{}, error)
	{
		fnc = jwks.NewCachingProvider(musUrl(c.Iss), 5*time.Minute).KeyFunc
	}

	var val *validator.Validator
	{
		val, err = validator.New(fnc, validator.RS256, c.Iss, []string{c.Aud})
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var m *Middleware
	{
		m = &Middleware{
			log: c.Log,
		}
	}

	{
		m.jwt = jwtmiddleware.New(
			val.ValidateToken,
			jwtmiddleware.WithCredentialsOptional(true),
			jwtmiddleware.WithValidateOnOptions(false),
			jwtmiddleware.WithErrorHandler(m.errHan),
		)
	}

	return m
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

func (m *Middleware) errHan(wri http.ResponseWriter, req *http.Request, err error) {
	{
		m.log.Log(
			context.Background(),
			"level", "error",
			"message", err.Error(),
			"path", req.URL.String(),
			"stack", tracer.Stack(err),
		)
	}

	{
		err = twirp.WriteError(wri, err)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}
}

func musUrl(str string) *url.URL {
	u, e := url.Parse(str)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return u
}
