package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
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
			w.expire()
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			w.search()
			time.Sleep(5 * time.Second)
		}
	}()

	{
		select {}
	}
}

func (w *Worker) expire() {
	err := w.res.Expire()
	if err != nil {
		w.lerror(tracer.Mask(err))
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

func (w *Worker) search() {
	var err error

	var tas *task.Task
	{
		tas, err = w.res.Search()
		if engine.IsTaskNotFound(err) {
			return
		} else if err != nil {
			w.lerror(tracer.Mask(err))
		}
	}

	var bud *budget.Budget
	{
		bud = budget.New()
	}

	// We track the current and desired amount of handlers for the current task in
	// order to decide whether to delete the task after all handlers got invoked.
	// The desired amount of handlers that can process the current task are those
	// that return true when calling Handler.Filter. The desired amount of workers
	// are then those that do not return an error when calling Handler.Ensure.
	var cur int
	var des int

	for _, h := range w.han {
		if !h.Filter(tas) {
			continue
		}

		{
			des++
		}

		{
			err := h.Ensure(tas, bud)
			if err != nil {
				w.lerror(tracer.Mask(err))
			} else {
				// We have to account for the worker budget when processing a task.
				// Calling Handler.Ensure may use up the entire budget and it may break
				// through the budget or it may not. Breaking through the budget means
				// that there is still work left to do. And so not breaking the worker
				// budget tells us here that Handler.Ensure successfully resolved the
				// task from its own point of view, allowing us to count with it towards
				// the desired amount of handlers we that we track.
				if !bud.Break() {
					cur++
				}
			}
		}
	}

	// If the current and desired amount of handlers match, we can delete the
	// task, assuming that it got properly resolved.
	if cur != 0 && des != 0 && cur == des {
		err := w.res.Delete(tas)
		if err != nil {
			w.lerror(tracer.Mask(err))
		}
	}
}
