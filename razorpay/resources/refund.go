package resources

import (
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"net/http"
)

type Refund struct {
	base BaseResource
}

func (refund *Refund) all(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return refund.base.getUrl(constants.REFUND_URL, data, options)
}

func (refund *Refund) fetch(
	refund_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.sprintf("%s/%s", constants.REFUND_URL, refund_id)

	return cust.base.getUrl(url, data, options)
}

func (refund *Refund) create(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return cust.base.postUrl(constants.REFUND_URL, data, options)
}
