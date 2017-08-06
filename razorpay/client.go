package razorpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/errors"
	"net/http"
	"time"
)

type Auth struct {
	key    string
	secret string
}

type Client struct {
	auth       Auth
	appDetails map[string]string
}

func (c *Client) setAppDetails(data map[string]string) {
	for k, v := range data {
		c.appDetails[k] = v
	}
}

func (c *Client) getAppDetails() map[string]string {
	return c.appDetails
}

func (c *Client) makeRequestAndGetResponse(req *http.Request) (*http.Response, error) {

	client := c.getHttpClient()

	response, err := client.Do(req)

	if response.StatusCode >= constants.HTTP_STATUS_OK &&
		response.StatusCode < constants.HTTP_STATUS_REDIRECT {
		return response, err
	}

	// Raise an error depending on the type of error in response
	var jsonResponse errors.ErrorJson
	json.NewDecoder(response.Body).Decode(jsonResponse)
	error_data := jsonResponse.ErrorData

	switch error_data.GetInternalErrorCode() {
	case errors.SERVER_ERROR:
		return nil, &errors.ServerError{Message: error_data.GetDescription()}
	case errors.GATEWAY_ERROR:
		return nil, &errors.GatewayError{Message: error_data.GetDescription()}
	case errors.BAD_REQUEST_ERROR:
	default:
		return nil, &errors.BadRequestError{Message: error_data.GetDescription()}
	}

	return response, err
}

func (c *Client) Get(
	url string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	newUrl := fmt.Sprintf("%s/%s", constants.BASE_URL, url)

	encoded_url := parseUrlOptions(newUrl, data)

	req, _ := http.NewRequest("GET", encoded_url, nil)

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	c.addHeaders(req, options)

	return c.makeRequestAndGetResponse(req)
}

func (c *Client) MakeRequestAndGetResponse(
	req *http.Request) (*http.Response, error) {

	client := c.getHttpClient()

	response, err := client.Do(req)

	if response.StatusCode >= constants.HTTP_STATUS_OK &&
		response.StatusCode < constants.HTTP_STATUS_REDIRECT {
		return response, err
	}

	// Raise an error depending on the type of error in response
	var jsonResponse errors.ErrorJson
	json.NewDecoder(response.Body).Decode(jsonResponse)
	error_data := jsonResponse.errorData

	switch error_data.getInternalErrorCode() {
	case errors.SERVER_ERROR:
		return nil, errors.ServerError(error_data.getDescription())
	case errors.GATEWAY_ERROR:
		return nil, errors.GatewayError(error_data.getDescription())
	case errors.BAD_REQUEST_ERROR:
	default:
		return nil, errors.BAD_REQUEST_ERROR(error_data.getDescription())
	}

}

func (c *Client) PostJson(
	url string,
	payload map[string]interface{},
	options map[string]string) (*http.Response, error) {

	jsonStr, _ := json.Marshal(payload)

	newUrl := fmt.Sprintf("%s/%s", constants.BASE_URL, url)

	req, _ := http.NewRequest("POST", newUrl, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	req.Header.Set("Content-Type", "application/json")

	c.addHeaders(req, options)

	return c.makeRequestAndGetResponse(req)
}

func (c *Client) PutJson(
	url string,
	payload map[string]interface{},
	options map[string]string) (*http.Response, error) {

	jsonStr, _ := json.Marshal(payload)

	newUrl := fmt.Sprintf("%s/%s", constants.BASE_URL, url)

	req, _ := http.NewRequest("PUT", newUrl, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	req.Header.Set("Content-Type", "application/json")

	c.addHeaders(req, options)

	return c.makeRequestAndGetResponse(req)
}

func (c *Client) Delete(
	url string,
	options map[string]string) (*http.Response, error) {

	newUrl := fmt.Sprintf("%s/%s", constants.BASE_URL, url)

	req, _ := http.NewRequest("DELETE", newUrl, nil)

	req.SetBasicAuth(c.auth.key, c.auth.secret)

	c.addHeaders(req, options)

	return c.makeRequestAndGetResponse(req)
}

func (c *Client) addHeaders(req *http.Request, options map[string]string) {

	req.Header.Set("User-Agent", "Razorpay-Go/"+getSdkVersion())

	if val, ok := options["Content-Type"]; ok {
		req.Header.Set("Content-Type", val)
	}
}

func (c *Client) getHttpClient() *http.Client {

	httpClient := &http.Client{Timeout: 10 * time.Second}

	return httpClient
}
