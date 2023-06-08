package enum

// 地区编码
const (
	AREA_CODE_BR = "BR" // Brasil
	AREA_CODE_HK = "HK" // Hong Kong
	AREA_CODE_ID = "ID" // Indonesia
	AREA_CODE_MY = "MY" // Malaysia
	AREA_CODE_MX = "MX" // Mexico
	AREA_CODE_PH = "PH" // Philippines
	AREA_CODE_SG = "SG" // Singapore
	AREA_CODE_TW = "TW" // Taiwan
	AREA_CODE_TH = "TH" // Thailand
	AREA_CODE_VN = "VN" // Vietnam	
)

// 语言
const (
	LANG_EN_BR = "en_BR"
	LANG_PT_BR = "pt_BR"
	LANG_EN_HK = "en_HK"
	LANG_ZH_HK = "zh_HK"
	LANG_EN_ID = "en_ID"
	LANG_ID_ID = "id_ID"
	LANG_EN_MY = "en_MY"
	LANG_MS_MY = "ms_MY"
	LANG_EN_MX = "en_MX"
	LANG_ES_MX = "es_MX"
	LANG_EN_PH = "en_PH"
	LANG_EN_SG = "en_SG"
	LANG_ZH_TW = "zh_TW"
	LANG_TH_TH = "th_TH"
	LANG_EN_TH = "en_TH"
	LANG_EN_VN = "en_VN"
	LANG_VI_VN = "vi_VN"
)

// The type of vehicle
const (
	SERVICE_TYPE_WALKER     = "WALKER"
    SERVICE_TYPE_MOTORCYCLE = "MOTORCYCLE"
    SERVICE_TYPE_CAR        = "CAR"
    SERVICE_TYPE_SEDAN      = "SEDAN"
    SERVICE_TYPE_VAN        = "VAN"
    SERVICE_TYPE_TRUCK175   = "TRUCK175"
    SERVICE_TYPE_TRUCK330   = "TRUCK330"
    SERVICE_TYPE_TRUCK550   = "TRUCK550"
)

// The stops of quotation
const (
	QUOT_STOPS_MIN = 2
	QUOT_STOPS_MAX = 16
)

// Order Status
const (
	ORDER_STATUS_ASSIGN    = "ASSIGNING_DRIVER"
	ORDER_STATUS_GOING     = "ON_GOING"
	ORDER_STATUS_PICKUP    = "PICKED_UP"
	ORDER_STATUS_COMPLETED = "COMPLETED"
	ORDER_STATUS_CANCELED  = "CANCELED"
	ORDER_STATUS_REJECTED  = "REJECTED"
	ORDER_STATUS_EXPIRED   = "EXPIRED"
)

// Proof Of Delivery (POD) Status
const (
	// "PENDING" - The driver hasn't completed the delivery to the stop yet
	POD_STATUS_PENDING = "PENDING"
	// "DELIVERED" -  The driver has completed the order and has taken a photo at the stop
	POD_STATUS_DELIVERED = "DELIVERED"	
	// "SIGNED" - The driver has completed the order and received recipient's signature
	POD_STATUS_SIGNED = "SIGNED"	
	// "FAILED" - The driver couldn't complete the delivery to the stop
	POD_STATUS_FAILED = "FAILED"	
)

// The reason Of change driver
const (
	RESON_LATE         = "DRIVER_LATE"         // Driver is late for delivery
	RESON_CHANGED      = "DRIVER_ASKED_CHANGE" // Driver requested the user to change
	RESON_UNRESPONSIVE = "DRIVER_UNRESPONSIVE" // Driver is not responding
	RESON_RUDE         = "DRIVER_RUDE"         // Driver is rude
)