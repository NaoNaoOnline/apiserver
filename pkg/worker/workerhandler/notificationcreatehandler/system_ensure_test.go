package notificationcreatehandler

import (
	"fmt"
	"slices"
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/google/go-cmp/cmp"
)

func Test_Worker_Handler_Notification_filIDs(t *testing.T) {
	testCases := []struct {
		eob *eventstorage.Object
		sli rulestorage.Slicer
		fil []objectid.ID
	}{
		// Case 000
		{
			eob: &eventstorage.Object{
				Cate: []objectid.ID{
					"c44",
					"c66",
				},
				Host: []objectid.ID{
					"h44",
					"h66",
				},
				User: "u88",
			},
			sli: rulestorage.Slicer{
				{List: "l33", Kind: "cate", Excl: []objectid.ID{"c33"}},
			},
			fil: []objectid.ID{
				"l33",
			},
		},
		// Case 001
		{
			eob: &eventstorage.Object{
				Cate: []objectid.ID{
					"c44",
					"c66",
				},
				Host: []objectid.ID{
					"h44",
					"h66",
				},
				User: "u88",
			},
			sli: rulestorage.Slicer{
				{List: "l33", Kind: "cate", Excl: []objectid.ID{"c33"}},
				{List: "l33", Kind: "host", Excl: []objectid.ID{"h33"}},
				{List: "l33", Kind: "user", Excl: []objectid.ID{"u33"}},
			},
			fil: []objectid.ID{
				"l33",
			},
		},
		// Case 002
		{
			eob: &eventstorage.Object{
				Cate: []objectid.ID{
					"c44",
					"c66",
				},
				Host: []objectid.ID{
					"h44",
					"h66",
				},
				User: "u88",
			},
			sli: rulestorage.Slicer{
				{List: "l33", Kind: "cate", Excl: []objectid.ID{"c33"}},
				{List: "l33", Kind: "host", Excl: []objectid.ID{"h33"}},
				{List: "l33", Kind: "user", Excl: []objectid.ID{"u33"}},

				{List: "l44", Kind: "cate", Excl: []objectid.ID{"c44"}}, // exclude by category
				{List: "l44", Kind: "host", Excl: []objectid.ID{"h44"}}, // exclude by host
				{List: "l44", Kind: "user", Excl: []objectid.ID{"u44"}}, // exclude by user

				{List: "l55", Kind: "cate", Excl: []objectid.ID{"c55"}},
				{List: "l55", Kind: "host", Excl: []objectid.ID{"h55"}},
				{List: "l55", Kind: "user", Excl: []objectid.ID{"u55"}},
			},
			fil: []objectid.ID{
				"l33",
				"l55",
			},
		},
		// Case 003
		{
			eob: &eventstorage.Object{
				Cate: []objectid.ID{
					"c44",
					"c66",
				},
				Host: []objectid.ID{
					"h44",
					"h66",
				},
				User: "u88",
			},
			sli: rulestorage.Slicer{
				{List: "l44", Kind: "user", Excl: []objectid.ID{"u44"}},

				{List: "l66", Kind: "cate", Excl: []objectid.ID{"c66"}}, // exclude by category
			},
			fil: []objectid.ID{
				"l44",
			},
		},
		// Case 004
		{
			eob: &eventstorage.Object{
				Cate: []objectid.ID{
					"c44",
					"c66",
				},
				Host: []objectid.ID{
					"h44",
					"h66",
				},
				User: "u88",
			},
			sli: rulestorage.Slicer{
				{List: "l44", Kind: "user", Excl: []objectid.ID{"u44"}},

				{List: "l66", Kind: "host", Excl: []objectid.ID{"h66"}}, // exclude by host
			},
			fil: []objectid.ID{
				"l44",
			},
		},
		// Case 005
		{
			eob: &eventstorage.Object{
				Cate: []objectid.ID{
					"c44",
					"c66",
				},
				Host: []objectid.ID{
					"h44",
					"h66",
				},
				User: "u88",
			},
			sli: rulestorage.Slicer{
				{List: "l88", Kind: "user", Excl: []objectid.ID{"u88"}}, // exclude by user

				{List: "l66", Kind: "host", Excl: []objectid.ID{"h66"}},
			},
			fil: []objectid.ID{},
		},
		// Case 006
		{
			eob: &eventstorage.Object{
				Cate: []objectid.ID{
					"c44",
					"c66",
				},
				Host: []objectid.ID{
					"h44",
					"h66",
				},
				User: "u88",
			},
			sli: rulestorage.Slicer{
				{List: "l22", Kind: "user", Excl: []objectid.ID{"u22"}},

				{List: "l33", Kind: "cate", Excl: []objectid.ID{"c33"}},
				{List: "l33", Kind: "host", Excl: []objectid.ID{"h33"}},
				{List: "l33", Kind: "user", Excl: []objectid.ID{"u33"}},

				{List: "l44", Kind: "host", Excl: []objectid.ID{"h44"}}, // exclude by host

				{List: "l55", Kind: "host", Excl: []objectid.ID{"h55"}},
				{List: "l55", Kind: "user", Excl: []objectid.ID{"u55"}},

				{List: "l66", Kind: "cate", Excl: []objectid.ID{"c66"}}, // excluded by category

				{List: "l77", Kind: "cate", Excl: []objectid.ID{"c77"}},
				{List: "l77", Kind: "user", Excl: []objectid.ID{"u77"}},

				{List: "l88", Kind: "user", Excl: []objectid.ID{"u88"}}, // excluded by user
			},
			fil: []objectid.ID{
				"l22",
				"l33",
				"l55",
				"l77",
			},
		},
		// Case 007
		{
			eob: &eventstorage.Object{
				Cate: []objectid.ID{
					"c44",
					"c66",
				},
				Host: []objectid.ID{
					"h44",
					"h66",
				},
				User: "u88",
			},
			sli: rulestorage.Slicer{
				{List: "l22", Kind: "cate", Excl: []objectid.ID{"c22"}},
				{List: "l22", Kind: "user", Excl: []objectid.ID{"u22"}},

				{List: "l33", Kind: "user", Excl: []objectid.ID{"u33"}},

				{List: "l44", Kind: "cate", Excl: []objectid.ID{"c44"}}, // exclude by category
				{List: "l44", Kind: "user", Excl: []objectid.ID{"u44"}}, // exclude by user

				{List: "l66", Kind: "user", Excl: []objectid.ID{"u66"}},

				{List: "l77", Kind: "cate", Excl: []objectid.ID{"c77"}},

				{List: "l88", Kind: "host", Excl: []objectid.ID{"h88"}},
			},
			fil: []objectid.ID{
				"l22",
				"l33",
				"l66",
				"l77",
				"l88",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			fil := filIDs(tc.eob, tc.sli)

			slices.Sort(fil)

			if !slices.Equal(fil, tc.fil) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.fil, fil))
			}
		})
	}
}

func Test_Worker_Handler_Notification_slcIDs(t *testing.T) {
	testCases := []struct {
		uid []objectid.ID
		lid []objectid.ID
		slc []objectid.ID
		use []objectid.ID
		lis []objectid.ID
	}{
		// Case 000
		{
			uid: []objectid.ID{},
			lid: []objectid.ID{},
			slc: []objectid.ID{},
			use: []objectid.ID{},
			lis: []objectid.ID{},
		},
		// Case 001
		{
			uid: []objectid.ID{
				"u33",
			},
			lid: []objectid.ID{
				"l33",
			},
			slc: []objectid.ID{
				"l33",
			},
			use: []objectid.ID{
				"u33",
			},
			lis: []objectid.ID{
				"l33",
			},
		},
		// Case 002
		{
			uid: []objectid.ID{
				"u33",
				"u44",
				"u55",
			},
			lid: []objectid.ID{
				"l33",
				"l44",
				"l55",
			},
			slc: []objectid.ID{
				"l33",
				"l55",
			},
			use: []objectid.ID{
				"u33",
				"u55",
			},
			lis: []objectid.ID{
				"l33",
				"l55",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			use, lis := slcIDs(tc.uid, tc.lid, tc.slc)

			if !slices.Equal(use, tc.use) {
				t.Fatalf("use\n\n%s\n", cmp.Diff(tc.use, use))
			}
			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("lis\n\n%s\n", cmp.Diff(tc.lis, lis))
			}
		})
	}
}
