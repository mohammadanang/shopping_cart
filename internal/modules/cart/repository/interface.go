package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
)

type Repository interface {
	Create(ctx context.Context, item dbgen.AddCartParams) (*dbgen.Cart, error)
	List(ctx context.Context) ([]*dbgen.Cart, error)
	Show(ctx context.Context, id int64) (*dbgen.Cart, error)
	Update(ctx context.Context, id int64, item dbgen.EditCartParams) (*dbgen.Cart, error)
	Delete(ctx context.Context, id int64) error
}
