package define

import "errors"

var (
    SystemError = errors.New("system error")
    GenerateApiError = errors.New("generate api error")
    RedirectApiError = errors.New("redirect api error")
)
