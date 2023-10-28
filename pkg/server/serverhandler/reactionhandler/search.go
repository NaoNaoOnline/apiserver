package reactionhandler

import (
	"context"
	"sort"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *reaction.SearchI) (*reaction.SearchO, error) {
	var err error

	var kin []string
	for _, x := range req.Object {
		if x.Public.Kind != "" {
			kin = append(kin, x.Public.Kind)
		}
	}

	var out []*reactionstorage.Object
	{
		out, err = h.rct.SearchKind(generic.Uni(kin))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Sort reactions by creation time with first priority. This order should
	// reflect our intended row based classification.
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Crtd.UnixNano() < out[j].Crtd.UnixNano()
	})

	//
	// Construct RPC response.
	//

	var res *reaction.SearchO
	{
		res = &reaction.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response got truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "reaction",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		// Reactions marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &reaction.SearchO_Object{
			Intern: &reaction.SearchO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Rctn: x.Rctn.String(),
				User: x.User.String(),
			},
			Public: &reaction.SearchO_Object_Public{
				Html: x.Html,
				Kind: x.Kind,
				Name: x.Name,
			},
		})
	}

	return res, nil
}
