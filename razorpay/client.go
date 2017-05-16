package razorpay

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Auth struct {
	key    string
	secret string
}

type Options map[string]string

type Payload map[string]interface{}

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

func (c *Client) setBaseUrl(url string) {
	c.base_url = url
}

func (c *Client) parseOptions(options Options) {
}

func (c *Client) Get(url string, data Payload, options Options) (*http.Response, error) {

	client := c.getHttpClient()

	encoded_url := parseUrlOptions(url, data)

	req, _ := http.NewRequest("GET", encoded_url, nil)

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	c.addHeaders(req, options)

	return client.Do(req)
}

func (c *Client) PostJson(requestUrl string, payload Payload, options Options) (*http.Response, error) {

	client := c.getHttpClient()

	jsonStr, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	req.Header.Set("Content-Type", "application/json")

	c.addHeaders(req, options)

	return client.Do(req)
}

func (c *Client) PutJson(requestUrl string, payload Payload, options Options) (*http.Response, error) {

	client := c.getHttpClient()

	jsonStr, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	req.Header.Set("Content-Type", "application/json")

	c.addHeaders(req, options)

	return client.Do(req)
}

func (c *Client) Delete(requestUrl string, options Options) (*http.Response, error) {

	client := c.getHttpClient()

	req, _ := http.NewRequest("DELETE", requestUrl, nil)

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	c.addHeaders(req, options)

	return client.Do(req)
}

func (c *Client) addHeaders(req *http.Request, options Options) {

	req.Header.Set("User-Agent", "Razorpay-Go/"+getSdkVersion())

	if val, ok := options["Content-Type"]; ok {
		req.Header.Set("Content-Type", val)
	}
}

func (c *Client) getHttpClient() *http.Client {

	httpClient := &http.Client{}

	return httpClient
}
