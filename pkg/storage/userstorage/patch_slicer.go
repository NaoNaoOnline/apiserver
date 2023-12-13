package userstorage

import "strings"

type PatchSlicer [][]*Patch

// AddPat returns the JSON Patch paths that define operation add.
func (s PatchSlicer) AddPat(ind int) []string {
	var pat []string

	for _, x := range s[ind] {
		if x.Ope == "add" {
			pat = append(pat, x.Pat)
		}
	}

	return pat
}

// AddPro expresses whether any patch object within the indexed list adds the
// given user profile. Here ind describes the list of patches linked to the user
// object in question.
func (s PatchSlicer) AddPro(ind int, key string) bool {
	for _, x := range s[ind] {
		if x.Ope == "add" && x.Pat == "/prfl/data/"+key {
			return true
		}
	}

	return false
}

// AddPat returns the JSON Patch paths that define operation remove.
func (s PatchSlicer) RemPat(ind int) []string {
	var pat []string

	for _, x := range s[ind] {
		if x.Ope == "remove" {
			pat = append(pat, x.Pat)
		}
	}

	return pat
}

// RemPro expresses whether any patch object within the indexed list removes the
// given user profile. Here ind describes the list of patches linked to the user
// object in question.
func (s PatchSlicer) RemPro(ind int, key string) bool {
	for _, x := range s[ind] {
		if x.Ope == "remove" && x.Pat == "/prfl/data/"+key {
			return true
		}
	}

	return false
}

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

func (s PatchSlicer) RplPro(ind int) bool {
	for _, x := range s[ind] {
		if (x.Ope == "add" || x.Ope == "remove") && strings.HasPrefix(x.Pat, "/prfl/data/") {
			return true
		}
	}

	return false
}
