package resources

import (
	"github.com/sreeram-boyapati/razorpay-go/razorpay"
	"net/http"
)

type BaseResource struct {
	client razorpay.Client
}

func (b *BaseResource) getUrl(
	url string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {
	return b.client.Get(url, data, options)
}

func (b *BaseResource) postUrl(
	url string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {
	return b.client.PostJson(url, data, options)
}

func (b *BaseResource) putUrl(
	url string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {
	return b.client.PutJson(url, data, options)
}

func (b *BaseResource) deleteUrl(
	url string,
	options map[string]string) (*http.Response, error) {
	return b.client.Delete(url, options)
}
