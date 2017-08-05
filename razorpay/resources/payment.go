package resources

import (
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"net/http"
)

type Payment struct {
	base BaseResource
}

func (p *Payment) all(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return p.base.getUrl(constants.PAYMENT_URL, data, options)
}

func (p *Payment) fetch(
	id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s", constants.PAYMENT_URL, id)

	return p.base.getUrl(url, data, options)
}

func (p *Payment) capture(
	id string,
	amount int,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s/capture", constants.PAYMENT_URL, id)
	// Amount should be in paisa

	data["amount"] = amount

	return p.base.postUrl(url, data, options)
}

func (p *Payment) refund(
	id string,
	amount int,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s/refund", constants.PAYMENT_URL, id)
	// Amount should be in paisa

	data["amount"] = amount

	return p.base.postUrl(url, data, options)
}
