package subscriptionhandler

import (
	"github.com/xh3b4sd/tracer"
)

var updateSyncLockError = &tracer.Error{
	Kind: "updateSyncLockError",
	Desc: "The request expects a single background process to synchronize state at a time. A background process to synchronize state was found to be already ongoing. Therefore the request failed.",
}

var updateSyncInvalidError = &tracer.Error{
	Kind: "updateSyncInvalidError",
	Desc: `The request expects symbol.sync to be set to "dflt". symbol.sync was not found to be set to "dflt". Therefore the request failed.`,
}
