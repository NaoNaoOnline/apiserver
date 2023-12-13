package liststorage

import (
	"strconv"
	"time"
)

const (
	Layout = "2006-01-02T15:04:05.999999Z"
)

type PatchSlicer [][]*Patch

// RplDes expresses whether any patch object within the indexed list defines the
// list description path to be replaced. Here ind describes the list of patches
// linked to the list object in question.
func (s PatchSlicer) RplDes(ind int) bool {
	for _, x := range s[ind] {
		if x.Ope == "replace" && x.Pat == "/desc/data" {
			return true
		}
	}

	return false
}

// RplFee expresses whether any patch object within the indexed list defines the
// feed time path to be replaced. Here ind describes the list of patches linked
// to the list object in question.
func (s PatchSlicer) RplFee(ind int) bool {
	for _, x := range s[ind] {
		if x.Ope == "replace" && x.Pat == "/feed/data" {
			return true
		}
	}

	return false
}

func (s PatchSlicer) UniTim(ind int) []*Patch {
	for i, x := range s[ind] {
		if x.Ope == "replace" && x.Pat == "/feed/data" {
			if x.Val == "" || x.Val == "0" {
				s[ind][i].Val = "0001-01-01T00:00:00Z" // zero time
			} else {
				s[ind][i].Val = time.Unix(musNum(x.Val), 0).UTC().Format(Layout)
			}
		}
	}

	return s[ind]
}

func musNum(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}

	return num
}
