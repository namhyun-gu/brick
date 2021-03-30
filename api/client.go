package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	http *http.Client
}

type HTTPError struct {
	StatusCode int
	RequestURL *url.URL
}

func (err HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d (%s)", err.StatusCode, err.RequestURL)
}

type ClientOption = func(http.RoundTripper) http.RoundTripper

func NewHTTPClient(opts ...ClientOption) *http.Client {
	tr := http.DefaultTransport
	for _, opt := range opts {
		tr = opt(tr)
	}
	return &http.Client{Transport: tr}
}

func NewClient(opts ...ClientOption) *Client {
	client := &Client{NewHTTPClient(opts...)}
	return client
}

func NewClientFromHTTP(httpClient *http.Client) *Client {
	client := &Client{http: httpClient}
	return client
}

func (c Client) REST(hostname string, method string, p string, body io.Reader, data interface{}) error {
	reqUrl := hostname + p
	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !success {
		return HTTPError{
			StatusCode: resp.StatusCode,
			RequestURL: resp.Request.URL,
		}
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) GET(hostname string, p string) ([]byte, error) {
	reqUrl := hostname + p

	resp, err := c.http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !success {
		return nil, HTTPError{
			StatusCode: resp.StatusCode,
			RequestURL: resp.Request.URL,
		}
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
