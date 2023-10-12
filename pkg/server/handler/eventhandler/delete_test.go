package eventhandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Event_Delete_Fuzz(t *testing.T) {
	var han event.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *event.DeleteI
		{
			inp = &event.DeleteI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Delete(tesCtx(), inp)
		}
	}
}
