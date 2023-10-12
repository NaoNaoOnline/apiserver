package userhandler

import (
	"context"
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Policy_Create_Fuzz(t *testing.T) {
	var han user.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *user.CreateI
		{
			inp = &user.CreateI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Create(context.Background(), inp)
		}
	}
}
