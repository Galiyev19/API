package helpers

type BaseResponse struct {
	Result  any  `json:"result"`
	Success bool `json:"success"`
}

func GenerateResponse(result any, success bool) *BaseResponse {
	return &BaseResponse{
		Result:  result,
		Success: success,
	}
}
