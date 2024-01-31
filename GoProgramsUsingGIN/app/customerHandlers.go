package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ravi276084/TestGoLangUsingGIN/GoProgramsUsingGIN/service"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(c *gin.Context) {
	// customers := []Customer{
	// 	{"Prasad", "Hyd", "500032"},
	// 	{"Ravi", "Hyd", "500032"},
	// }
	//status := c.Request.URL.Query().Get("status")
	customers, _ := ch.service.GetAllCustomers()
	c.JSON(http.StatusOK, customers)

}

func (ch *CustomerHandler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("Received request for Customer ID: %s", id)
	customerid, _ := strconv.Atoi(id)
	fmt.Println("customer_id", +customerid)
	customers, err := ch.service.GetCustomer(customerid)

	if err != nil {
		fmt.Printf("Error: %+v", err)
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	} else {
		c.JSON(http.StatusOK, customers)
	}
}
