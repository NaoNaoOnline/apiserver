package wallethandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
)

func (h *Handler) Delete(ctx context.Context, req *wallet.DeleteI) (*wallet.DeleteO, error) {
	fmt.Printf("/wallet.API/Delete not implemented\n")
	return &wallet.DeleteO{}, nil
}
