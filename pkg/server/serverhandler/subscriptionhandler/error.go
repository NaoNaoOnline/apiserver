package subscriptionhandler

import (
	"github.com/xh3b4sd/tracer"
)

var updateStatusSuccessError = &tracer.Error{
	Kind: "updateStatusSuccessError",
	Desc: "The request expects to verify subscription objects that are not yet verified. The subscription object was found to be already verified. Therefore the request failed.",
}

var updateSyncLockError = &tracer.Error{
	Kind: "updateSyncLockError",
	Desc: "The request expects a single background process to synchronize state at a time. A background process to synchronize state was found to be already ongoing. Therefore the request failed.",
}

var updateSyncInvalidError = &tracer.Error{
	Kind: "updateSyncInvalidError",
	Desc: `The request expects symbol.sync to be set to "default". symbol.sync was not found to be set to "default". Therefore the request failed.`,
}
