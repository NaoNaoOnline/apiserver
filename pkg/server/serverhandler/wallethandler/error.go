package wallethandler

import (
	"github.com/xh3b4sd/tracer"
)

var searchInternConflictError = &tracer.Error{
	Kind: "searchInternConflictError",
	Desc: "The request expects intern.wllt to be the only field provided within the given query object. Fields other than intern.wllt were found to be set within the given query object. Therefore the request failed.",
}

var searchPublicConflictError = &tracer.Error{
	Kind: "searchPublicConflictError",
	Desc: "The request expects public.kind to be the only field provided within the given query object. Fields other than public.kind were found to be set within the given query object. Therefore the request failed.",
}

var searchSymbolConflictError = &tracer.Error{
	Kind: "searchSymbolConflictError",
	Desc: "The request expects symbol.crtr to be the only field provided within the given query object. Fields other than symbol.crtr were found to be set within the given query object. Therefore the request failed.",
}

var searchSymbolInvalidError = &tracer.Error{
	Kind: "searchSymbolInvalidError",
	Desc: `The request expects symbol.crtr to be set to "default". symbol.crtr was not found to be set to "default". Therefore the request failed.`,
}

var updateEmptyError = &tracer.Error{
	Kind: "updateEmptyError",
	Desc: "The request expects the query object to contain all of [intern update]. The query object was not found to contain all of [intern update]. Therefore the request failed.",
}

var updateSymbolConflictError = &tracer.Error{
	Kind: "updateSymbolConflictError",
	Desc: "The request expects the query object to contain one of [symbol update]. The query object was not found to contain one of [symbol update]. Therefore the request failed.",
}

var walletLabelAccountingError = &tracer.Error{
	Kind: "walletLabelAccountingError",
	Desc: "The request expects the caller to designate only one accounting wallet. The caller was found to designate another accounting wallet. Therefore the request failed.",
}

var walletLabelAlreadyExistsError = &tracer.Error{
	Kind: "walletLabelAlreadyExistsError",
	Desc: "The request expects the provided wallet label not to exist already. The provided wallet label was found to already exist. Therefore the request failed.",
}

var walletLabelNotFoundError = &tracer.Error{
	Kind: "walletLabelNotFoundError",
	Desc: "The request expects the provided wallet label to exist already. The provided wallet label was not found to already exist. Therefore the request failed.",
}
