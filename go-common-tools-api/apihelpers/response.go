package apihelpers

// RandString 产生随机字符串，可用于密码等
func ResultSuccess(param map[string]string) *JsonResponse {
	return ResultJson(200, "success", param)
}

func ResultCode(code int, msg string) *JsonResponse {
	return ResultJson(code, msg, make(map[string]string))
}

type JsonResponse struct {
	//必须的大写开头
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"result"` //key重命名,最外面是反引号
}

func ResultJson(code int, msg string, param map[string]string) *JsonResponse {
	var jsonResponse JsonResponse

	jsonResponse.Code = code
	jsonResponse.Msg = msg
	jsonResponse.Data = param

	return &jsonResponse
}
