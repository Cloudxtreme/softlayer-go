package slapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

// SoftLayerAPIError is an error returned when there's an error returned from the SoftLayer API
type SoftLayerAPIError struct {
	Code   string `json:"code"`
	String string `json:"error"`
}

// Error shows a nice error string
func (err SoftLayerAPIError) Error() string {
	return fmt.Sprintf("(%s) %s", err.Code, err.String)
}

// Client is the SoftLayer API Client
type Client struct {
	Endpoint   string
	Username   string
	APIKey     string
	HTTPClient *http.Client
	Debug      bool
}

// Call performs the API call. The result parameter will have the API response marshalled into it
func (client *Client) Call(req Request, result interface{}) error {
	if req.Service == "" || req.Method == "" {
		return errors.New("service and method is required to make an API call")
	}

	if client.HTTPClient == nil {
		client.HTTPClient = http.DefaultClient
	}
	httpClient := client.HTTPClient

	method := "GET"
	if req.Method == "deleteObject" {
		method = "DELETE"
	} else if req.Method == "createObject" || req.Method == "createObjects" {
		method = "POST"
	} else if req.Method == "editObject" || req.Method == "editObjects" {
		method = "PUT"
	} else if req.Arguments != nil && len(req.Arguments) > 0 {
		method = "POST"
	}

	urlParts := []string{client.Endpoint, req.Service}
	if req.ID != nil {
		bytes, err := json.Marshal(req.ID)
		if err != nil {
			return err
		}

		urlParts = append(urlParts, string(bytes))
	}

	urlParts = append(urlParts, req.Method)

	httpReq, err := http.NewRequest(method, strings.Join(urlParts, "/"), nil)
	if err != nil {
		return err
	}

	httpReq.SetBasicAuth(client.Username, client.APIKey)
	httpReq.Header.Add("Content-Type", "application/json")

	// Set object mask, filter, limit, offset, arguments settings
	values := httpReq.URL.Query()
	body := make(map[string]interface{})

	if req.Mask != "" {
		values.Set("objectMask", req.Mask)
	}

	if req.Limit != 0 {
		values.Set("resultLimit", fmt.Sprintf("%d,%d", req.Offset, req.Limit))
	}

	if req.Filter != nil {
		filterBytes, err := json.Marshal(req.Filter)
		if err != nil {
			return err
		}
		values.Set("objectFilter", string(filterBytes))
	}

	if req.Arguments != nil {
		body["parameters"] = req.Arguments
	}

	if len(body) > 0 {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		httpReq.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
		method = "POST"
	}

	httpReq.URL.RawQuery = values.Encode()

	if client.Debug == true {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err != nil {
			log.Println(string(dump))
		}
	}

	// Make the API Request
	res, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	// See if this response is an error
	// note(kmcdonald): We have to do this since there are some errors which return with 200 status codes. This, I feel, is a bug in the REST API.
	slError := SoftLayerAPIError{}
	err = json.Unmarshal(respBody, &slError)
	if err != nil && res.StatusCode < 200 && res.StatusCode >= 300 {
		// We could not parse an error, but there is a non-2xx status code
		return SoftLayerAPIError{
			Code:   res.Status,
			String: res.Status,
		}
	} else if slError.Code != "" && slError.String != "" {
		// We could parse the error and we were able to parse the error code and error string.
		return slError
	}

	// Unmarshal Result
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return err
	}

	return nil
}

// Request holds request-level information
type Request struct {
	Service   string
	Method    string
	ID        interface{}
	Arguments []interface{}
	Mask      string
	Filter    interface{}
	Limit     int
	Offset    int
}
