package models

import (
	"net/http"
	"time"
)

// Service : Service type
type Service struct {
	URL     string `json:"url" form:"url"`
	Type    string `json:"type" form:"type"`
	Timeout uint   `json:"timeout" form:"timeout"`
	Status  int    `json:"status" form:"status"`
	Error   string `json:"error" form:"error"`
}

// GetServices : function that do requests simultaneously with goroutines
func GetServices(requestedServices []Service) (services []Service) {
	// channel for services
	chanService := make(chan Service)

	// iterates requestedServices
	for _, reqService := range requestedServices {

		// call doGet with goroutines
		go doGet(reqService, chanService)
	}

	// receives services from chanServices when ready
	for range requestedServices {
		services = append(services, <-chanService)
	}

	return
}

// doGet : function that make the real http requests
func doGet(reqService Service, chanService chan Service) {
	// copy basic fields
	service := Service{
		URL:     reqService.URL,
		Type:    reqService.Type,
		Timeout: reqService.Timeout,
	}

	// prepare client
	client := &http.Client{Timeout: time.Duration(reqService.Timeout) * time.Second}

	// call to method
	response, err := client.Get(reqService.URL)
	if err != nil {
		service.Status = 500
		service.Error = err.Error()
	} else {
		response.Body.Close()
		service.Status = response.StatusCode
	}

	// send to channel
	chanService <- service
}
