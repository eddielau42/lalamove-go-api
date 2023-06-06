package lalamove

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/eddielau42/lalamove-go-api/logger"
	"github.com/eddielau42/lalamove-go-api/model/city"
	"github.com/eddielau42/lalamove-go-api/model/driver"
	"github.com/eddielau42/lalamove-go-api/model/order"
	"github.com/eddielau42/lalamove-go-api/model/quotation"
	"github.com/eddielau42/lalamove-go-api/util"
)

const (
	baseURL = "https://rest.lalamove.com"
	sandboxURL = "https://rest.sandbox.lalamove.com"

	Version = "v3"
)

// 客户端
type Client struct {
	apiKey string
	apiSecret string
	country string

	sandboxMode bool

	debug bool
}

type Config struct {
	Apikey string
	Secret string
	Country string
	Logfile string
}

// 创建客户端实例
func NewClient(conf Config) *Client {
	if conf.Logfile == "" {
		conf.Logfile = "../lalamove.log"
	}
	logger.SetFile(conf.Logfile)

	return &Client{
		apiKey: conf.apikey,
		apiSecret: conf.secret,
		country: conf.country,
	}
}
// 设置沙箱环境
func (cli *Client) Sandbox() *Client {
	cli.sandboxMode = true
	return cli
}
// 是否沙箱环境
func (cli Client) IsSandbox() bool {
	return cli.sandboxMode
}
// 调式模式开关; 用于调式打印输出请求过程
func (cli *Client) Debug(toggle bool) *Client {
	cli.debug = toggle
	return cli
}

// 设置国家地区
func (cli *Client)SetCountry(country string) *Client {
	cli.country = country
	return cli
}
// 返回当前国家地区
func (cli Client) GetCountry() string {
	return cli.country
}


// GetQuotations	获取报价单
func (cli *Client) GetQuotations(q *quotation.Quotation) (*quotation.QuotationDetail, error) {
	// [POST] /v3/quotations
	uri := "/" + Version + "/quotations"

	payload, err := json.Marshal(map[string]interface{}{"data": q})
	if err != nil {
		// 解析请求数据失败
		log.Fatalln(err)
		return nil, err
	}
	
	result, err := cli.Request(METHOD_POST, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	data := &struct{
		Data *quotation.QuotationDetail `json:"data"`
	}{}
	err = result.Parse(data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}

// GetQuotationDetail	获取报价单详情
func (cli *Client) GetQuotationDetail(quotationID string) (*quotation.QuotationDetail, error) {
	// [GET] /v3/quotations/{quotationId}
	uri := "/" + Version + "/quotations/" + quotationID
	
	var payload []byte
	result, err := cli.Request(METHOD_GET, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	data := &struct{
		Data *quotation.QuotationDetail `json:"data"`
	}{}
	err = result.Parse(data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
} 

// PlaceOrder	下单
func (cli *Client) PlaceOrder(o *order.Order) (*order.OrderDetail, error) {
	// [POST] /v3/orders
	uri := "/" + Version + "/orders"

	payload, err := json.Marshal(map[string]interface{}{"data": o})
	if err != nil {
		// 解析请求数据失败
		log.Fatalln(err)
		return nil, err
	}

	// fmt.Printf("-----request payload:\n%s\n", string(payload))

	result, err := cli.Request(METHOD_POST, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	data := &struct{
		Data *order.OrderDetail `json:"data"`
	}{}
	err = result.Parse(data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}

// GetOrderDetail	获取订单详情
func (cli *Client) GetOrderDetail(orderID string) (*order.OrderDetail, error) {
	// [GET] /v3/orders/{id}
	uri := "/" + Version + "/orders/" + orderID

	var payload []byte
	result, err := cli.Request(METHOD_GET, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	data := &struct{
		Data *order.OrderDetail `json:"data"`
	}{}
	err = result.Parse(data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
} 

// GetDriverDetail	获取司机信息
func (cli *Client) GetDriverDetail(orderID, driverID string) (*driver.DriverDetail, error) {
	// [GET] /v3/orders/{orderId}/drivers/{driverId}
	uri := "/" + Version + "/orders/" + orderID + "/drivers/" + driverID

	var payload []byte
	result, err := cli.Request(METHOD_GET, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	data := &struct{
		Data *driver.DriverDetail `json:"data"`
	}{}
	err = result.Parse(data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}

// AddPriorityFee	添加小费
func (cli *Client) AddPriorityFee(orderID, fee string) (*order.OrderDetail, error) {
	// [POST] /v3/orders/{orderId}/priority-fee
	uri := "/" + Version + "/orders/" + orderID + "/priority-fee"

	payload, err := json.Marshal(map[string]interface{}{
		"data": map[string]interface{}{
			"priorityFee": fee,
		},
	})
	if err != nil {
		// 解析请求数据失败
		log.Fatalln(err)
		return nil, err
	}

	result, err := cli.Request(METHOD_POST, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	data := &struct{
		Data *order.OrderDetail `json:"data"`
	}{}
	err = result.Parse(data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}

// EditOrder	编辑修改订单
func (cli *Client) EditOrder(orderID string, stops []quotation.DeliveryStop) (*order.OrderDetail, error) {
	// [PATCH] /v3/orders/{orderId}
	uri := "/" + Version + "/orders/" + orderID

	payload, err := json.Marshal(map[string]interface{}{
		"data": map[string]interface{}{
			"stops": stops,
		},
	})
	if err != nil {
		// 解析请求数据失败
		log.Fatalln(err)
		return nil, err
	}

	result, err := cli.Request(METHOD_PATCH, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	data := &struct{
		Data *order.OrderDetail `json:"data"`
	}{}
	err = result.Parse(data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}

// CancelOrder	取消订单
func (cli *Client) CancelOrder(orderID string) (bool, error) {
	// [DELETE] /v3/orders/{orderId}
	uri := "/" + Version + "/orders/" + orderID

	var payload []byte
	result, err := cli.Request(METHOD_DELETE, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	
	if result.Response.StatusCode == http.StatusNoContent {
		return true, nil
	}

	err = result.Parse(nil)
	if err != nil {
		return false, err
	}
	return false, err
}

// ChangeDriver	更换司机
func (cli *Client) ChangeDriver(orderID, driverID, reason string) (bool, error) {
	// [DELETE] /v3/orders/{orderId}/drivers/{driverId}
	uri := "/" + Version + "/orders/" + orderID + "/drivers/" + driverID

	payload, err := json.Marshal(map[string]interface{}{
		"data": map[string]interface{}{
			"reason": reason,
		},
	})

	result, err := cli.Request(METHOD_DELETE, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}

	if result.Response.StatusCode == http.StatusNoContent {
		return true, nil
	}

	err = result.Parse(nil)
	if err != nil {
		return false, err
	}
	return false, err
}

// GetCityInfo	获取某一市场的所有城市检索信息和支持的配置。信息包括城市、车辆（服务），以及正在支持的特殊要求。
func (cli *Client) GetCityInfo() ([]city.City, error) {
	// [GET] /v3/cities
	uri := "/" + Version + "/cities"

	var payload []byte
	result, err := cli.Request(METHOD_GET, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// fmt.Printf("------ cities: %s\n", string(result.Body))
	data := &struct{
		Data []city.City `json:"data"`
	}{}
	err = result.Parse(data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (cli *Client) SetWebhook(url string) (bool, error) {
	// [PATCH] /v3/webhook
	uri := "/" + Version + "/webhook"

	payload, err := json.Marshal(map[string]interface{}{
		"data": map[string]interface{}{
			"url": url,
		},
	})
	result, err := cli.Request(METHOD_PATCH, uri, payload)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}

	if result.Response.StatusCode == http.StatusOK {
		return true, nil
	}

	err = result.Parse(nil)
	if err != nil {
		return false, err
	}
	return false, err
}


// 请求方法
const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_PATCH  = "PATCH"
	METHOD_DELETE = "DELETE"
)

// 封装API返回结果
type APIResult struct {
	ReqID string
	Request *http.Request
	Response *http.Response
	Body []byte
}
// Parse	解析返回数据
func (r APIResult) Parse(bindData interface{}) error {
	var err error

	respStatusCode := r.Response.StatusCode
	// 成功
	if respStatusCode >= http.StatusOK && respStatusCode < http.StatusMultipleChoices {
		err = json.Unmarshal(r.Body, &bindData)
		if err != nil {
			fmt.Printf("----- 解析数据错误! error: %s\n", err.Error())
			return err
		}
		return nil
	}

	r.printStackLog()

	// 失败
	if respStatusCode >= http.StatusBadRequest && respStatusCode < http.StatusInternalServerError {
		errData := struct{
			Errors []APIError `json:"errors"`
		}{}
		err = json.Unmarshal(r.Body, &errData)
		if err != nil {
			fmt.Printf("----- 解析数据错误! error: %s\n", err.Error())
			return err
		}
		
		// 取最后一条错误信息内容
		msg := errData.Errors[len(errData.Errors)-1]
		errMsg := "[" + msg.ID + "] " + msg.Message
		if msg.Detail != "" {
			errMsg += " (detail: " + msg.Detail + ")"
		}
		return errors.New(errMsg)
	}

	// 第三方服务异常 (HttpStatusCode >= 500)
	if respStatusCode >= http.StatusInternalServerError {
		errData := struct{
			Message string `json:"message"`
		}{}
		err = json.Unmarshal(r.Body, &errData)
		if err != nil {
			fmt.Printf("----- 解析数据错误! error: %s\n", err.Error())
			return err
		}

		return errors.New(errData.Message)
	}

	return nil
}

// printStackLog	打印API调用信息
func (r APIResult) printStackLog() {
	
	reqHeader, _ := json.Marshal(r.Request.Header)

	logger.Error(
		"----> 请求失败! 打印API调用信息:\n"+
		"------------------------------\n"+
		"Req-ID: %s\nReq-URL: %s\nReq-Method: %s\nReq-Header: %+v\nResp-StatusCode: %d\nResp-Body: %s\n"+
		"------------------------------\n",
		r.ReqID,
		r.Request.URL,
		r.Request.Method,
		string(reqHeader),
		r.Response.StatusCode,
		string(r.Body),
	)
}


type APIError struct {
	ID string `json:"id"`
	Message string `json:"message"`
	Detail string `json:"detail"`
}

// 发起请求
func (cli Client) Request(method, uri string, params []byte) (*APIResult, error) {
	var (
		err error
	)
	result := &APIResult{}

	url := baseURL
	if cli.IsSandbox() { // 沙箱环境
		url = sandboxURL
	}
	url = url + uri

	result.Request, err = http.NewRequest(method, url, bytes.NewBuffer(params))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// 当前时间戳
	ms := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 请求唯一标识
	result.ReqID = util.UniqueID()

	// 签名
	message := fmt.Sprintf("%s\r\n%s\r\n%s\r\n\r\n", ms, method, uri)
	if method != METHOD_GET {
		message = message + string(params)
	}
	signature := util.Signature(cli.apiSecret, message)
	
	// 设置请求头
	result.Request.Header.Add("Content-type", "application/json")
	result.Request.Header.Add("Accept", "application/json")
	result.Request.Header.Add("Request-ID", result.ReqID)
	result.Request.Header.Add("Market", strings.ToUpper(cli.country))
	result.Request.Header.Add("Authorization", fmt.Sprintf("hmac %s:%s:%s", cli.apiKey, ms, signature))	
	
	httpCli := &http.Client{
		Timeout: 30 * time.Second,
	}
	result.Response, err = httpCli.Do(result.Request)
	if err != nil {
		// 请求错误
		log.Fatalln(err)
		return nil, err
	}

	// 将响应数据读取存放到 "result.Body" 中
	result.Body, err = ioutil.ReadAll(result.Response.Body)
	defer result.Response.Body.Close()

	return result, nil
}
