package handler

import (
	"bank/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	custSer service.CustomerService
}

func NewCustomerHandler(custSer service.CustomerService) customerHandler {
	return customerHandler{custSer: custSer}
}

func (h customerHandler) GetCustomers(c *gin.Context) {

	customers, err := h.custSer.GetCustomers()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (h customerHandler) GetCustomer(c *gin.Context) {

	id := c.Param("customer_id")
	customers, err := h.custSer.GetCustomer(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers)
}
