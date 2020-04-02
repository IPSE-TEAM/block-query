package route

import (
    "github.com/gin-gonic/gin"
    "token-query/handler"
)

func Router() *gin.Engine {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    router.POST("/query/tx", handler.DoTxQuery)
    router.GET("/query/price", handler.DoPriceQuery)
    router.POST("/verify/tx", handler.DoTxVerify)

    return router
}
