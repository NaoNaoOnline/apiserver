package labelhandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
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
