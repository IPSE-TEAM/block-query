package transaction

import (
    "log"
    "testing"
    "token-query/logger"
)

func TestUSDTBaseEthTxQueryWithEthBtc(t *testing.T) {
    if err := logger.InitLogWithAdapterConsole("DEBUG", true); err != nil {
        log.Fatal("init logger error: ", err)
    }

    tx := "0xdaaf39ec1b91ea82b9f2ef438eff43009ba7ba960845df93a6b75d875b80ffa6"
    txDetail, err := USDTBaseEthTxQueryWithEthBtc(tx)
    if err != nil {
        t.Error("USDTBaseEthTxQueryWithEthBtc error: ", err)
    }

    if txDetail.Data.TotalCount >= 1 {
        if txDetail.Data.List[0].TxHash != tx {
            t.Error("USDTBaseEthTxQueryWithEthBtc query failed")
        }
    } else {
        t.Error("USDTBaseEthTxQueryWithEthBtc query failed")
    }
}
