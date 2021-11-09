package bni

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-hexagonal-auth/business/bni"
	"go-hexagonal-auth/config"
	"go-hexagonal-auth/modules/models"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

const(

)
type Object struct {
	CID			string
	ClientID    string
	Key         string
	ContentType string
}


type Configuration struct {
	cfg *config.BNIConfig
	db  *gorm.DB
}

func NewBNIConfiguration(cfg *config.BNIConfig, db *gorm.DB) *Configuration {
	return &Configuration{
		cfg: cfg,
		db:db,
	}
}


func (s *Configuration) Billing(data string) (*bni.Content, error) {
	data, err:= s.executionRequest(data)
	if err != nil {
		return nil, err
	}
	var content *bni.Content
 	err = json.Unmarshal([]byte(data), content)
	if err != nil {
		return nil, err
	}
	if content.Status !="000"{
		return  nil, errors.New("error:"+content.Message)
	}
	return content,nil
}

func (s *Configuration) UpdateTransaction(data string) (*string, error) {
	data, err:= s.executionRequest(data)
	if err != nil {
		return nil, err
	}
	return &data,nil
}

func (s *Configuration) InquiryBilling(data string) (*string, error) {
	data, err:= s.executionRequest(data)
	if err != nil {
		return nil, err
	}
	return &data,nil
}

func (s *Configuration) executionRequest(content string) (string, error) {
	httpPostUrl := s.cfg.Url

	data := models.Content{
		ClientID: s.cfg.ClientID,
		Prefix:   s.cfg.Cid,
		Data:     content,
	}
	var jsonData = []byte( fmt.Sprintf("%v",data))
	request, error := http.NewRequest("POST", httpPostUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return "",error
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body),nil
}