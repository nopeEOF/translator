package client

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"html"
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
		var value string
		for _, d := range(dict["data"].([]interface{})[0].([]interface{})) {
			value = value + d.([]interface{})[0].(string)
		}
		return value, nil
	}
	return "nil", nil
} 

func (c *Client) KdialogMessageBody(body string) string {
	body =  html.EscapeString(body)
	return fmt.Sprintf("<html><body dir='rtl'><p align=\"justify\">%s</p></body></html>", body)
}
