package labelstorage

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
// given label profile. Here ind describes the list of patches linked to the
// label object in question.
func (s PatchSlicer) AddPro(ind int, pro string) bool {
	for _, x := range s[ind] {
		if x.Ope == "add" && x.Pat == "/prfl/"+pro+"/data" {
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
// given label profile. Here ind describes the list of patches linked to the
// label object in question.
func (s PatchSlicer) RemPro(ind int, pro string) bool {
	for _, x := range s[ind] {
		if x.Ope == "remove" && x.Pat == "/prfl/"+pro+"/data" {
			return true
		}
	}

	return false
}
