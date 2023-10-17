package policycache

func (m *Memory) SearchRcrd() []*Record {
	m.mut.Lock()
	defer m.mut.Unlock()
	return m.cac
}

func (m *Memory) searchRcrd(sys int64, mem string) *Record {
	{
		_, exi := m.rec[sys]
		if !exi {
			return nil
		}
	}

	var rec *Record
	{
		rec = m.rec[sys][mem]
	}

	if rec != nil {
		return rec
	}

	return nil
}
