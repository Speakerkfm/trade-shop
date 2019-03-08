package httperrors

import "trade-shop/pkg/models"

var WrongUsernameOrPassword = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "001",
	Description: "Wrong email or password",
}}
var EmailIsTaken = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "002",
	Description: "Email address is already taken",
}}
