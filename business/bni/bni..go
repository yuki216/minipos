package bni

//User product User that available to rent or sell
type Content struct {
	Status string
	Message string
	Data string
}

type Data struct {
	ClientId string
	TrxAmount string
	CustomerName string
	CustomerEmail string
	CustomerPhone string
	VirtualAccount string
	TrxId string
	DatetimeExpired string
	Description string
	Type string
}

type CreateBilling struct {
	ClientID string
	TrxAmount string
	CustomerName string
	CustomerEmail string
	CustomerPhone string
	VirtualAccount string
	TrxID string
	DatetimeExpired string
	Description string
	Type string
}

type CreateResponse struct {
	VirtualAccount string
	TrxID string
}