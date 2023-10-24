package eventhandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Event_Search_Fuzz(t *testing.T) {
	var han event.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *event.SearchI
		{
			inp = &event.SearchI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Search(tesCtx(), inp)
		}
	}
}
