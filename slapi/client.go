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
	// Error code
	Code string `json:"code"`
	// Error string
	String string `json:"error"`
}

// Error shows a nice error string
func (err SoftLayerAPIError) Error() string {
	return fmt.Sprintf("(%s) %s", err.Code, err.String)
}

// Client is the SoftLayer API Client
type Client struct {
	// Endpoint is the URL to the SoftLayer API
	Endpoint string
	// Username is the SoftLayer Username
	Username string
	// APIKey is the SoftLayer API key
	APIKey string
	// HTTPClient is a golang HTTP client
	HTTPClient *http.Client
	// Debug is a flag which tells if request/response is logged
	Debug bool
}

// Request holds request-level information
type Request struct {
	// Service is the SoftLayer API service
	Service string
	// Method is the SoftLayer API method
	Method string
	// ID is the InitParameter for the request
	ID interface{}
	// Parameters are API parameters for the request
	Parameters []interface{}
	// Mask is the mask for the request
	Mask string
	// Filter is the filter for the request
	Filter interface{}
	// Limit is the result limit for the request
	Limit int
	// offset is the result offset for the request
	Offset int
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

	httpReq, err := client.makeHTTPRequest(req)
	if err != nil {
		return errors.New(fmt.Sprint("Error making request: %s", err))
	}

	if client.Debug == true {
		reqDump, err := httputil.DumpRequestOut(httpReq, true)
		if err == nil {
			log.Println(string(reqDump))
		}
	}

	// Make the API Request
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}

	if client.Debug == true {
		respDump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			log.Println(string(respDump))
		}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	// See if this response is an error
	// NOTE(kmcdonald): We have to do this since there are some errors which return with 200 status codes. This, I feel, is a bug in the REST API.
	slError := SoftLayerAPIError{}
	err = json.Unmarshal(respBody, &slError)
	if err != nil && resp.StatusCode < 200 && resp.StatusCode >= 300 {
		// We could not parse an error, but there is a non-2xx status code
		return SoftLayerAPIError{
			Code:   resp.Status,
			String: resp.Status,
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

func (client *Client) makeHTTPRequest(req Request) (*http.Request, error) {

	method := "GET"
	if req.Method == "deleteObject" {
		method = "DELETE"
	} else if req.Method == "createObject" || req.Method == "createObjects" {
		method = "POST"
	} else if req.Method == "editObject" || req.Method == "editObjects" {
		method = "PUT"
	} else if req.Parameters != nil && len(req.Parameters) > 0 {
		method = "POST"
	}

	urlParts := []string{client.Endpoint, req.Service}
	if req.ID != nil {
		bytes, err := json.Marshal(req.ID)
		if err != nil {
			return nil, err
		}

		urlParts = append(urlParts, string(bytes))
	}

	urlParts = append(urlParts, req.Method)

	httpReq, err := http.NewRequest(method, strings.Join(urlParts, "/"), nil)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(client.Username, client.APIKey)
	httpReq.Header.Add("Content-Type", "application/json")

	// Set object mask, filter, limit, offset, parameters settings
	values := httpReq.URL.Query()

	if req.Mask != "" {
		values.Set("objectMask", req.Mask)
	}

	if req.Limit != 0 {
		values.Set("resultLimit", fmt.Sprintf("%d,%d", req.Offset, req.Limit))
	}

	if req.Filter != nil {
		filterBytes, err := json.Marshal(req.Filter)
		if err != nil {
			return nil, err
		}
		values.Set("objectFilter", string(filterBytes))
	}

	body := make(map[string]interface{})
	if req.Parameters != nil {
		body["parameters"] = req.Parameters
	}

	if len(body) > 0 {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		httpReq.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
		method = "POST"
	}

	httpReq.URL.RawQuery = values.Encode()
	return httpReq, nil
}
