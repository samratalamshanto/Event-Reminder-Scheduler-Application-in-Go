package base

type CommonResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	MSG     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Err     interface{} `json:"err"`
}
