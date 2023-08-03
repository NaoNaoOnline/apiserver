package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/user"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type MiddlewareConfig struct {
	Log logger.Interface
	Use user.Interface
}

type Middleware struct {
	log logger.Interface
	use user.Interface
}

func NewMiddleware(c MiddlewareConfig) *Middleware {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Use == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Use must not be empty", c)))
	}

	return &Middleware{
		log: c.Log,
		use: c.Use,
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

		var obj *user.Object
		{
			obj, err = m.use.Search(cla.RegisteredClaims.Subject, "")
			if user.IsNotFound(err) {
				h.ServeHTTP(w, r)
				return
			} else if err != nil {
				m.werror(ctx, w, err)
				return
			}
		}

		// Finally we looked up our internal user ID and add it to the request
		// context for further use.
		{
			r = r.Clone(userid.NewContext(ctx, obj.User))
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
