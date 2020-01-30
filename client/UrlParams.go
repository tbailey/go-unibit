package client

import "strings"

const (
	ParamToken = "accessKey"
)

type Request struct {
	Tickers []string
	StartDate string
	EndDate string
	SelectedFields []string
	DataType string
}

func (r Request) GetUrlParams() URLParams {
	params := URLParams{}
	params["tickers"] = strings.Join(r.Tickers,",")
	params["startDate"] = r.StartDate
	params["endDate"] = r.EndDate
	params["selectedFields"] = strings.Join(r.SelectedFields, ",")
	params["dataType"] = r.DataType

	return params
}

type URLParams map[string]string
