package subscriptionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
)

func (h *Handler) Create(ctx context.Context, req *subscription.CreateI) (*subscription.CreateO, error) {
	fmt.Printf("/subscription.API/Create not implemented\n")
	return &subscription.CreateO{}, nil
}
