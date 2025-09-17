package xendit

import (
	"context"

	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
	"github.com/xendit/xendit-go"
)

type Xendit interface {
	CreateVirtualAccount(ctx context.Context, order domain.Order, bank string) (*xendit.VirtualAccount, error)
	CreateEWallet(ctx context.Context, order domain.Order, channel string) (*xendit.EWalletCharge, error)
	CreateQRIS(ctx context.Context, order domain.Order) (*xendit.QRCode, error)
	CreateRetail(ctx context.Context, order domain.Order, outlet string) (*xendit.RetailOutlet, error)
}
