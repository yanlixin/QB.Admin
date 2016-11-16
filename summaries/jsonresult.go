package summaries

type JsonResult struct {
	Code    string        `json:"code"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

func JsonResultSuccess(msg string, data []interface{}) *JsonResult {
	result := new(JsonResult)
	result.Code = "0000"
	result.Status = "OK"
	result.Message = msg
	result.Data = data
	return result
}

func JsonResultUnmarshalFailed(msg string, data []interface{}) *JsonResult {
	result := new(JsonResult)
	result.Code = "0010"
	result.Status = "Fail"
	result.Message = msg
	result.Data = data
	return result
}

func JsonResultLogicFailed(msg string, data []interface{}) *JsonResult {
	result := new(JsonResult)
	result.Code = "0020"
	result.Status = "Failed"
	result.Message = msg
	result.Data = data
	return result
}
