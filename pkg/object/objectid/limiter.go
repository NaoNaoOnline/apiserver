package objectid

const (
	// Limit is the default limit per limiter instance. Any limiter instance does
	// not pass through more object IDs than its configured limit allows.
	Limit = 100
)

type Limiter struct {
	lim int
}

func NewLimiter(lim ...int) *Limiter {
	if len(lim) == 1 && lim[0] > 0 {
		return &Limiter{lim: lim[0]}
	}

	return &Limiter{lim: Limit}
}

func (l *Limiter) Limit(num int) int {
	if num > l.lim {
		num = l.lim
	}

	{
		l.lim -= num
	}

	return num
}
