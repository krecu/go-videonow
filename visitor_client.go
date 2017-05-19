package videonow

import (
	"net/http"
	"time"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"stats/videonow/model"
)

type VisitorClient struct {
	Host string
	Port string
}

type VisitorBody struct {
	Ip string
	Ua string
	Id string
}

func NewVisitorClient(host string, port string) (*VisitorClient) {

	return &VisitorClient{
		Host: host,
		Port: port,
	}
}

func (v *VisitorClient) Get(id string, ip string, ua string) (*model.Visitor, error){
	var proto model.Visitor

	url := v.Host + ":" + v.Port + "/api/visitor"

	body, err := json.Marshal(&VisitorBody{
		Ip: ip,
		Id: id,
		Ua: ua,
	})

	if err != nil {
		return &model.Visitor{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := client.Do(req)
	defer response.Body.Close()

	if err != nil || response.StatusCode != 200 {
		return &model.Visitor{}, err
	}

	result, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(result, &proto)

	return &proto, err
}