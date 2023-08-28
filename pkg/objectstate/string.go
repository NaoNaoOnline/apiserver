package objectstate

const (
	Deleted String = "deleted"
	Updated String = "Updated"
)

type String string

func (s String) String() string {
	return string(s)
}
