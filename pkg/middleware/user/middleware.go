package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/google/uuid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

type MiddlewareConfig struct {
	Log logger.Interface
	Red redigo.Interface
}

type Middleware struct {
	log logger.Interface
	red redigo.Interface
}

func NewMiddleware(c MiddlewareConfig) *Middleware {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	return &Middleware{
		log: c.Log,
		red: c.Red,
	}
}

func (m *Middleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		var ctx context.Context
		{
			ctx = r.Context()
		}

		// Lookup the OAuth subject claim if available. If no subject claim is
		// present we are dealing with an unauthenticated request and simply execute
		// the next handler, since data may be returned to anyone on the internet.
		var cla *validator.ValidatedClaims
		{
			cla, _ = ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
			if cla == nil || cla.RegisteredClaims.Subject == "" {
				h.ServeHTTP(w, r)
				return
			}
		}

		// At this point an OAuth subject claim was provided with the request and we
		// can trust that the originating user was successfully authenticated.
		var key string
		{
			key = fmt.Sprintf(keyfmt.SubjectClaim, cla.RegisteredClaims.Subject)
		}

		// Now search for our internal user mapping given the external subject
		// claim. Create the mapping if one does not yet exist for the current user.
		// Here we leverage redis over a remote connection and return an internal
		// server error if anything unexpected happens.
		var val string
		{
			val, err = m.red.Simple().Search().Value(key)
			if simple.IsNotFound(err) {
				{
					val = uuid.NewString()
				}

				{
					err = m.red.Simple().Create().Element(key, val)
					if err != nil {
						m.werror(ctx, w, err)
						return
					}
				}
			} else if err != nil {
				m.werror(ctx, w, err)
				return
			}
		}

		// Finally we looked up our internal user ID and add it to the request
		// context for further use.
		{
			r = r.Clone(userid.NewContext(ctx, val))
		}

		// Continue processing the request. The next handler may execute another
		// middleware or the RPC handler for the actual business logic.
		{
			h.ServeHTTP(w, r)
		}
	})
}

func (m *Middleware) werror(ctx context.Context, wri http.ResponseWriter, err error) {
	m.log.Log(
		ctx,
		"level", "error",
		"code", strconv.Itoa(http.StatusInternalServerError),
		"message", err.Error(),
	)

	{
		wri.WriteHeader(http.StatusInternalServerError)
		_, _ = wri.Write([]byte(`{"message":"` + err.Error() + `"}`))
	}
}
