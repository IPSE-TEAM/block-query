package main

import (
    "log"
    "token-query/handler"
    "token-query/logger"
    "token-query/route"
    "token-query/tasks"
    "token-query/utils"
)

func main() {
    if err := logger.InitLogWithAdapterConsole("DEBUG", true); err != nil {
        log.Fatal("init logger error: ", err)
    }

    logger.Log.Info("current version: %s", utils.Version)
    logger.Log.Info("start coin and token price query task")

    // 启动定时查询虚拟货币价格的cron
    spc := tasks.NewSymbolPriceCron()
    go spc.Start()

    logger.Log.Info("start coin and token price query service")
    if err := handler.InitPriceService(spc.SymbolPriceWithTether); err != nil {
        logger.Log.Fatal("initialize price service error: %s", err)
    }

    logger.Log.Info("start web server, listen on 0.0.0.0:8080")
    r := route.Router()
    if err := r.Run("0.0.0.0:8080"); err != nil {
        logger.Log.Fatal("can't start token query service, error: %s", err)
    }
}
