package service

import (
    "github.com/gin-gonic/gin"
)

func GenerateShort(context *gin.Context) {
    // todo
    // 存数据库(id, url), 获取last insert id
    // return To62
}

func RedirectShort(context *gin.Context) {
    // todo
    // get To62 string
    // transform to decimal
    // get url through decimal
    // redirect url
}