package types

import (
    "fmt"
    "strings"
    "token-query/utils"
)

var SupportTxSymbols = []string{utils.EOS, utils.BTC, utils.USDT, utils.ETH}

type TransactionQueryRequest struct {
    Tx     string `json:"tx"`
    Symbol string `json:"symbol"`
}

func (tr *TransactionQueryRequest) IsValid() error {
    var valid = false
    var err error
    for _, symbol := range SupportTxSymbols {
        if symbol == strings.ToUpper(tr.Symbol) {
            valid = true
            break
        }
    }

    if !valid {
        err = fmt.Errorf("unsupport symbol: %s", tr.Symbol)
    }

    return err
}

var SupportPriceSymbols = []string{utils.EOS, utils.BTC, utils.ETH}

//type PriceQueryRequest struct {
//    Symbol string `json:"symbol"`
//}
//
//func (tr *PriceQueryRequest) IsValid() error {
//    var valid = false
//    var err error
//    for _, symbol := range SupportPriceSymbols {
//        if symbol == tr.Symbol {
//            valid = true
//            break
//        }
//    }
//
//    if !valid {
//        err = fmt.Errorf("unsupport symbol: %s", tr.Symbol)
//    }
//
//    return err
//}

type TransactionVerifyRequest struct {
    Tx        string `json:"tx"`
    Symbol    string `json:"symbol"`
    Quantity  uint64 `json:"quantity"`
    Decimal   uint32 `json:"decimal"`
    From      string `json:"from"`
    To        string `json:"to"`
    Timestamp int64  `json:"timestamp"`
}

func (tr *TransactionVerifyRequest) IsValid() error {
    var valid = false
    var err error
    for _, symbol := range SupportTxSymbols {
        if symbol == strings.ToUpper(tr.Symbol) {
            valid = true
            break
        }
    }

    if !valid {
        err = fmt.Errorf("unsupport symbol: %s", tr.Symbol)
    }

    return err
}
