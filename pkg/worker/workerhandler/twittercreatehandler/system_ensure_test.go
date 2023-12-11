package twittercreatehandler

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/google/go-cmp/cmp"
)

// go test ./pkg/worker/workerhandler/twittercreatehandler -update
var update = flag.Bool("update", false, "update .golden files")

func Test_Worker_Twitter_Create_ensureTmpl(t *testing.T) {
	testCases := []struct {
		dob *descriptionstorage.Object
		eob *eventstorage.Object
		lob []*labelstorage.Object
	}{
		// Case 000 ensures a normal tweet can be generated with all the given
		// content.
		{
			dob: newDesc("where do I sign up I'm straight yes and amen"),
			eob: newEvnt(),
			lob: newLabl(3, 1),
		},
		// Case 001 ensures a trimmed tweet can be generated with the description
		// being partially cut off.
		{
			dob: newDesc("where do I sign up I'm straight yes and amen coulda, shoulda, woulda, where have you been I'm sick check is in the mail"),
			eob: newEvnt(),
			lob: newLabl(2, 1),
		},
		// Case 002 ensures a trimmed tweet can be generated with the description
		// and a category label being fully cut off.
		{
			dob: newDesc("where do I sign up I'm straight yes and amen coulda, shoulda, woulda, where have you been I'm sick check is in the mail"),
			eob: newEvnt(),
			lob: newLabl(4, 3),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var cur string
			{
				cur, err = ensureTmpl(tc.dob, tc.eob, tc.lob)
				if err != nil {
					t.Fatal(err)
				}
			}

			var pat string
			{
				pat = gldFil(i)
			}

			if *update {
				err := os.WriteFile(pat, []byte(cur), 0600)
				if err != nil {
					t.Fatal(err)
				}
			}

			var des []byte
			{
				des, err = os.ReadFile(pat)
				if err != nil {
					t.Fatal(err)
				}
			}

			if cur != string(des) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(des), cur))
			}

			if len(cur) > 280 {
				t.Fatalf("\n\n%s\n", cmp.Diff(280, len(cur)))
			}
		})
	}
}

func gldFil(i int) string {
	return fmt.Sprintf("golden/case-%03d.txt", i)
}

func newDesc(txt string) *descriptionstorage.Object {
	return &descriptionstorage.Object{
		Text: objectfield.String{
			Data: txt,
		},
	}
}

func newEvnt() *eventstorage.Object {
	return &eventstorage.Object{
		Link: "https://twitter.com/i/spaces/7gqBvHzjUDjOB",
		Time: time.Unix(1700342507, 0).UTC(),
	}
}

func newLabl(x int, y int) []*labelstorage.Object {
	var cat []*labelstorage.Object
	{
		cat = []*labelstorage.Object{
			{
				Kind: "cate",
				Name: objectfield.String{
					Data: "Ethereum Ethereum Ethereum",
				},
			},
			{
				Kind: "cate",
				Name: objectfield.String{
					Data: "Rollups Rollups Rollups",
				},
			},
			{
				Kind: "cate",
				Name: objectfield.String{
					Data: "Real World Assets",
				},
			},
			{
				Kind: "cate",
				Name: objectfield.String{
					Data: "Real World Assets",
				},
			},
		}
	}

	var hos []*labelstorage.Object
	{
		hos = []*labelstorage.Object{
			{
				Kind: "host",
				Name: objectfield.String{
					Data: "Arbitrum",
				},
				Prfl: objectfield.MapStr{
					Data: map[string]string{
						objectlabel.ProfileTwitter: "arbitrum_foundation",
					},
				},
			},
			{
				Kind: "host",
				Name: objectfield.String{
					Data: "Vance",
				},
				Prfl: objectfield.MapStr{
					Data: map[string]string{
						objectlabel.ProfileTwitter: "pythianism_pythianism",
					},
				},
			},
			{
				Kind: "host",
				Name: objectfield.String{
					Data: "Sisyphus",
				},
				Prfl: objectfield.MapStr{
					Data: map[string]string{
						objectlabel.ProfileTwitter: "Sisyphus_Sisyphus",
					},
				},
			},
		}
	}

	var all []*labelstorage.Object
	{
		all = append(all, cat[:x]...)
		all = append(all, hos[:y]...)
	}

	return all
}
