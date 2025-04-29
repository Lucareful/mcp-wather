package weather

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	API = "https://restapi.amap.com/v3/weather/weatherInfo"
	KEY = "your personal secret."
)

var httpclient = &fasthttp.Client{
	ReadTimeout:                   500 * time.Millisecond,
	WriteTimeout:                  500 * time.Millisecond,
	MaxIdleConnDuration:           1 * time.Hour,
	NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
	DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
	DisablePathNormalizing:        true,
	// increase DNS cache time to an hour instead of default minute
	Dial: (&fasthttp.TCPDialer{
		Concurrency:      4096,
		DNSCacheDuration: time.Hour,
	}).Dial,
}

func FetchWeatherData(city, infoType string) (string, error) {
	// 1. Create a new request
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// 2. Set the URL and method
	req.SetRequestURI(API + fmt.Sprintf("?key=%s&city=%s&extensions=%s", KEY, city, infoType))
	req.Header.SetMethod("GET")

	// 3. Create a response object
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 4. Perform the request
	if err := httpclient.Do(req, resp); err != nil {
		return "", err
	}

	// 5. Get the response body
	body := resp.Body()

	return string(body), nil
}
