package objectstate

const (
	Deleted String = "deleted"
	Started String = "started"
	Updated String = "updated"
)

type String string

func (s String) String() string {
	return string(s)
}
