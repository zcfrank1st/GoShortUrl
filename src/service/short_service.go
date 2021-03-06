package service

import (
    "github.com/gin-gonic/gin"
    "define"
)

var engine define.Engine // interface is just a pointer

func init () {
    engine = &define.MysqlEngine{}
}

type Url struct {
    UrlContent string `json:"url"`
}


func GenerateShort(context *gin.Context) {
    var urlJson Url
    if context.BindJSON(&urlJson) == nil {
        lastInsertId, err := engine.StoreUrl(urlJson.UrlContent)

        if err != nil {
            context.JSON(200, gin.H{
                "code": define.SystemErrorCode,
                "body": define.GenerateApiError,
            })
        } else {
            context.JSON(200, gin.H{
                "code": define.SuccessCode,
                "body": define.To62(lastInsertId),
            })
        }
    } else {
        context.JSON(200, gin.H{
            "code": define.SystemErrorCode,
            "body": define.SystemError,
        })
    }
}

func RedirectShort(context *gin.Context) {
    id := define.ToDecimal(context.Param("id"))

    url, err := engine.GetUrl(id)
    if err != nil {
        context.JSON(200, gin.H{
            "code": define.SystemErrorCode,
            "body": define.RedirectApiError,
        })
    } else {
        context.Redirect(302, url)
    }
}