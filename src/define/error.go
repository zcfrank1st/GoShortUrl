package define

import "errors"

var (
    SystemError = errors.New("system error")
    StoreError = errors.New("store url error")
    GetError = errors.New("get short url error")
)
