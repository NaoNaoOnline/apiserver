package listhandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_List_Delete_Fuzz(t *testing.T) {
	var han list.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *list.DeleteI
		{
			inp = &list.DeleteI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Delete(tesCtx(), inp)
		}
	}
}
