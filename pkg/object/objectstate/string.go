package objectstate

const (
	Deleted String = "deleted"
	Updated String = "updated"
)

type String string

func (s String) String() string {
	return string(s)
}
