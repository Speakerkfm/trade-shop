package sales

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/satori/go.uuid"
	"trade-shop/pkg/httperrors"
	"trade-shop/pkg/restapi/operations/sales"
)

func (c *Context) BuyLot(params sales.BuyParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return sales.NewBuyUnauthorized()
	}

	sellerID, err := c.st.GetSellerBySaleID(uuid.FromStringOrNil(params.SaleID.String()))
	if err != nil {
		return sales.NewBuyBadRequest().WithPayload(&httperrors.DefaultError)
	}

	if sellerID == *userID {
		return sales.NewBuyBadRequest().WithPayload(&httperrors.LotOwner)
	}

	if err := c.sale.Purchase(*userID, sellerID, uuid.FromStringOrNil(params.SaleID.String())); err != nil {
		return sales.NewBuyBadRequest().WithPayload(&httperrors.NotEnoughMoney)
	}

	return sales.NewBuyOK()
}
