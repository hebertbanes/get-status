package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertbanes/get-status/models"
	res "github.com/hebertbanes/get-status/responses"
)

// GetStatusHandler : Endpoint Get
func GetStatusHandler(c *gin.Context) {

	// bind and validate params
	var params []models.Service
	if err := c.ShouldBind(&params); err != nil {
		res.APIBadRequestMsg(c, err.Error())
		return
	}

	// validate not zero length
	if len(params) == 0 {
		res.APIBadRequestMsg(c, "Al menos debe especificar un servicio")
		return
	}

	// get response
	response := getStatus(params)

	// write response
	res.APIOkData(c, response)
}

// getStatus :
func getStatus(requestedServices []models.Service) (services []models.Service) {
	services = models.GetServices(requestedServices)
	return
}
