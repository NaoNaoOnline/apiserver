package policyhandler

import (
	"context"
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Policy_Update_Fuzz(t *testing.T) {
	var han policy.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *policy.UpdateI
		{
			inp = &policy.UpdateI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Update(context.Background(), inp)
		}
	}
}
