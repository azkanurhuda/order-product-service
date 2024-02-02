package costumer_sdk

import (
	"context"
)

type SDK interface {
	GetCustomerByID(ctx context.Context, ID int) (*CostumerResponse, error)
}
