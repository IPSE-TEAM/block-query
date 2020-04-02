package transaction

import (
    "log"
    "testing"
    "token-query/logger"
)

func TestBTCTxQueryWithBlockChain(t *testing.T) {
    if err := logger.InitLogWithAdapterConsole("DEBUG", true); err != nil {
        log.Fatal("init logger error: ", err)
    }

    tx := "ceee32a2b528591aa92376812dfea6f1c714243387dce19f224cd38405cbc37e"
    txDetail, err := BTCTxQueryWithBlockChain(tx)
    if err != nil {
        t.Error("BTCTxQueryWithBlockChain error: ", err)
    }

    if txDetail.TxId != tx {
        t.Error("BTCTxQueryWithBlockChain query failed")
    }
}
