package slapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Endpoint string
	Username string
	APIKey   string
}

// Call performs the API call. The result parameter will have the API response marshalled into it
func (client *Client) Call(req Request, result interface{}) error {
	if req.Service == "" || req.Method == "" {
		return errors.New("service and method is required to make an API call")
	}

	// TODO: the http client should be re-used
	httpClient := &http.Client{}

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

	httpReq, err := http.NewRequest(method, client.Endpoint+req.Service+"/"+req.Method, nil)
	if err != nil {
		return err
	}

	httpReq.SetBasicAuth(client.Username, client.APIKey)
	httpReq.Header.Add("Content-Type", "application/json")

	// Set object mask, filter, limit, offset, arguments settings
	values := httpReq.URL.Query()
	body := make(map[string]interface{})

	if req.Mask != "" {
		values.Set("objectMask", string(req.Mask))
	}

	if req.Limit != 0 {
		values.Set("limit", string(req.Limit))
	}

	if req.Offset != 0 {
		values.Set("offset", string(req.Offset))
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
	}

	httpReq.URL.RawQuery = values.Encode()

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

	// Return error on incorrect status
	if res.StatusCode != 200 {
		slError := SoftLayerAPIError{}
		err = json.Unmarshal(respBody, &slError)
		if err != nil {
			return SoftLayerAPIError{
				String: res.Status,
				Code:   res.Status,
			}
		}
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
