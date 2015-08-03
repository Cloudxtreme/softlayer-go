package slapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SoftLayerAPIError struct {
	Code   string `json:"code"`
	String string `json:"error"`
}

func (err SoftLayerAPIError) Error() string {
	return fmt.Sprintf("(%s) %s", err.Code, err.String)
}

type Client struct {
	Endpoint string
	Username string
	APIKey   string
}

func (client *Client) Context() RequestContext {
	return RequestContext{Client: client}
}

func (ctx *RequestContext) Call(result interface{}) error {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", ctx.Client.Endpoint+ctx.Service+"/"+ctx.Method, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(ctx.Client.Username, ctx.Client.APIKey)
	req.Header.Add("Content-Type", "application/json")
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
		paramBytes, err := json.Marshal(ctx.Arguments)
		if err != nil {
			return err
		}
		body["parameters"] = string(paramBytes)
	}

	if len(body) > 0 {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	}

	req.URL.RawQuery = values.Encode()
	log.Println(req.URL.String())

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

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

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return err
	}

	return nil
}

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
