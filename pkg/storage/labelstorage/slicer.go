package labelstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

// Cate returns all underlying label objects with kind cate.
func (s Slicer) Cate() Slicer {
	var obj []*Object

	for _, x := range s {
		if x.Kind == "cate" {
			obj = append(obj, x)
		}
	}

	return obj
}

// Host returns all underlying label objects with kind host.
func (s Slicer) Host() Slicer {
	var obj []*Object

	for _, x := range s {
		if x.Kind == "host" {
			obj = append(obj, x)
		}
	}

	return obj
}

// Labl returns all the label IDs for the underling list of label objects.
func (s Slicer) Labl() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Labl)
	}

	return ids
}

// Cate returns all underlying label object names.
func (s Slicer) Name() []string {
	var nam []string

	for _, x := range s {
		nam = append(nam, x.Name.Data)
	}

	return nam
}

// Prfl returns all underlying profile names with the given profile key.
func (s Slicer) Prfl(key string) []string {
	var nam []string

	for _, x := range s {
		// Since we are looking for e.g. Twitter handles, we only want to process
		// host labels.
		if x.Kind != "host" {
			continue
		}

		// If there is no profile data we record the label name instead.
		if x.Prfl.Data == nil {
			nam = append(nam, x.Name.Data)
			continue
		}

		var val string
		{
			val = x.Prfl.Data[key]
		}

		if val != "" {
			nam = append(nam, val)
		} else {
			// If there is no profile value we record the label name instead.
			nam = append(nam, x.Name.Data)
		}
	}

	return nam
}
