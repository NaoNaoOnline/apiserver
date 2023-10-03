package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/worker/handler"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue/engine"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Han are the worker specific handlers implementing the actual business
	// logic.
	Han []handler.Interface
	Log logger.Interface
	// Res is the rescue engine used to participate in the distributed task queue.
	Res engine.Interface
}

type Worker struct {
	han []handler.Interface
	log logger.Interface
	res engine.Interface
}

func New(c Config) *Worker {
	if len(c.Han) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Han must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	return &Worker{
		han: c.Han,
		log: c.Log,
		res: c.Res,
	}
}

func (w *Worker) Daemon() {
	{
		w.log.Log(
			context.Background(),
			"level", "info",
			"message", "worker searching for tasks",
			"addr", w.res.Listen(),
		)
	}

	go func() {
		for {
			{
				time.Sleep(5 * time.Second)
			}

			{
				err := w.res.Expire()
				if err != nil {
					w.lerror(tracer.Mask(err))
				}
			}
		}
	}()

	go func() {
		for {
			{
				time.Sleep(5 * time.Second)
			}

			var err error

			var tas *task.Task
			{
				tas, err = w.res.Search()
				if engine.IsTaskNotFound(err) {
					continue
				} else if err != nil {
					w.lerror(tracer.Mask(err))
				}
			}

			var cou int

			for _, h := range w.han {
				if !h.Filter(tas) {
					continue
				}

				{
					err := h.Ensure(tas)
					if err != nil {
						w.lerror(tracer.Mask(err))
					} else {
						cou++
					}
				}
			}

			if cou == len(w.han) {
				err := w.res.Delete(tas)
				if err != nil {
					w.lerror(tracer.Mask(err))
				}
			}
		}
	}()

	{
		select {}
	}
}
func (w *Worker) lerror(err error) {
	w.log.Log(
		context.Background(),
		"level", "error",
		"message", err.Error(),
		"stack", tracer.Stack(err),
	)
}
