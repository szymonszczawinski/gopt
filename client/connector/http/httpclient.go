package http

import (
	"encoding/json"
	"gopt/client/connector"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	instance := new(HttpClient)
	instance.client = &http.Client{}
	return instance
}
func (c *HttpClient) PerformRequest(method connector.HttpMethod, url string, payload any, callback connector.IConnectorCallback) {
	switch method {
	case connector.GET:
		c.performGet(url, callback)
	}
}

func (c *HttpClient) performGet(url string, callback connector.IConnectorCallback) {
	request, err := http.NewRequest(string(connector.GET), url, nil)
	if err != nil {
		log.Println(err.Error())
		callback.OnMessage(url, err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	response, err := c.client.Do(request)
	if err != nil {
		log.Println(err.Error())
		callback.OnMessage(url, err)
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		callback.OnMessage(url, err)

	}
	var responseObject map[string]any
	json.Unmarshal(bodyBytes, &responseObject)
	log.Printf("API Response as struct %+v\n", responseObject)

}
