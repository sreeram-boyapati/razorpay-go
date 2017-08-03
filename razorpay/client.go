package razorpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"net/http"
)

type Auth struct {
	key    string
	secret string
}

type Client struct {
	auth        Auth
	base_url    string
	app_details map[string]string
}

func (c *Client) setAppDetails(data map[string]string) {
	for k, v := range data {
		c.app_details[k] = v
	}
}

func (c *Client) getAppDetails() map[string]string {
	return c.app_details
}

func (c *Client) Get(
	url string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	client := c.getHttpClient()

	newUrl := fmt.Sprintf("%s/%s", constants.BASE_URL, url)

	encoded_url := parseUrlOptions(newUrl, data)

	req, _ := http.NewRequest("GET", encoded_url, nil)

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	c.addHeaders(req, options)

	return client.Do(req)
}

func (c *Client) PostJson(
	url string,
	payload map[string]interface{},
	options map[string]string) (*http.Response, error) {

	client := c.getHttpClient()

	jsonStr, _ := json.Marshal(payload)

	newUrl := fmt.Sprintf("%s/%s", constants.BASE_URL, url)

	req, _ := http.NewRequest("POST", newUrl, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	req.Header.Set("Content-Type", "application/json")

	c.addHeaders(req, options)

	return client.Do(req)
}

func (c *Client) PutJson(
	url string,
	payload map[string]interface{},
	options map[string]string) (*http.Response, error) {

	client := c.getHttpClient()

	jsonStr, _ := json.Marshal(payload)

	newUrl := fmt.Sprintf("%s/%s", constants.BASE_URL, url)

	req, _ := http.NewRequest("PUT", newUrl, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	req.Header.Set("Content-Type", "application/json")

	c.addHeaders(req, options)

	return client.Do(req)
}

func (c *Client) Delete(
	url string,
	options map[string]string) (*http.Response, error) {

	client := c.getHttpClient()

	newUrl := fmt.Sprintf("%s/%s", constants.BASE_URL, url)

	req, _ := http.NewRequest("DELETE", newUrl, nil)

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	c.addHeaders(req, options)

	return client.Do(req)
}

func (c *Client) addHeaders(req *http.Request, options map[string]string) {

	req.Header.Set("User-Agent", "Razorpay-Go/"+getSdkVersion())

	if val, ok := options["Content-Type"]; ok {
		req.Header.Set("Content-Type", val)
	}
}

func (c *Client) getHttpClient() *http.Client {

	httpClient := &http.Client{}

	return httpClient
}
