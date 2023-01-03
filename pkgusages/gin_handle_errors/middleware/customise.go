package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnyErrorWrapper is used gin.HandlerFunc as a return type.
func AnyErrorWrapper() gin.HandlerFunc {
	return handleErrors(gin.ErrorTypeAny)
}

// handleErrors is used gin.HandlerFunc as a return type, so the return is a func.
func handleErrors(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // The code writes from this line will run after the request is completed, just before it is transmitted to the client.

		// got the caught errors from gin according to the error type that came as a parameter to our function
		detectedErrors := c.Errors.ByType(errType)

		// If the number of caught errors is greater than 0, we returned an error message with 500 internal server code.
		if len(detectedErrors) > 0 {
			c.String(http.StatusInternalServerError, "An error occurred.")
			c.Abort() // Finish the request with the c.Abort(). If you don’t use abort, the request won’t actually be completed.
			return
		}
	}
}
