package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/handler"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
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
	Res rescue.Interface
}

type Worker struct {
	han []handler.Interface
	log logger.Interface
	res rescue.Interface
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

	{
		w.create()
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

	go func() {
		for {
			w.ticker()
			time.Sleep(5 * time.Second)
		}
	}()

	{
		select {}
	}
}

func (w *Worker) create() {
	var err error

	// Ensure any task template the respective handlers require.
	for _, x := range w.han {
		var tas *task.Task
		{
			tas = x.Create()
		}

		if tas == nil {
			continue
		}

		var exi bool
		{
			exi, err = w.res.Exists(tas)
			if err != nil {
				w.lerror(tracer.Mask(err))
			}
		}

		if exi {
			continue
		}

		{
			err := w.res.Create(tas)
			if err != nil {
				w.lerror(tracer.Mask(err))
			}
		}
	}

	// Emit any task the system requires during the program's startup sequence.
	for _, x := range w.han {
		var tas *task.Task
		{
			tas = x.Create()
		}

		if tas == nil {
			continue
		}

		var cre bool
		for _, y := range w.filter() {
			if tas.Meta.Has(y) {
				cre = true
			}
		}

		if !cre {
			continue
		}

		// If we want to initially create a task during the program's startup
		// sequence, then we need to ensure that any Task.Cron definition is
		// removed, since we want to emit the task in its scheduled task form right
		// now, instead of creating a task template.
		{
			tas.Cron = nil
		}

		{
			err := w.res.Create(tas)
			if err != nil {
				w.lerror(tracer.Mask(err))
			}
		}
	}
}

func (w *Worker) expire() {
	err := w.res.Expire()
	if err != nil {
		w.lerror(tracer.Mask(err))
	}
}

func (w *Worker) filter() []map[string]string {
	return []map[string]string{
		// We want to emit every task for every chain to buffer policy records
		// during the program's startup sequence in order to ensure the internal
		// caching of policy records.
		{
			objectlabel.PlcyAction: objectlabel.ActionBuffer,
			objectlabel.PlcyOrigin: objectlabel.OriginSystem,
		},
	}
}

func (w *Worker) lerror(err error) {
	e, o := err.(*tracer.Error)
	if o {
		w.log.Log(
			context.Background(),
			"level", "error",
			"message", e.Error(),
			"description", e.Desc,
			"docs", e.Docs,
			"kind", e.Kind,
			"stack", tracer.Stack(e),
		)
	} else {
		w.log.Log(
			context.Background(),
			"level", "error",
			"message", err.Error(),
			"stack", tracer.Stack(err),
		)
	}
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

func (w *Worker) ticker() {
	err := w.res.Ticker()
	if err != nil {
		w.lerror(tracer.Mask(err))
	}
}
