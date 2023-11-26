package subscriptionscrapehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *ScrapeHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	// TODO scrape subscription from the respective chain
	// TODO ensure only current month is allowed
	// TODO forward valid flag in Task.Sync

	return nil
}
