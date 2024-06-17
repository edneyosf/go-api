package controller

import (
	"base-api/core/app/data/response"
)

func errorResponse(message string) response.ErrorData {
	data := response.ErrorData{ Error: message }

	return data
}
