package client

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	netUrl "net/url"
)

type Client struct {
	Client *http.Client
}

func (r *Client) UrlQueryEncode(text string) string {
	return netUrl.QueryEscape(text)
}

func (r *Client) BodyToString(body io.ReadCloser) (string, error) {
	var bodyStr string
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		return bodyStr, err
	}

	return string(bodyByte), nil
}

func (r *Client) NewGetRequest(url string, headers map[string]string) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return request, nil
}

func (r *Client) GetTranslateTextWithSplitBody(body string) (string, error) {
	data := fmt.Sprintf("{\"data\": %s}", body)
	var dict map[string]interface{}
    err := json.Unmarshal([]byte(data), &dict)
    if err != nil {
        return data, err
    }
	if dict["data"].([]interface{})[0] != nil {
		return dict["data"].([]interface{})[0].([]interface{})[0].([]interface{})[0].(string), nil
	}
	return "nil", nil
} 
