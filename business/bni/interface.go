package bni

import "go-hexagonal-auth/api/v1/bni/request"

//Service outgoing port for user
type Service interface {
	//Login If data not found will return nil without error
	CreateBilling(data *request.CreateBilling) ( *CreateResponse, error)
}
type Repository interface {
	Billing(data string) (*Content, error)
	UpdateTransaction(data string) (*string, error)
	InquiryBilling(data string) (*string, error)
}