package resources

import (
	"github.com/razorpay"
	"net/http"
)

type BaseResource struct {
	client   razorpay.Client
	base_url string
}

type Options map[string]interface{}

func (b *BaseResource) all(id string, data map[string]interface{}, options Options) *http.Response {
	return b.get_url(b.base_url, data, options)
}

func (b *BaseResource) get_url(url string, data map[string]interface{}, options Options) *http.Response {
	return b.client.get(url, data, options)
}
