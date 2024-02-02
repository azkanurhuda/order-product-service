package product_sdk

import "context"

type SDK interface {
	GetProductByID(ctx context.Context, ID int) (*ProductResponse, error)
}
