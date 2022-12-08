package services

import (
	"io"
	"net/http"
	"time"

	breaker "github.com/sony/gobreaker"
)

var (
	CbSetting = breaker.Settings{
		Name:        "Ticket Get",
		MaxRequests: 0,
		Interval:    10 * time.Second,
		Timeout:     3 * time.Second,
		ReadyToTrip: func(counts breaker.Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
		OnStateChange: nil,
		IsSuccessful:  nil,
	}

	bonusCb  = breaker.NewCircuitBreaker(CbSetting)
	ticketCb = breaker.NewCircuitBreaker(CbSetting)
	flightCb = breaker.NewCircuitBreaker(CbSetting)
)

var cb *breaker.CircuitBreaker

func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	//// we need to buffer the body if we want to read it here and send it
	//// in the request.
	//body, err := io.ReadAll(req.Body)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//// you can reassign the body if you need to parse it as multipart
	//req.Body = io.NopCloser(bytes.NewReader(body))
	//
	//// create a new url from the raw RequestURI sent by the client
	//url := fmt.Sprintf("%s://%s%s", proxyScheme, proxyHost, req.RequestURI)
	//
	//proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(body))
	//
	//// We may want to filter some headers, otherwise we could just use a shallow copy
	//// proxyReq.Header = req.Header
	//proxyReq.Header = make(http.Header)
	//for h, val := range req.Header {
	//	proxyReq.Header[h] = val
	//}
	//
	//resp, err := httpClient.Do(proxyReq)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadGateway)
	//	return
	//}
	//defer resp.Body.Close()

	// legacy

	//r.URL.Host = "example.com"
	//r.RequestURI = ""
	//client := &http.Client{}
	//
	//delete(r.Header, "Accept-Encoding")
	//delete(r.Headers, "Content-Length")
	//resp, err := client.Do(r.WithContext(context.Background())
	//if err != nil {
	//	return nil, err
	//}
	//return resp, nil
}
