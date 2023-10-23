package policycache

import "github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"

func (m *Memory) create(rec *policystorage.Object) error {
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
				m.rec[rec.Syst] = map[string]*policystorage.Object{}
			}
		}

		{
			m.rec[rec.Syst][rec.Memb] = rec
		}
	}

	return nil
}
