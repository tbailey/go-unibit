package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"
)
const (
	// APIEndpoint the api url
	APIEndpoint = "https://api.unibit.ai"

	// APIVersion the api version
	APIVersion = "v2"
)

var (
	// ErrEmptyResponse if the api returns an empty response
	ErrEmptyResponse = errors.New("empty response given - you may of entered an incorrect symbol")

	// ErrTooManyRequests if the api returns that you have made too many requests
	ErrTooManyRequests = errors.New("you are over the request limit - you may of not entered a valid token")

	// ErrTickerNotFound the err for when the api returns that the ticker supplied doesn't exist
	ErrTickerNotFound = errors.New("ticker not found")

	// ErrServer the error for when the api returns a status code implying a server failure
	ErrServer = errors.New("server err")

	// APIErrTickerNotFound the return body if the ticker doesn't exist
	APIErrTickerNotFound = "Ticker Not Found."
)

type Result struct {
	MetaData MetaData `json:"meta_data"`
	ResultData json.RawMessage `json:"result_data"`
}

type MetaData struct {
	ApiName string `json:"api_name"`
	TotalDataPoints uint64 `json:"num_total_data_points"`
	CreditCost uint64 `json:"credit_cost"`
}

// Client is the data structure for holding the Client details
type Client struct {
	Key           string
	Client        http.Client

	endpoint string
}

// NewClient returns a new api client
func NewClient(key string) *Client {
	a := &Client{Key: key, endpoint: APIEndpoint}
	a.Client = http.Client{
		Timeout: time.Second * 30,
	}
	return a
}

// Get requests a get from the api
func (a *Client) Get(path string, params URLParams, response interface{}) error {
	return a.Call(http.MethodGet, path, params, response)
}

// Call calls the api using a supplied method
func (a *Client) Call(method string, path string, params URLParams, response interface{}) error {
	params[ParamToken] = a.Key
	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}

	endpoint := fmt.Sprintf("%v/%v/%v", a.endpoint, APIVersion, path)
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = q.Encode()
	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 400 {
		if resp.StatusCode == http.StatusTooManyRequests {
			return ErrTooManyRequests
		}
		return ErrServer
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) == APIErrTickerNotFound {
		return ErrTickerNotFound
	}
	r := Result{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(r.ResultData, &response)
	if err != nil {
		return err
	}

	// checking if response is empty as api doesn't return 404
	if isEmptyResponse(response) {
		return ErrEmptyResponse
	}

	return nil
}

func isEmptyResponse(response interface{}) bool {
	return reflect.DeepEqual(response, getEmptyResponse(response))
}

func getEmptyResponse(p interface{}) interface{} {
	v := reflect.ValueOf(p)
	if v.Kind() == reflect.Ptr {
		return reflect.New(v.Elem().Type()).Interface()
	}
	return reflect.Zero(v.Type()).Interface()
}