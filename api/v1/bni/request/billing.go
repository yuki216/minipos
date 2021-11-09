package request

type CreateBilling struct {
	ClientID string `json:"client_id"`
	TrxAmount string `json:"trx_amount"`
	CustomerName string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
	CustomerPhone string `json:"customer_phone"`
	VirtualAccount string `json:"virtual_account"`
	TrxID string `json:"trx_id"`
	DatetimeExpired string `json:"datetime_expired"`
	Description string `json:"description"`
	Type string `json:"type"`
}

//{
//“type” : “createbilling”,
//“client_id” : “001”,
//“trx_id” : “1230000001”,
//“trx_amount” : “100000”,
//“billing_type” : “c”,
//“customer_name” : “Mr. X”,
//“customer_email” : “xxx@email.com”,
//“customer_phone” : “08123123123”,
//“virtual_account” : “8001000000000001”,
//“datetime_expired” : “2016-03-01T16:00:00+07:00”,
//“description” : “Payment of Trx 123000001”
//}
