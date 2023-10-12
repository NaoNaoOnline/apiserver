package labelhandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Label_Search_Fuzz(t *testing.T) {
	var han label.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *label.SearchI
		{
			inp = &label.SearchI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Search(tesCtx(), inp)
		}
	}
}
