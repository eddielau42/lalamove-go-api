package lalamove

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eddielau42/lalamove-go-api/enum"
	"github.com/eddielau42/lalamove-go-api/model/order"
	"github.com/eddielau42/lalamove-go-api/model/quotation"
)

const (
	apikey = "pk_test_********************************"
	secret = "sk_test_****************************************************************"
)

var (
	cli *Client
)

func init() {
	cli = NewClient(Config{
		apikey: apikey,
		secret: secret,
		country: enum.COUNTRY_HONGKONG,
	})
	// Set sandbox mode
	cli.Sandbox()
	// Set debug mode
	cli.Debug(true)
}

func TestClient(t *testing.T) {
	// t.Logf("----> lalamove-client: %+v\n", cli)
	assert.True(t, cli.IsSandbox())

	// Change country
	cli.SetCountry(enum.COUNTRY_TAIWAN)
	assert.EqualValues(t, cli.GetCountry(), enum.COUNTRY_TAIWAN)
}

func TestGetQuotations(t *testing.T) {
	q := &quotation.Quotation{
		ServiceType: enum.SERVICE_TYPE_MOTORCYCLE,
		Language: enum.LANG_ZH_HK,
	}
	q.AddStop(quotation.DeliveryStop{
		Address: "Innocentre, 72 Tat Chee Ave, Kowloon Tong",
		Coordinates: quotation.Coordinates{
			Lat: "22.33547351186244",
			Lng: "114.17615807116502",
		},
	}).
	AddStop(quotation.DeliveryStop{
		Address: "Canton Rd, Tsim Sha Tsui",
		Coordinates: quotation.Coordinates{
			Lat: "22.29553167157697",
			Lng: "114.16885175766998",
		},
	}).
	// Optional fields - item
	SetItem(quotation.QuotationItem{
		Quantity: "12",
		Weight: "LESS_THAN_3_KG",
		Categories: []string{"FOOD_DELIVERY", "OFFICE_ITEM"},
		HandlingInstructions: []string{"KEEP_UPRIGHT"},
	}).
	// Optional fields - specialRequests
	AddSpecialRequest("TOLL_FEE_10")
	
	// Optional fields - isRouteOptimized
	q.IsRouteOptimized = true

	// Optional fields - scheduleAt
	// scheduledAt := time.Now()
	// scheduledAt.Add(2 * time.Hour)
	// q.SetScheduleAt(scheduledAt)

	t.Logf("\n----> 提交 quotation : %+v \n", q)
	
	_q, err := cli.GetQuotations(q)
	t.Logf("\n----> response: %+v", _q)
	if err != nil {
		t.Logf("\n----> GetQuotations_error: %s", err.Error())
	}

	assert.NotEmpty(t, _q.ID)
}

func TestGetQuotationDetail(t *testing.T) {
	id := "2717235714549903754"
	t.Logf("\n----> quotationID: %s \n", id)

	q, err := cli.GetQuotationDetail(id)
	if err != nil {
		t.Logf("\n----> GetQuotationDetail_error: %s", err.Error())
	}
	
	if assert.NotNil(t, q) {
		t.Logf("\n----> result - quotationDetail is: %+v", q)
		assert.NotEmpty(t, q)
		assert.Equal(t, id, q.ID)
	}
}

func TestPlaceOrder(t *testing.T) {
	qID := "2717235714549903754"
	qd, _ := cli.GetQuotationDetail(qID)
	// t.Logf("\n----> quotation_detail:\n%+v\n", qd)

	o := &order.Order{
		QuotationId: qID,
		Sender: order.Contact{
			// StopId: "2714578206857392161",
			StopId: qd.SenderStop().ID,
			Name: "Michal",
			Phone: "+85238485765",
		},

		// Optional fields
		// IsPODEnabled: true,
		// IsRecipientSMSEnabled: true,
		// Partner: "Lalamove Partner ",
	}

	// o.DisablePOD()
	// o.EnablePOD()

	// o.DisableRecipient()
	// o.EnableRecipient()

	for _, recipient := range qd.RecipientStops() {
		o.AddRecipient(order.DeliveryDetail{
			// StopId: "2714578206857392162",
			StopId: recipient.ID,
			Name: "Katrina",
			Phone: "+85238485760",
			Remarks: "YYYYYY",
		})
	}

	t.Logf("\n----> 提交 order : %+v \n", o)

	od, err := cli.PlaceOrder(o)
	if err != nil {
		t.Logf("\n----> PlaceOrder_error: %s", err.Error())

		// API:PlaceOrder >> response : 
		// ----> PlaceOrder_error: [ERR_INSUFFICIENT_CREDIT] You do not have sufficient credit for the action, please top-up your wallet
	}
	if assert.NotNil(t, od) {
		assert.NotEmpty(t, od.ID)
		t.Logf("\n----> orderID: %s", od.ID)
	}
}

func TestGetOrderDetail(t *testing.T) {
	id := "107900701184"
	t.Logf("\n----> orderID: %s \n", id)

	o, err := cli.GetOrderDetail(id)
	if err != nil {
		t.Logf("\n----> GetOrderDetail_error: %s", err.Error())
	}
	if assert.NotNil(t, o) {
		t.Logf("\n----> result - orderDetail is: %+v", o)
		assert.NotEmpty(t, o)
		assert.Equal(t, id, o.ID)
	}
}

func TestEditOrder(t *testing.T) {
	id := "107900701184"

	stops := make([]quotation.DeliveryStop, 0)
	
	stops = append(stops, quotation.DeliveryStop{
		Name: "Michal",
		Phone: "+85238485765",
		Address: "Innocentre, 72 Tat Chee Ave, Kowloon Tong",
		Coordinates: quotation.Coordinates{
			Lat: "22.3354735",
			Lng: "114.1761581",
		},
	})
	stops = append(stops, quotation.DeliveryStop{
		Name: "Michal",
		Phone: "+85212345679",
		Address: "Telegraph Bay, Cyberport Rd, 薄扶林 Cyberport 1",
		Coordinates: quotation.Coordinates{
			Lat: "22.26308035863828",
			Lng: "114.13081794602759",
		},
	})

	o, err := cli.EditOrder(id, stops)
	if err != nil {
		t.Logf("\n----> GetOrderDetail_error: %s", err.Error())
	}
	if assert.NotNil(t, o) {
		t.Logf("\n----> result - orderDetail is: %+v", o)
		assert.NotEmpty(t, o)
		assert.Equal(t, id, o.ID)
	}
}

func TestChangeDriver(t *testing.T) {
	orderID := ""
	driverID := ""
	reason := enum.RESON_LATE

	ok, err := cli.ChangeDriver(orderID, driverID, reason)
	if err != nil {
		t.Logf("\n----> PlaceOrder_error: %s", err.Error())
	}

	assert.True(t, ok)
}

func TestGetCityInfo(t *testing.T) {
	// cli.SetCountry(enum.COUNTRY_PHILIPPINES)
	cli.SetCountry("CN") // invalidate value
	cities, err := cli.GetCityInfo()
	if err != nil {
		t.Logf("\n----> GetCityInfo_error: %s", err.Error())
	}
	if assert.NotNil(t, cities) {
		assert.NotEmpty(t, cities)
		// t.Logf("\n----> city: %+v", cities)

		assert.GreaterOrEqual(t, len(cities), 0)

		assert.NotEmpty(t, cities[0].Locode)
		assert.NotEmpty(t, cities[0].Name)

		assert.NotEmpty(t, cities[0].Services)
		assert.GreaterOrEqual(t, len(cities[0].Services), 0)
		assert.NotEmpty(t, cities[0].Services[0].Key)
		assert.NotEmpty(t, cities[0].Services[0].Description)
		assert.NotEmpty(t, cities[0].Services[0].Dimension)
		assert.NotEmpty(t, cities[0].Services[0].Load)
		assert.NotEmpty(t, cities[0].Services[0].SpecialRequests)
		// assert.NotEmpty(t, cities[0].Services[0].DeliveryItemSpecification)
	}
}