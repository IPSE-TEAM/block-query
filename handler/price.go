package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "token-query/types"
)

var priceQueryResult *types.PriceQueryResponse

func InitPriceService(pqr *types.PriceQueryResponse) error {
    if pqr == nil {
        return fmt.Errorf("parameter 'pqr' is nil")
    }

    priceQueryResult = pqr

    return nil
}

func DoPriceQuery(c *gin.Context) {
    if priceQueryResult == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "no data available"})
        return
    }

    c.JSON(http.StatusOK, priceQueryResult)
    return
}
