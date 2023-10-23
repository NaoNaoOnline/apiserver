package permission

import (
	"github.com/xh3b4sd/tracer"
)

var policyScrapeFailedError = &tracer.Error{
	Kind: "policyScrapeFailedError",
}

var taskBudgetLimitError = &tracer.Error{
	Kind: "taskBudgetLimitError",
}
