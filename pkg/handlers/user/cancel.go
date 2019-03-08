package user

import (
	"trade-shop/pkg/httperrors"
	"trade-shop/pkg/restapi/operations/sales"
	"trade-shop/pkg/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	uuid "github.com/satori/go.uuid"
)

func (c *Context) SaleCancel(params user.CancelParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return user.NewCancelUnauthorized()
	}

	sellerID, err := c.st.GetSellerBySaleID(uuid.FromStringOrNil(params.SaleID.String()))
	if err != nil {
		return user.NewCancelBadRequest().WithPayload(&httperrors.LotDoesNotExist)
	}

	if sellerID != *userID {
		return user.NewCancelBadRequest().WithPayload(&httperrors.NotYourLot)
	}

	if err := c.sale.Cancel(*userID, uuid.FromStringOrNil(params.SaleID.String())); err != nil {
		return user.NewCancelBadRequest().WithPayload(&httperrors.DefaultError)
	}

	return sales.NewBuyOK()
}
