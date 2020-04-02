package transaction

import (
    "log"
    "testing"
    "token-query/logger"
)

func TestEOSTxQueryWithGreymass(t *testing.T) {
    if err := logger.InitLogWithAdapterConsole("DEBUG", true); err != nil {
        log.Fatal("init logger error: ", err)
    }

    tx := "3cff044f209675a45df7b634fbada39633421071faaf1a75c0497dd007a2474b"

    txDetail, err := EOSTxQueryWithGreymass(tx)
    if err != nil {
        t.Error("EOSTxQueryWithGreymass error: ", err)
    }

    if txDetail != nil && txDetail.Trx.Trx.Actions[0].Data.From != "huanglingbin" {
        t.Error("EOSTxQueryWithGreymass query failed!")
    }
}
