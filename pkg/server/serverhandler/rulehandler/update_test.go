package rulehandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Rule_Update_Fuzz(t *testing.T) {
	var han rule.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *rule.UpdateI
		{
			inp = &rule.UpdateI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Update(tesCtx(), inp)
		}
	}
}
