package votehandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Vote_Create_Fuzz(t *testing.T) {
	var han vote.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *vote.CreateI
		{
			inp = &vote.CreateI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Create(tesCtx(), inp)
		}
	}
}
