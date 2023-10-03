package budget

const (
	// Default is the default budget per Budget instance. Every Budget
	// instance does not pass through more object IDs than its configured limit
	// allows.
	Default = 100
)

type Budget struct {
	bud int
}

func New(bud ...int) *Budget {
	if len(bud) == 1 && bud[0] > 0 {
		return &Budget{bud: bud[0]}
	}

	return &Budget{bud: Default}
}

func (l *Budget) Break() bool {
	return l.bud < 0
}

func (l *Budget) Claim(num int) int {
	if l.bud <= 0 {
		l.bud = -1
		return 0
	}

	if num > l.bud {
		num = l.bud
		l.bud = -1
	} else {
		l.bud -= num
	}

	return num
}
