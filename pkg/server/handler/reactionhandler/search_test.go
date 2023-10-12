package reactionhandler

import (
	"context"
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Reaction_Search_Fuzz(t *testing.T) {
	var han reaction.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *reaction.SearchI
		{
			inp = &reaction.SearchI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Search(context.Background(), inp)
		}
	}
}
