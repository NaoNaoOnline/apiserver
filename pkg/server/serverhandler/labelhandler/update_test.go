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
		pro objectfield.MapStr
		pat []*labelstorage.Patch
		err error
	}{
		// Case 000 ensures that adding a new profile is allowed.
		{
			pro: objectfield.MapStr{},
			pat: []*labelstorage.Patch{
				{Ope: "add" /*****/, Pat: "/prfl/data/Twitter", Val: "flashbots"},
			},
			err: nil,
		},
		// Case 001 ensures that adding a second profile is allowed.
		{
			pro: objectfield.MapStr{
				Data: map[string]string{
					"Twitter": "FlashbotsFDN",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "add" /*****/, Pat: "/prfl/data/Warpcast", Val: "flashyboys"},
			},
			err: nil,
		},
		// Case 002 ensures that removing an existing profile is allowed.
		{
			pro: objectfield.MapStr{
				Data: map[string]string{
					"Twitter": "FlashbotsFDN",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "remove" /**/, Pat: "/prfl/data/Twitter", Val: "FlashbotsFDN"},
			},
			err: nil,
		},
		// Case 003 ensures that replacing an existing profile is allowed.
		{
			pro: objectfield.MapStr{
				Data: map[string]string{
					"Twitter":  "FlashbotsFDN",
					"Warpcast": "flashbots",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "remove" /**/, Pat: "/prfl/data/Twitter", Val: "FlashbotsFDN"},
				{Ope: "add" /*****/, Pat: "/prfl/data/Twitter", Val: "flashbots"},
			},
			err: nil,
		},
		// Case 004 ensures that adding an existing profile is not allowed.
		{
			pro: objectfield.MapStr{
				Data: map[string]string{
					"Twitter": "FlashbotsFDN",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "add" /**/, Pat: "/prfl/data/Twitter", Val: "flashbots"},
			},
			err: labelProfileAlreadyExistsError,
		},
		// Case 005 ensures that removing a profile that does not exist is not
		// allowed.
		{
			pro: objectfield.MapStr{
				Data: map[string]string{
					"Twitter":  "FlashbotsFDN",
					"Warpcast": "flashbots",
				},
			},
			pat: []*labelstorage.Patch{
				{Ope: "remove" /**/, Pat: "/prfl/data/Foo", Val: "Bar"},
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
