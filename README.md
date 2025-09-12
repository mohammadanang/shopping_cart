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

Using multiple programming language: Golang (primary language) and Node.js (documentation only)

## Plans

Using xendit as a payment gateway.

## Endpoints

- [cart] Add and/or edit (remove included) items to cart (also create order first) ✅
- [cart] List all carts with order data ✅
- [order] Submit order & payment ❌
- [order] List order history with paging ❌
- [order] Detail order with payment & cart data ❌
- [payment] Complete payment & order (with payment gateway process) ❌
- [payment] List payment history with paging ❌
- [access] Register user access ❌
- [access] Generate access token ❌
- [access] Refresh access token ❌

### Swagger Documentation Page

link: **{baseUrl}/docs/index.html**
json data: **{baseUrl}/openapi.json**
