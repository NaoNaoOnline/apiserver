package descriptionhandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Description_Create_Fuzz(t *testing.T) {
	var han description.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *description.CreateI
		{
			inp = &description.CreateI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Create(tesCtx(), inp)
		}
	}
}
