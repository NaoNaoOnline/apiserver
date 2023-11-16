package walletstorage

import "strings"

type PatchSlicer [][]*Patch

// AddLab expresses whether any patch object within the indexed list adds the
// given wallet label. Here ind describes the list of patches linked to the
// wallet object in question.
func (s PatchSlicer) AddLab(ind int, lab string) bool {
	for _, x := range s[ind] {
		if x.Ope == "add" && strings.HasPrefix(x.Pat, "/labl/data") && (x.Val == lab || lab == "*") {
			return true
		}
	}

	return false
}

// RemLab expresses whether any patch object within the indexed list removes the
// given wallet label. Here ind describes the list of patches linked to the
// wallet object in question.
func (s PatchSlicer) RemLab(ind int, lab string) bool {
	for _, x := range s[ind] {
		if x.Ope == "remove" && strings.HasPrefix(x.Pat, "/labl/data") && (x.Val == lab || lab == "*") {
			return true
		}
	}

	return false
}
