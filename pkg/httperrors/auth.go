package httperrors

import "trade-shop/pkg/models"

var WrongUsernameOrPassword = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "001",
	Description: "Wrong username or password",
}}
var WrongAccess = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "002",
	Description: "Wrong access",
}}
