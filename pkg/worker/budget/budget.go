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

func (b *Budget) Break() bool {
	return b.bud < 0
}

func (b *Budget) Claim(num int) int {
	if b.bud <= 0 {
		b.bud = -1
		return 0
	}

	if num > b.bud {
		num = b.bud
		b.bud = -1
	} else {
		b.bud -= num
	}

	return num
}
