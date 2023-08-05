package labelstorage

type Object struct {
	// Name is the label name.
	Name string `json:"name"`
}

type Interface interface {
	Create() error
}
