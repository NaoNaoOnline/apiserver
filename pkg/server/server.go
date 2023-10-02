package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/gorilla/mux"
	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Han are the server specific handlers implementing the actual business
	// logic.
	Han []handler.Interface
	// Int are the Twirp specific interceptors wrapping the endpoint handlers.
	Int []twirp.Interceptor
	// Lis is the main HTTP listener bound to some configured host and port.
	Lis net.Listener
	Log logger.Interface
	// Mid are the protocol specific transport layer middlewares executed before
	// any RPC handler.
	Mid []mux.MiddlewareFunc
}

type Server struct {
	lis net.Listener
	log logger.Interface
	srv *http.Server
}

func New(c Config) *Server {
	if len(c.Han) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Han must not be empty", c)))
	}
	if c.Lis == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lis must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	var rtr *mux.Router
	{
		rtr = mux.NewRouter()
	}

	{
		rtr.Use(c.Mid...)
	}

	for _, x := range c.Han {
		x.Attach(rtr, twirp.WithServerInterceptors(c.Int...), twirp.WithServerPathPrefix(""))
	}

	return &Server{
		lis: c.Lis,
		log: c.Log,
		srv: &http.Server{Handler: rtr},
	}
}

func (s *Server) Daemon() {
	{
		s.log.Log(
			context.Background(),
			"level", "info",
			"message", "server listening for calls",
			"addr", s.lis.Addr().String(),
		)
	}

	{
		err := s.srv.Serve(s.lis)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}
}
