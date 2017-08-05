package resources

import (
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"net/http"
)

type Customer struct {
	base BaseResource
}

func (cust *Customer) fetch(
	customer_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.sprintf("%s/%s", constants.CUSTOMER_URL, customer_id)

	return cust.base.getUrl(url, data, options)
}

func (cust *Customer) create(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return cust.base.postUrl(constants.CUSTOMER_URL, data, options)
}

func (cust *Customer) edit(
	customer_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.sprintf(constants.CUSTOMER_URL, customer_id)

	return cust.base.PutUrl(url, data, options)
}
