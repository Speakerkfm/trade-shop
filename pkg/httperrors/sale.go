package httperrors

import "trade-shop/pkg/models"

var NotEnoughItems = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "003",
	Description: "Not enough items",
}}

var DefaultError = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "000",
	Description: "Something goes wrong! :c",
}}

var LotOwner = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "004",
	Description: "You can't buy your lot",
}}

var NotEnoughMoney = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "003",
	Description: "Not enough money",
}}
