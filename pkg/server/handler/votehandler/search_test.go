package votehandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Vote_Search_Fuzz(t *testing.T) {
	var han vote.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *vote.SearchI
		{
			inp = &vote.SearchI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Search(tesCtx(), inp)
		}
	}
}
