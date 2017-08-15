package service

import (
    "github.com/gin-gonic/gin"
    "define"
)

var engine *define.MysqlEngine

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
                "code": define.SystemError,
                "body": nil,
            })
        } else {
            context.JSON(200, gin.H{
                "code": define.SuccessCode,
                "body": define.To62(lastInsertId),
            })
        }
    } else {
        context.JSON(200, gin.H{
            "code": define.SystemError,
            "body": nil,
        })
    }
}

func RedirectShort(context *gin.Context) {
    id := define.ToDecimal(context.Param("id"))

    url, err := engine.GetUrl(id)
    if err != nil {
        context.JSON(200, gin.H{
            "code": define.SystemError,
            "body": nil,
        })
    } else {
        context.Redirect(302, url)
    }
}