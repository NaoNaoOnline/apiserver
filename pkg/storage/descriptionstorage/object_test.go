package descriptionstorage

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Storage_Description_Object_Verify_Text(t *testing.T) {
	testCases := []struct {
		txt string
		err error
	}{
		// Case 000
		{
			txt: "this is a legit description",
			err: nil,
		},
		// Case 001
		{
			txt: "this too short",
			err: descriptionTextLengthError,
		},
		// Case 002
		{
			txt: "this way too long this way too long this way too long this way too long this way too long this way too long this way too long this way too long this way too long this way too long ",
			err: descriptionTextLengthError,
		},
		// Case 003
		{
			txt: "this one contains weird characters <>",
			err: descriptionTextFormatError,
		},
		// Case 004
		{
			txt: "this one contains my.fake.drop.uk",
			err: descriptionTextURLError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := witTxt(tc.txt).Verify()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %#v got %#v", tc.err, err)
			}
		})
	}
}

func witTxt(txt string) *Object {
	return &Object{
		Evnt: "1234",
		Text: txt,
	}
}
