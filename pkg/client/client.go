package client

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	netUrl "net/url"

	"github.com/nopeEOF/translator/pkg/config"
)

type Client struct {
	Client *http.Client
}

func (c *Client) UrlQueryEncode(text string) string {
	return netUrl.QueryEscape(text)
}

func (c *Client) BodyToString(body io.ReadCloser) (string, error) {
	var bodyStr string
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		return bodyStr, err
	}

	return string(bodyByte), nil
}

func (c *Client) NewGetRequest(url string, headers map[string]string) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return request, nil
}

func (c *Client) GetTranslateTextWithSplitBody(body string) (string, error) {
	data := fmt.Sprintf("{\"data\": %s}", body)
	var dict map[string]interface{}
	err := json.Unmarshal([]byte(data), &dict)
	if err != nil {
		return data, err
	}
	if dict["data"].([]interface{})[0] != nil {
		var value string
		for _, d := range dict["data"].([]interface{})[0].([]interface{}) {
			value = value + d.([]interface{})[0].(string)
		}
		return value, nil
	}
	return "nil", nil
}

func (c *Client) KdialogMessageBody(body string, config config.Config) string {
	body = html.EscapeString(body)
	fmt.Println(config.FontSize)
	return fmt.Sprintf("<html><font size='%d'><body dir='%s'><p align=\"justify\">%s</p></body></font></html>", config.FontSize, config.Direction, body)
}
