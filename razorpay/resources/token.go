package resources

import (
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"net/http"
)

type Token struct {
	base BaseResource
}

func (token *Token) all(
	customer_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s/tokens", constants.CUSTOMER_URL, customer_id)

	return token.base.getUrl(url, data, options)
}

func (token *Token) fetch(
	customer_id string,
	token_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, customer_id, token_id)

	return cust.base.getUrl(url, data, options)
}

func (token *Token) create(
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	return cust.base.postUrl(constants.TOKEN_URL, data, options)
}
