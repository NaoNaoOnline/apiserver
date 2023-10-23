package objectlabel

const (
	ActionBuffer string = "buffer"
	ActionDelete string = "delete"
	ActionScrape string = "scrape"
	ActionUpdate string = "update"
)

const (
	OriginCustom string = "custom"
	OriginSystem string = "system"
)

const (
	DescAction string = "description.naonao.io/action"
	DescObject string = "description.naonao.io/object"
	DescOrigin string = "description.naonao.io/origin"
)

const (
	EvntAction string = "event.naonao.io/action"
	EvntObject string = "event.naonao.io/object"
	EvntOrigin string = "event.naonao.io/origin"
)

const (
	PlcyAction string = "policy.naonao.io/action"
	PlcyBuffer string = "policy.naonao.io/buffer"
	PlcyChanid string = "policy.naonao.io/chanid"
	PlcyCntrct string = "policy.naonao.io/cntrct"
	PlcyRpcUrl string = "policy.naonao.io/rpcurl"
	PlcyOrigin string = "policy.naonao.io/origin"
	PlcyUnique string = "policy.naonao.io/%06d"
)
