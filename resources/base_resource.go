package resources

import (
	"github.com/razorpay"
	"net/http"
)

type BaseResource struct {
	client   razorpay.Client
	base_url string
}

func (b *BaseResource) all(id string, data razorpay.Payload, options razorpay.Options) *http.Response {
	return b.getUrl(b.base_url, data, options)
}

func (b *BaseResource) getUrl(url string, data razorpay.Payload, options razorpay.Options) *http.Response {
	return b.client.Get(url, data, options)
}
