package razorpay

import "net/http"

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

func (c *Client) setBaseUrl(url string) {
	c.base_url = url
}

func (c *Client) parseOptions(options Options) {
}

func (c *Client) Get(url string, data Payload, options Options) (*http.Response, error) {

	client := c.getHttpClient()

	encoded_url := parseUrlOptions(url, data)

	req, _ := http.NewRequest("GET", encoded_url, nil)

	c.constructHeaders(req, options)

	return client.Do(req)
}

func (c *Client) Post(url string, data Payload, options Options) (*http.Response, error) {

	client := c.getHttpClient()

	req, _ := http.NewRequest("GET", encoded_url, nil)

	c.constructHeaders(req, options)

	return client.Do(req)
}

func (c *Client) constructHeaders(req *http.Request, options Options) {

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	req.Header.Set("User-Agent", "Razorpay-Go/"+getSdkVersion())

	if val, ok := options["Content-Type"]; ok {
		req.Header.Set("Content-Type", val)
	} else {
		req.Header.Set("Content-Type", "application/json")
	}
}

func (c *Client) getHttpClient() *http.Client {

	httpClient := &http.Client{}

	return httpClient
}
