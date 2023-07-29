package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func (h *Handler) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
	fmt.Printf("apigocode/pkg/label.Search not implemented\n")
	claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	fmt.Printf("%#v\n", ok)
	fmt.Printf("%#v\n", claims)
	return &label.SearchO{}, nil
}
