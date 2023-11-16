package labelhandler

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Label_Update_Fuzz(t *testing.T) {
	var han label.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *label.UpdateI
		{
			inp = &label.UpdateI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Update(tesCtx(), inp)
		}
	}
}

func Test_Server_Handler_Label_updateVrfyPtch(t *testing.T) {
	testCases := []struct {
		pro map[string]objectfield.String
		pat []*labelstorage.Patch
		err error
	}{
		// Case 000 verifies that adding a new profile is allowed.
		{
			pro: map[string]objectfield.String{},
			pat: []*labelstorage.Patch{
				{Ope: "add" /*****/, Pat: "/prfl/Twitter/data", Val: "flashbots"},
			},
			err: nil,
		},
		// Case 001 verifies that removing an existing profile is allowed.
		{
			pro: map[string]objectfield.String{
				"Twitter": {
					Data: "FlashbotsFDN",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "remove" /**/, Pat: "/prfl/Twitter/data", Val: "FlashbotsFDN"},
			},
			err: nil,
		},
		// Case 002 verifies that replacing an existing profile is allowed.
		{
			pro: map[string]objectfield.String{
				"Twitter": {
					Data: "FlashbotsFDN",
				},
				"Warpcast": {
					Data: "flashbots",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "remove" /**/, Pat: "/prfl/Twitter/data", Val: "FlashbotsFDN"},
				{Ope: "add" /*****/, Pat: "/prfl/Twitter/data", Val: "flashbots"},
			},
			err: nil,
		},
		// Case 003 verifies that adding an existing profile is not allowed.
		{
			pro: map[string]objectfield.String{
				"Twitter": {
					Data: "FlashbotsFDN",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "add" /**/, Pat: "/prfl/Twitter/data", Val: "flashbots"},
			},
			err: labelProfileAlreadyExistsError,
		},
		// Case 004 verifies that removing a profile that does not exist is not
		// allowed.
		{
			pro: map[string]objectfield.String{
				"Twitter": {
					Data: "FlashbotsFDN",
				},
				"Warpcast": {
					Data: "flashbots",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "remove" /**/, Pat: "/prfl/Foo/data", Val: "Bar"},
			},
			err: labelProfileNotFoundError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var use objectid.ID
			{
				use = objectid.ID("user1")
			}

			var ctx context.Context
			{
				ctx = userid.NewContext(context.Background(), use)
			}

			var han *Handler
			{
				han = &Handler{}
			}

			var obj *labelstorage.Object
			{
				obj = &labelstorage.Object{
					Prfl: tc.pro,
					User: objectfield.ID{
						Data: use,
					},
				}
			}

			var pat []*labelstorage.Patch
			{
				pat = tc.pat
			}

			{
				err = han.updateVrfyPtch(ctx, []*labelstorage.Object{obj}, [][]*labelstorage.Patch{pat})
				if !errors.Is(err, tc.err) {
					fmt.Printf("%#v\n", err)
					t.Fatal("expected", true, "got", false)
				}
			}
		})
	}
}
