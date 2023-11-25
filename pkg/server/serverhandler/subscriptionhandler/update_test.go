package subscriptionhandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Subscription_Update_Fuzz(t *testing.T) {
	var han subscription.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *subscription.UpdateI
		{
			inp = &subscription.UpdateI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Update(tesCtx(), inp)
		}
	}
}
