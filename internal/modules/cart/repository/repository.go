package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
)

type CartRepository struct {
	q *dbgen.Queries
}

func NewCartRepository(queries *dbgen.Queries) Repository {
	return &CartRepository{
		q: queries,
	}
}

func (r *CartRepository) Create(ctx context.Context, item dbgen.AddCartParams) (*dbgen.Cart, error) {
	cart, err := r.q.AddCart(ctx, &item)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *CartRepository) List(ctx context.Context) ([]*dbgen.Cart, error) {
	carts, err := r.q.ListCarts(ctx)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *CartRepository) Show(ctx context.Context, id int64) (*dbgen.Cart, error) {
	cart, err := r.q.GetCart(ctx, id)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *CartRepository) Update(ctx context.Context, id int64, item dbgen.EditCartParams) (*dbgen.Cart, error) {
	cart, err := r.q.EditCart(ctx, &item)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *CartRepository) Delete(ctx context.Context, id int64) error {
	return r.q.RemoveCart(ctx, id)
}
