package usermiddleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NaoNaoOnline/apiserver/pkg/server/context/isprem"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/subjectclaim"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type MiddlewareConfig struct {
	Log logger.Interface
	Use userstorage.Interface
}

type Middleware struct {
	log logger.Interface
	use userstorage.Interface
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
		var sub string
		{
			sub = subjectclaim.FromContext(ctx)
			if sub == "" {
				h.ServeHTTP(w, r)
				return
			}
		}

		var obj *userstorage.Object
		{
			obj, err = m.use.SearchSubj(sub)
			if userstorage.IsSubjectClaimMapping(err) {
				h.ServeHTTP(w, r)
				return
			} else if err != nil {
				m.werror(ctx, w, tracer.Mask(err))
				return
			}
		}

		// Add relevant user internals to the request context for further use.
		{
			ctx = userid.NewContext(ctx, obj.User)
			ctx = isprem.NewContext(ctx, obj.Prem)
		}

		{
			r = r.Clone(ctx)
		}

		// Continue processing the request. The next handler may execute another
		// middleware or the RPC handler for the actual business logic.
		{
			h.ServeHTTP(w, r)
		}
	})
}

func (m *Middleware) werror(ctx context.Context, wri http.ResponseWriter, err error) {
	e, o := err.(*tracer.Error)
	if o {
		m.log.Log(
			ctx,
			"level", "error",
			"message", e.Error(),
			"code", strconv.Itoa(http.StatusInternalServerError),
			"description", e.Desc,
			"docs", e.Docs,
			"kind", e.Kind,
			"stack", tracer.Stack(e),
		)
	} else {
		m.log.Log(
			ctx,
			"level", "error",
			"code", strconv.Itoa(http.StatusInternalServerError),
			"message", err.Error(),
			"stack", tracer.Stack(err),
		)
	}

	{
		wri.WriteHeader(http.StatusInternalServerError)
		_, _ = wri.Write([]byte(`{"message":"` + err.Error() + `"}`))
	}
}
