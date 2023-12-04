package userstorage

type PatchSlicer [][]*Patch

// RplHom expresses whether any patch object within the indexed list defines the
// custom default view path to be replaced. Here ind describes the list of
// patches linked to the user object in question.
func (s PatchSlicer) RplHom(ind int) bool {
	for _, x := range s[ind] {
		if x.Ope == "replace" && x.Pat == "/home/data" {
			return true
		}
	}

	return false
}

// RplNam expresses whether any patch object within the indexed list defines the
// user name path to be replaced. Here ind describes the list of patches linked
// to the user object in question.
func (s PatchSlicer) RplNam(ind int) bool {
	for _, x := range s[ind] {
		if x.Ope == "replace" && x.Pat == "/name/data" {
			return true
		}
	}

	return false
}
