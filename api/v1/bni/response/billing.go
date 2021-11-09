package response


//Login response payload
type CreateResponse struct {
	VirtualAccount string `json:"virtual_account" `
	TrxID string `json:"trx_id"`
}

//NewMediaResponse construct MediaResponse
func NewCreateResponse(content *CreateResponse) *CreateResponse {
	var CreateResponse CreateResponse


	return &CreateResponse
}
