package helper

import "github.com/gin-gonic/gin"

func getErrorMap(err *APIError) map[string]interface{} {
	errObj := map[string]interface{}{
		"error": map[string]interface{}{
			"err_code": err.Code,
			"err_str":  err.Title,
			"err_msg":  err.Message,
		},
	}
	return errObj
}

func WriteErrorResponse(ctx *gin.Context, err *APIError) {
	ctx.JSON(err.Status, getErrorMap(err))
}

func getSuccessMap(data interface{}) map[string]interface{} {
	successMap := map[string]interface{}{
		"results": data,
		"error": map[string]interface{}{
			"err_msg":  nil,
			"err_code": nil,
			"err_str":  nil,
		},
	}
	return successMap
}

func WriteSuccessResponse(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, getSuccessMap(data))
}
