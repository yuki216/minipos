package bni

import (
	"encoding/json"
	"fmt"
	"go-hexagonal-auth/api/v1/bni/request"
	"go-hexagonal-auth/config"
	"go-hexagonal-auth/util/bniEnc"
)

//=============== The implementation of those interface put below =======================
type service struct {
	cfg        *config.Config
	repository Repository
}

//NewService Construct user service object
func NewService( cfg *config.Config, repository Repository) Service {
	return &service{
		cfg,
		repository,
	}
}


//Login by given user Username and Password, return error if not exist
func (s *service) CreateBilling(data *request.CreateBilling) ( *CreateResponse, error) {

	cid := s.cfg.BNIConfig.Cid// from BNI
	sck := s.cfg.BNIConfig.Key // from BNI


	jsonData := fmt.Sprintf(`{
	 "type" : "createbilling",
	 "client_id" : "%s",
	 "trx_id" : "%s",
	 "trx_amount" : "%s",
	 "billing_type" : "c",
	 "customer_name" : "%s",
	 "customer_email" : "%s",
	 "customer_phone" : "%s",
	 "virtual_account" : "%s",
	 "datetime_expired" : "%s",
	 "description" : "%s"+
	"}`, data.ClientID,data.TrxID, data.TrxAmount, data.CustomerName, data.CustomerEmail, data.CustomerPhone, data.VirtualAccount, data.DatetimeExpired, data.Description)

	str := bniEnc.Encrypt(jsonData, cid, sck)
	content, err := s.repository.Billing(str)
	if err != nil {
		return nil,err
	}

	decStr,err := bniEnc.Decrypt(content.Message, cid, sck)
	if err != nil {
		return nil, err
	}
	var res CreateResponse
	err = json.Unmarshal([]byte(decStr), res)
	if err != nil {
		return nil, err
	}

	if content.Status != "000"{
		return nil, err
	}

	return &res,nil
}


func (s *service) UpdateBilling(data Data) (*Content, error) {

	cid := s.cfg.BNIConfig.Cid// from BNI
	sck := s.cfg.BNIConfig.Key // from BNI

	json :=fmt.Sprintf(`{
	 "client_id" : "%s",
	 "trx_id" : "$%s,
	 "trx_amount" : "%s",
	 "customer_name" : "%s",
	 "customer_email" : "%s",
	 "customer_phone" : "%s",
	 "datetime_expired" : "%s",
	 "description" : "%s",
	 "type" : "updateBilling"
	}`, data.ClientId,data.TrxId, data.TrxAmount, data.CustomerName, data.CustomerEmail, data.CustomerPhone, data.DatetimeExpired, data.Description)

	str := bniEnc.Encrypt(json, cid, sck)
	content, err := s.repository.Billing(str)
	if err != nil {
		return nil,err
	}

	return content, nil
}

func (s *service) InquiryBilling(data Data) (*Content, error) {

	cid := s.cfg.BNIConfig.Cid// from BNI
	sck := s.cfg.BNIConfig.Key // from BNI

	json :=fmt.Sprintf(`{
			 "type" : "inquirybilling",
			 "client_id" : "%s",
			 "trx_id" : "%s",
			}`, data.ClientId,data.TrxId)

	str := bniEnc.Encrypt(json, cid, sck)
	content, err := s.repository.Billing(str)
	if err != nil {
		return nil,err
	}

	return content, nil
}