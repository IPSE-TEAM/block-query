package transaction

import (
    "log"
    "testing"
    "token-query/logger"
)

func TestETHTxQueryWithEthBtc(t *testing.T) {
    if err := logger.InitLogWithAdapterConsole("DEBUG", true); err != nil {
        log.Fatal("init logger error: ", err)
    }

    tx := "0xc5bba9b4659a0fe4f4e12bb12cadfa6689a042073025eca0cb9340b092ceb851"
    txDetail, err := ETHTxQueryWithEthBtc(tx)
    if err != nil {
        t.Error("ETHTxQueryWithEthBtc error: ", err)
    }

    if txDetail.Data.TxHash != tx {
        t.Error("ETHTxQueryWithEthBtc query failed")
    }
}
