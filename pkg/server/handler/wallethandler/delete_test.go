package wallethandler

import (
	"testing"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	fuzz "github.com/google/gofuzz"
)

func Test_Server_Handler_Wallet_Delete_Fuzz(t *testing.T) {
	var han wallet.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *wallet.DeleteI
		{
			inp = &wallet.DeleteI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Delete(tesCtx(), inp)
		}
	}
}
