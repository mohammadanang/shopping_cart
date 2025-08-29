# Shopping Cart

- Go: v1.25
- fiber: v2
- oapi-codegen: v2
- pgx: v5
- paseto
- sqlc
- viper
- golang-migrate
- nodejs: v22 (for documentation/swagger UI only)

## Plans

Using xendit as a payment gateway.

## Endpoints

- [cart] Add and/or edit (remove included) items to cart (also create order first)
- [cart] List all carts with order data
- [order] Submit order & payment
- [order] List order history with paging
- [order] Detail order with payment & cart data
- [payment] Complete payment & order (with payment gateway process)
- [payment] List payment history with paging
