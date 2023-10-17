package policycache

func (m *Memory) ExistsAcce(sys int64, mem string, acc int64) bool {
	{
		m.mut.Lock()
		defer m.mut.Unlock()
	}

	var rec *Record
	{
		rec = m.searchRcrd(sys, mem)
	}

	return rec != nil && rec.Acce == acc
}

func (m *Memory) ExistsMemb(mem string) bool {
	{
		m.mut.Lock()
		defer m.mut.Unlock()
	}

	var exi bool
	{
		_, exi = m.mem[mem]
	}

	return exi
}

func (m *Memory) ExistsSyst(sys int64, mem string) bool {
	{
		m.mut.Lock()
		defer m.mut.Unlock()
	}

	var rec *Record
	{
		rec = m.searchRcrd(sys, mem)
	}

	return rec != nil
}
