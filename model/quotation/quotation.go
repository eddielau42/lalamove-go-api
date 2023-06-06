package quotation

import (
	"time"
	
	"github.com/eddielau42/lalamove-go-api/enum"
)

type Quotation struct {
	ServiceType string `json:"serviceType"`
	Stops []DeliveryStop `json:"stops"`
	Language string `json:"language"`

	// optional fields
	ScheduleAt string `json:"scheduleAt,omitempty"`
	SpecialRequests []string `json:"specialRequests,omitempty"`
	IsRouteOptimized bool `json:"isRouteOptimized"`
	Item QuotationItem `json:"item,omitempty"`
}

// SetScheduleAt	设置取货时间; 如果是立即下单，请省略
func (q *Quotation) SetScheduleAt(t time.Time) *Quotation {
	// Note: time in UTC timezone and ISO 8601 format
	const iso8601 = "2006-01-02T15:04:05Z"
	q.ScheduleAt = t.UTC().Format(iso8601)
	return q
}
// SetItem	设置配送物品的信息
func (q *Quotation) SetItem(item QuotationItem) *Quotation {
	q.Item = item
	return q
}
// AddSpecialRequest	添加配送特别要求信息
func (q *Quotation) AddSpecialRequest(strs ...string) *Quotation {
	for _, str := range strs {
		q.SpecialRequests = append(q.SpecialRequests, str)
	}
	return q
}
// AddStop	添加站点 (最少2个, 最多16个)
func (q *Quotation) AddStop(stop DeliveryStop) *Quotation {
	if len(q.Stops) < enum.QUOT_STOPS_MAX {
		q.Stops = append(q.Stops, stop)
	}
	return q
}
func (q *Quotation) SenderStop() DeliveryStop {
	return q.Stops[0]
}
func (q *Quotation) RecipientStops() []DeliveryStop {
	return q.Stops[1:]
}

type QuotationDetail struct {
	ID string `json:"quotationId"`
	ExpiresAt string `json:"expiresAt"`
	PriceBreakdown PriceBreakdown `json:"priceBreakdown"`
	Distance Distance `json:"distance"`

	Quotation
}

type QuotationItem struct {
	Quantity string `json:"quantity"`
	Weight string `json:"weight"`
	Categories []string `json:"categories"`
	HandlingInstructions []string `json:"handlingInstructions"`
}

type DeliveryStop struct {
	ID string `json:"stopId,omitempty"`
	Coordinates Coordinates `json:"coordinates"`
	Address string `json:"address"`

	Name string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Remarks string `json:"remarks,omitempty"`
}

type Coordinates struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type PriceBreakdown struct {
	Base string `json:"base"`
	TotalExcludePriorityFee string `json:"totalExcludePriorityFee"`
	Total string `json:"total"`
	Currency string `json:"currency"`
	PriorityFee string `json:"priorityFee"`

	ExtraMileage string `json:"extraMileage"`
	Surcharge string `json:"surcharge"`
	SpecialRequests string `json:"specialRequests"`
	
	Vat string `json:"vat"`
	TotalBeforeOptimization string `json:"totalBeforeOptimization"`
}
type Distance struct {
	Value string `json:"value"`
	Unit string `json:"unit"`
}
