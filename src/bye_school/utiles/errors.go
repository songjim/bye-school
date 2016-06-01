package utiles

var errorCode = map[string]string{
	"10001": "连接超时",
}

// APIError 自定义错误
type APIError struct {
	Code string
}

func (e APIError) Error() string {
	return errorCode[e.Code]
}

// NewAPIError 返回指定Code的错误
func NewAPIError(code string) error {
	return APIError{code}
}