package order

import (
	"github.com/eddielau42/lalamove-go-api/model/quotation"
)

type Order struct {
	QuotationId string `json:"quotationId"`
	Sender Contact `json:"sender"`
	Recipients []DeliveryDetail `json:"recipients"`

	// optional fields
	IsRecipientSMSEnabled bool `json:"isRecipientSMSEnabled"`
	IsPODEnabled bool `json:"isPODEnabled"`
	Partner string `json:"partner,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// AddRecipient	添加收件人信息;
// 收件人信息包含: stopId - 站点ID, name - 收件人姓名, phone - 收件人手机, remarks - 备注
func (o *Order) AddRecipient(recipient DeliveryDetail) *Order {
	o.Recipients = append(o.Recipients, recipient)
	return o
}
func (o *Order) DisablePOD() *Order {
	o.IsPODEnabled = false
	return o
}
func (o *Order) EnablePOD() *Order {
	o.IsPODEnabled = true
	return o
}
func (o *Order) DisableRecipient() *Order {
	o.IsRecipientSMSEnabled = false
	return o
}
func (o *Order) EnableRecipient() *Order {
	o.IsRecipientSMSEnabled = true
	return o
}

type OrderDetail struct {
	ID string `json:"orderId"`
	QuotationId string `json:"quotationId"`
	DriverId string `json:"driverId"`
	Status string `json:"status"`
	PriorityFee string `json:"priorityFee"`
	ShareLink string `json:"shareLink"`
	Metadata map[string]string `json:"metadata"`
	
	Distance quotation.Distance `json:"distance"`
	Stops []quotation.DeliveryStop `json:"stops"`

	PriceBreakdown quotation.PriceBreakdown `json:"priceBreakdown"`
}

type DeliveryDetail struct {
	StopId string `json:"stopId"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	
	// optional fields
	Remarks string `json:"remarks"`
}

type Contact struct {
	StopId string `json:"stopId"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}