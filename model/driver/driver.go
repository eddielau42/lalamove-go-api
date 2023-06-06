package driver

import (
	"github.com/eddielau42/lalamove-go-api/model/quotation"
)

type Driver struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type DriverDetail struct {
	ID string `json:"driverId"`

	PlateNo string `json:"plateNumber"`
	Coordinates quotation.Coordinates `json:"coordinates"`
	Driver
}