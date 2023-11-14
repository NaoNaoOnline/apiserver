package policystorage

type Slicer []*Object

func (s Slicer) Memb() []string {
	var mem []string

	for _, x := range s {
		mem = append(mem, x.Memb)
	}

	return mem
}
