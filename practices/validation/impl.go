package validation

import (
	"github.com/gin-gonic/gin"
)

// User represents registration data
type User struct {
	// Must be 3-30 chars
	Username string `json:"username" binding:"required,min=3,max=30"`
	// Must be valid email
	Email string `json:"email" binding:"required,email"`
	// Min 8 chars
	Password string `json:"password" binding:"required,min=8"`
	// Must be 18 or older
	Age int `json:"age" binding:"required,gte=18"`
}

// OrderItem represents individual items in an order
type OrderItem struct {
	ProductID string  `json:"product_id" binding:"required"`
	Name      string  `json:"name" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0,lte=10"`
	Price     float64 `json:"price" binding:"required,gt=0"`
}

// Order represents the complete order structure
type Order struct {
	CustomerID  string      `json:"customer_id" binding:"required"`
	Items       []OrderItem `json:"items" binding:"required,dive,required"`
	TotalAmount float64     `json:"total_amount" binding:"required,gt=0"`
	Status      string      `json:"status" binding:"required,oneof=pending processing completed"`
	Notes       string      `json:"notes" binding:"omitempty,max=500"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// func main() {
// 	router := setupRouter()
// 	router.Run(":8002")
// }

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/register", handleRegistration)
	router.POST("/orders", handleOrders)

	return router
}

// handleRegistration process register request
func handleRegistration(c *gin.Context) {
	var user User

	if err := bindAndValidate(c, &user); err != nil {
		sendErrorResponse(c, 400, err.Error())
		return
	}

	sendSuccessResponse(c, "User registration successful", gin.H{
		"user": user,
	})
}

func handleOrders(c *gin.Context) {
	var order Order

	if err := bindAndValidate(c, &order); err != nil {
		sendErrorResponse(c, 400, err.Error())
		return
	}

	// Success response
	sendSuccessResponse(c, "Order created successfully", gin.H{
		"order": order,
	})
}

func bindAndValidate(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}
	return nil
}

func sendSuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func sendErrorResponse(c *gin.Context, code int, errorMsg string) {
	c.JSON(code, Response{
		Success: false,
		Error:   errorMsg,
	})
}

// Helper function for user-friendly error messages
func getErrorMsg(field string, tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "gt":
		if field == "TotalAmount" || field == "Price" {
			return "Must be greater than 0"
		}
		return "Must be greater than 0"
	case "lte":
		return "Must not exceed 10 items"
	case "oneof":
		return "Must be one of: pending, processing, completed"
	case "max":
		return "Must not exceed 500 characters"
	default:
		return "Invalid value"
	}
}
