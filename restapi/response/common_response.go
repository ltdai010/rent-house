package response

type ResponseBool struct {
	Data bool `json:"data" xml:"data"`
	*Err
}

type ResponseCommonSingle struct {
	Data interface{} `json:"data" xml:"data"`
	*Err
}

type ResponseCommonSingleWithValidate struct {
	Data interface{} `json:"data" xml:"data"`
	*ValidateErr
}

type ResponseCommonArray struct {
	Data       interface{} `json:"data" xml:"data"`
	TotalCount int64       `json:"total_count" xml:"total_count"`
	*Err
}

type ResponseCommonArrayPtr struct {
	Data       []*interface{} `json:"data" xml:"data"`
	TotalCount int64          `json:"total_count" xml:"total_count"`
	*Err
}
