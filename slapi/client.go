package slapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// Request creates a new request. Specify request-specific settings on the return value
func (client *Client) Request() RequestContext {
	return RequestContext{Client: client}
}

// RequestContext holds request-level information
type RequestContext struct {
	Service   string
	Method    string
	ID        interface{}
	Arguments []interface{}
	Mask      string
	Filter    interface{}
	Limit     int
	Offset    int
	Client    *Client
}

// Do performs the API call. The result parameter will have the API response marshalled into it
func (ctx *RequestContext) Do(result interface{}) error {
	// TODO: the http client should be re-used
	httpClient := &http.Client{}

	method := "GET"
	if ctx.Method == "deleteObject" {
		method = "DELETE"
	} else if ctx.Method == "createObject" || ctx.Method == "createObjects" {
		method = "POST"
	} else if ctx.Method == "editObject" || ctx.Method == "editObjects" {
		method = "PUT"
	} else if ctx.Arguments != nil && len(ctx.Arguments) > 0 {
		method = "POST"
	}

	req, err := http.NewRequest(method, ctx.Client.Endpoint+ctx.Service+"/"+ctx.Method, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(ctx.Client.Username, ctx.Client.APIKey)
	req.Header.Add("Content-Type", "application/json")

	// Set object mask, filter, limit, offset, arguments settings
	values := req.URL.Query()
	body := make(map[string]interface{})

	if ctx.Mask != "" {
		values.Set("objectMask", string(ctx.Mask))
	}

	if ctx.Limit != 0 {
		values.Set("limit", string(ctx.Limit))
	}

	if ctx.Offset != 0 {
		values.Set("offset", string(ctx.Offset))
	}

	if ctx.Filter != nil {
		filterBytes, err := json.Marshal(ctx.Filter)
		if err != nil {
			return err
		}
		values.Set("objectFilter", string(filterBytes))
	}

	if ctx.Arguments != nil {
		body["parameters"] = ctx.Arguments
	}

	if len(body) > 0 {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		log.Println(string(bodyBytes))
		req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	}

	req.URL.RawQuery = values.Encode()

	// TODO: Remove this log (add debug option?)
	log.Println(req.URL.String())

	// Make the API Request
	res, err := httpClient.Do(req)
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

	log.Println("RESPONSE", string(respBody))

	// Unmarshal Result
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return err
	}

	return nil
}
