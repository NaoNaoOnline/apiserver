package policyhandler

import (
	"github.com/xh3b4sd/tracer"
)

var searchLtstInvalidError = &tracer.Error{
	Kind: "updateSyncInvalidError",
	Desc: `The request expects symbol.ltst to be set to "default". symbol.ltst was not found to be set to "default". Therefore the request failed.`,
}

var searchSymbolEmptyError = &tracer.Error{
	Kind: "searchSymbolEmptyError",
	Desc: "The request expects symbol.ltst not to be empty. symbol.ltst was found to be empty. Therefore the request failed.",
}

var updateSyncInvalidError = &tracer.Error{
	Kind: "updateSyncInvalidError",
	Desc: `The request expects symbol.sync to be set to "default". symbol.sync was not found to be set to "default". Therefore the request failed.`,
}

var updateSyncLockError = &tracer.Error{
	Kind: "updateSyncLockError",
	Desc: `The request expects a single background process to synchronize state at a time. A background process to synchronize state was found to be already ongoing. Therefore the request failed.`,
}
