package razorpay

import "net/url"

func parseUrlOptions(requestUrl string, data map[string]interface{}) string {
	var Url *url.URL

	Url, err := url.Parse(requestUrl)

	if err != nil {
		panic(err)
	}

	parameters := url.Values{}

	for k, v := range data {
		parameters.Add(k, v.(string))
	}

	Url.RawQuery = parameters.Encode()

	return Url.String()
}

func getSdkVersion() string {
	return "1.0.0"
}
