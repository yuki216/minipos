package response

import (
	"go-hexagonal-auth/api/paginator"
	"go-hexagonal-auth/business/report"
)

type getAllOmzetResponse struct {
	Meta  paginator.Meta    `json:"meta"`
	Omzets []GetOmzetResponse `json:"omzet"`
}

type GetOmzetResponse struct {
	Date   				 string    `json:"date"`
	MerchantName         string       `jsom:"merchant_name"`
	Total       		 float64   `json:"total";`
}

//NewGetAllUserResponse construct GetAllUserResponse
func NewGetOmzetResponse(omzets []report.Omzet, page int, rowPerPage int) getAllOmzetResponse {

	var (
		lenOmzets = len(omzets)
	)

	getAllOmzetResponse := getAllOmzetResponse{}
	getAllOmzetResponse.Meta.BuildMeta(lenOmzets, page, rowPerPage)

	for index, value := range omzets {
		if index == getAllOmzetResponse.Meta.RowPerPage {
			continue
		}

		var getOmzetResponse GetOmzetResponse

		getOmzetResponse.Date = value.Date
		getOmzetResponse.MerchantName = value.MerchantName
		getOmzetResponse.Total = value.Total

		getAllOmzetResponse.Omzets = append(getAllOmzetResponse.Omzets, getOmzetResponse)
	}

	if getAllOmzetResponse.Omzets== nil {
		getAllOmzetResponse.Omzets = []GetOmzetResponse{}
	}

	return getAllOmzetResponse
}
