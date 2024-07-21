package translate

import (
	"fmt"

	"github.com/nopeEOF/translator/pkg/client"
	"github.com/nopeEOF/translator/pkg/config"
)
func Translate(text string, config config.Config, client client.Client) (string, error) {
	var translate string
	text = client.UrlQueryEncode(text)
	url := fmt.Sprintf(config.Url, config.Lang, text)
	header := map[string]string{
		"User-Agent":"Mozilla/5.0",
	}
	getRequest, err := client.NewGetRequest(url, header)
	if err != nil {
		return translate, err
	}
	response, err := client.Client.Do(getRequest)
	if err != nil {
		return translate, err
	}
	defer response.Body.Close()
	body, err := client.BodyToString(response.Body)
	if err != nil {
		return translate, err
	}
	return body, nil
}