package subscriptionscrapehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *ScrapeHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	// TODO scrape subscription from the respective chain
	// TODO create subscription object in redis for subscription found onchain

	return nil
}
