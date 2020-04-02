package handler

import (
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net/http"
    "strings"
    "token-query/handler/transaction"
    "token-query/logger"
    "token-query/types"
    "token-query/utils"
)

func DoTxQuery(c *gin.Context) {
    var txQuery types.TransactionQueryRequest
    if err := c.ShouldBindJSON(&txQuery); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := txQuery.IsValid(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    switch strings.ToUpper(txQuery.Symbol) {
    case utils.EOS:
        txDetail, err := transaction.EOSTxQueryWithGreymass(txQuery.Tx)
        if err != nil {
            logger.Log.Error("EOSTxQueryWithGreymass error: %s", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        queryData, err := transaction.NewTxQueryResponseWithGreymass(txDetail)
        if err != nil {
            logger.Log.Error("NewTxQueryResponseWithGreymass error: %s", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, queryData)
        return

    case utils.ETH:
        txDetail, err := transaction.ETHTxQueryWithEthBtc(txQuery.Tx)
        if err != nil {
            logger.Log.Error("ETHTxQueryWithEthBtc error: %s", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        queryData, err := transaction.NewTxQueryResponseWithEthBtc(txDetail)
        if err != nil {
            logger.Log.Error("NewTxQueryResponseWithEthBtc error: %s", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, queryData)
        return

    case utils.BTC:
        txDetail, err := transaction.BTCTxQueryWithBlockChain(txQuery.Tx)
        if err != nil {
            logger.Log.Error("BTCTxQueryWithBlockChain error: %s", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        queryData, err := transaction.NewTxQueryResponseWithBlockChain(txDetail)
        if err != nil {
            logger.Log.Error("NewTxQueryResponseWithBlockChain error: %s", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, queryData)
        return

    case utils.USDT:
        if len(txQuery.Tx) < 2 {
            err := fmt.Errorf("tx '%s' invalid", txQuery.Tx)
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        switch strings.ToUpper(txQuery.Tx[0:2]) {
        case "0X":
            txDetail, err := transaction.USDTBaseEthTxQueryWithEthBtc(txQuery.Tx)
            if err != nil {
                logger.Log.Error("USDTBaseEthTxQueryWithEthBtc error: %s", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            queryData, err := transaction.NewUSDTTxQueryResponseWithEthBtc(txDetail)
            if err != nil {
                logger.Log.Error("NewUSDTTxQueryResponseWithEthBtc error: %s", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            c.JSON(http.StatusOK, queryData)
            return
        }

        err := fmt.Errorf("tx '%s' unsupport", txQuery.Tx)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
}

func DoTxVerify(c *gin.Context) {
    var txVerify types.TransactionVerifyRequest

    originData, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        logger.Log.Error("read request body error: %s", err)
        c.JSON(http.StatusBadRequest, types.TransactionVerifyResponse{
            TransactionVerifyRequest: txVerify,
            VerifyCode:               40001,
            VerifyMsg:                err.Error(),
        })
        return
    }

    logger.Log.Debug("request origin data: %s", string(originData))

    if err := json.Unmarshal(originData, &txVerify); err != nil {
        logger.Log.Error("unmarshal request data error: %s", err)
        c.JSON(http.StatusBadRequest, types.TransactionVerifyResponse{
            TransactionVerifyRequest: txVerify,
            VerifyCode:               40001,
            VerifyMsg:                err.Error(),
        })
        return
    }

    //if err := c.ShouldBindJSON(&txVerify); err != nil {
    //    c.JSON(http.StatusBadRequest, types.TransactionVerifyResponse{
    //        TransactionVerifyRequest: txVerify,
    //        VerifyCode:               40001,
    //        VerifyMsg:                err.Error(),
    //    })
    //    return
    //}

    if err := txVerify.IsValid(); err != nil {
        c.JSON(http.StatusBadRequest, types.TransactionVerifyResponse{
            TransactionVerifyRequest: txVerify,
            VerifyCode:               40002,
            VerifyMsg:                err.Error(),
        })
        return
    }

    switch strings.ToUpper(txVerify.Symbol) {
    case utils.ETH:
        txDetail, err := transaction.ETHTxQueryWithEthBtc(txVerify.Tx)
        if err != nil {
            logger.Log.Error("ETHTxQueryWithEthBtc error: %s", err)
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               50001,
                VerifyMsg:                err.Error(),
            })
            return
        }
        verifyCode, verifyErr := transaction.NewETHTxVerifyResponseWithEthBtc(txDetail, &txVerify)
        if verifyErr != nil {
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               verifyCode,
                VerifyMsg:                verifyErr.Error(),
            })
            return
        } else {
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               verifyCode,
                VerifyMsg:                "verify pass",
            })
            return
        }
    case utils.USDT:
        if len(txVerify.Tx) < 2 {
            err := fmt.Errorf("tx '%s' invalid", txVerify.Tx)
            c.JSON(http.StatusBadRequest, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               40003,
                VerifyMsg:                err.Error(),
            })
            return
        }

        switch strings.ToUpper(txVerify.Tx[0:2]) {
        case "0X":
            txDetail, err := transaction.USDTBaseEthTxQueryWithEthBtc(txVerify.Tx)
            if err != nil {
                logger.Log.Error("USDTBaseEthTxQueryWithEthBtc error: %s", err)

                if err == transaction.ErrNotFind || err == transaction.ErrTxType {
                    c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                        TransactionVerifyRequest: txVerify,
                        VerifyCode:               40004,
                        VerifyMsg:                err.Error(),
                    })
                    return
                }

                c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                    TransactionVerifyRequest: txVerify,
                    VerifyCode:               50001,
                    VerifyMsg:                err.Error(),
                })
                return
            }
            verifyCode, verifyErr := transaction.NewUSDTTxVerifyResponseWithEthBtc(txDetail, &txVerify)
            if verifyErr != nil {
                c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                    TransactionVerifyRequest: txVerify,
                    VerifyCode:               verifyCode,
                    VerifyMsg:                verifyErr.Error(),
                })
                return
            } else {
                c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                    TransactionVerifyRequest: txVerify,
                    VerifyCode:               verifyCode,
                    VerifyMsg:                "verify pass",
                })
                return
            }
        }

        err := fmt.Errorf("tx '%s' unsupport", txVerify.Tx)
        c.JSON(http.StatusBadRequest, types.TransactionVerifyResponse{
            TransactionVerifyRequest: txVerify,
            VerifyCode:               40005,
            VerifyMsg:                err.Error(),
        })
        return

    case utils.BTC:
        txDetail, err := transaction.BTCTxQueryWithBlockChain(txVerify.Tx)
        if err != nil {
            logger.Log.Error("BTCTxQueryWithBlockChain error: %s", err)
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               50001,
                VerifyMsg:                err.Error(),
            })
            return
        }
        verifyCode, verifyErr := transaction.NewBTCTxVerifyResponseWithBlockChain(txDetail, &txVerify)
        if verifyErr != nil {
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               verifyCode,
                VerifyMsg:                verifyErr.Error(),
            })
            return
        } else {
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               verifyCode,
                VerifyMsg:                "verify pass",
            })
            return
        }
    case utils.EOS:
        txDetail, err := transaction.EOSTxQueryWithGreymass(txVerify.Tx)
        if err != nil {
            logger.Log.Error("EOSTxQueryWithGreymass error: %s", err)
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               50001,
                VerifyMsg:                err.Error(),
            })
            return
        }
        verifyCode, verifyErr := transaction.NewEOSTxVerifyResponseWithGreymass(txDetail, &txVerify)
        if verifyErr != nil {
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               verifyCode,
                VerifyMsg:                verifyErr.Error(),
            })
            return
        } else {
            c.JSON(http.StatusOK, types.TransactionVerifyResponse{
                TransactionVerifyRequest: txVerify,
                VerifyCode:               verifyCode,
                VerifyMsg:                "verify pass",
            })
            return
        }
    }
}
