package app

import (
	"github.com/Ravi276084/TestGoLangUsingGIN/GoProgramsUsingGIN/domain"
	"github.com/Ravi276084/TestGoLangUsingGIN/GoProgramsUsingGIN/service"
	"github.com/gin-gonic/gin"
)

func RouterSetUp() *gin.Engine {

	router := gin.Default()

	//customerRepo := domain.NewCustomerRepositoryStub()

	// create an instance of the repository
	customerRepo := domain.NewCustomerRepositoryDb()

	// create an instance of the service and inject the repository
	customerService := service.NewCustomerService(customerRepo)

	// create an instance of the handlers and inject the services
	customerHandler := &CustomerHandler{service: customerService}

	router.GET("/customers", customerHandler.GetAllCustomers)
	router.GET("/customers/:id", customerHandler.GetCustomer)

	return router

}

func Start() {
	r := RouterSetUp()
	r.Run("localhost:8000")
}
