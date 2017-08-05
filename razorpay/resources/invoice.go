package resources

import (
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"net/http"
)

type Invoice struct {
	base BaseResource
}

func (invoice *Invoice) all(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return invoice.base.getUrl(constants.INVOICE_URL, data, options)
}

func (invoice *Invoice) fetch(
	invoice_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s", constants.INVOICE_URL, card_id)

	return invoice.base.getUrl(url, data, options)
}

func (invoice *Invoice) create(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return invoice.base.postUrl(constants.INVOICE_URL, data, options)
}
