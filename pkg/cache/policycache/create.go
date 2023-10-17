package policycache

func (m *Memory) create(rec *Record) error {
	{
		m.cac = append(m.cac, rec)
	}

	{
		m.mem[rec.Memb] = struct{}{}
	}

	{
		{
			_, exi := m.rec[rec.Syst]
			if !exi {
				m.rec[rec.Syst] = map[string]*Record{}
			}
		}

		{
			m.rec[rec.Syst][rec.Memb] = rec
		}
	}

	return nil
}
