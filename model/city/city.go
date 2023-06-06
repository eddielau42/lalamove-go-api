package city

type City struct {
	Locode string `json:"locode"`
	Name string `json:"name"`
	Services []CityService `json:"services"`
}

type CityService struct {
	Key string `json:"key"`
	Description string `json:"description"`

	Dimension map[string]Dimension `json:"dimensions"`
	// Dimension map[string]interface{} `json:"dimensions"`

	Load Load `json:"load"`

	SpecialRequests []SpecialRequest `json:"specialRequests"`

	// DeliveryItemSpecification DeliveryItemSpecification `json:"deliveryItemSpecification"`
	DeliveryItemSpecification map[string]interface{} `json:"deliveryItemSpecification"`
}

type Dimension map[string]interface{}

type Load struct {
	Value string `json:"value"`
	Unit string `json:"unit"`
}

type SpecialRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type DeliveryItemSpecification struct {}