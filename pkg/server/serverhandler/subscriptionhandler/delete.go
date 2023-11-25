package subscriptionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
)

func (h *Handler) Delete(ctx context.Context, req *subscription.DeleteI) (*subscription.DeleteO, error) {
	fmt.Printf("/subscription.API/Delete not implemented\n")
	return &subscription.DeleteO{}, nil
}
