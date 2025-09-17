package xendit

import (
	"context"
	"strconv"

	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
	"github.com/mohammadanang/shopping-cart/pkg/config"
	"github.com/mohammadanang/shopping-cart/pkg/util"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
	"github.com/xendit/xendit-go/qrcode"
	"github.com/xendit/xendit-go/retailoutlet"
	"github.com/xendit/xendit-go/virtualaccount"
)

type XenditPayment struct {
	cfg *config.Config
}

func NewXendit(cfg *config.Config) Xendit {
	xendit.Opt.SecretKey = cfg.Env.XenditSecretKey

	return &XenditPayment{
		cfg: cfg,
	}
}

func (x *XenditPayment) CreateVirtualAccount(ctx context.Context, order domain.Order, bank string) (*xendit.VirtualAccount, error) {
	orderId := strconv.Itoa(int(order.ID))
	params := virtualaccount.CreateFixedVAParams{
		ExternalID:     orderId,
		BankCode:       bank,
		Name:           order.Buyer,
		ExpectedAmount: order.Total,
	}

	resp, err := virtualaccount.CreateFixedVAWithContext(ctx, &params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (x *XenditPayment) CreateEWallet(ctx context.Context, order domain.Order, channel string) (*xendit.EWalletCharge, error) {
	orderId := strconv.Itoa(int(order.ID))
	params := ewallet.CreateEWalletChargeParams{
		ReferenceID:    orderId,
		Currency:       "IDR",
		Amount:         order.Total,
		ChannelCode:    channel,
		CheckoutMethod: util.EWALLET_METHOD,
	}
	resp, err := ewallet.CreateEWalletChargeWithContext(ctx, &params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (x *XenditPayment) CreateQRIS(ctx context.Context, order domain.Order) (*xendit.QRCode, error) {
	orderId := strconv.Itoa(int(order.ID))
	params := qrcode.CreateQRCodeParams{
		ExternalID:  orderId,
		Type:        "qris",
		CallbackURL: "https://cart.anangm182.com/api/v1/orders/" + orderId,
		Amount:      order.Total,
	}
	resp, err := qrcode.CreateQRCodeWithContext(ctx, &params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (x *XenditPayment) CreateRetail(ctx context.Context, order domain.Order, outlet string) (*xendit.RetailOutlet, error) {
	outletName := xendit.RetailOutletNameEnum(outlet)
	orderId := strconv.Itoa(int(order.ID))
	params := retailoutlet.CreateFixedPaymentCodeParams{
		ExternalID:       orderId,
		RetailOutletName: outletName,
		Name:             order.Buyer,
		ExpectedAmount:   order.Total,
	}
	resp, err := retailoutlet.CreateFixedPaymentCodeWithContext(ctx, &params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
