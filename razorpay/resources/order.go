package resources

import (
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"net/http"
)

type Order struct {
	base BaseResource
}

func (order *Order) all(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return order.base.getUrl(constants.ORDER_URL, data, options)
}

func (order *Order) fetch(
	order_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.sprintf("%s/%s", constants.ORDER_URL, order_id)

	return cust.base.getUrl(url, data, options)
}

func (order *Order) create(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return cust.base.postUrl(constants.ORDER_URL, data, options)
}

func (order *Order) edit(
	order_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.sprintf("%s/%s", constants.ORDER_URL, order_id)

	return cust.base.PutUrl(url, data, options)
}

func (order *Order) payments(
	order_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s/payments", constants.ORDER_URL, order_id)
}
