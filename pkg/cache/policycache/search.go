package policycache

import "github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"

func (m *Memory) SearchRcrd() []*policystorage.Object {
	m.mut.Lock()
	defer m.mut.Unlock()
	return m.cac
}

func (m *Memory) searchRcrd(sys int64, mem string) *policystorage.Object {
	{
		_, exi := m.rec[sys]
		if !exi {
			return nil
		}
	}

	var rec *policystorage.Object
	{
		rec = m.rec[sys][mem]
	}

	if rec != nil {
		return rec
	}

	return nil
}
