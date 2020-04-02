package tasks

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/robfig/cron"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
    "time"
    "token-query/logger"
    "token-query/types"
    "token-query/utils"
)

// 定时查询Symbol的价格，单位为USDT
//type SymbolPriceWithTether struct {
//    EthPrice  string `json:"eth_price"`
//    BtcPrice  string `json:"btc_price"`
//    EosPrice  string `json:"eos_price"`
//    Timestamp int64  `json:"timestamp"`
//}

type SymbolPriceCron struct {
    cronTask              *cron.Cron
    SymbolPriceWithTether *types.PriceQueryResponse
    lastQueryTime         time.Time
}

func NewSymbolPriceCron() *SymbolPriceCron {
    return &SymbolPriceCron{
        cronTask: cron.New(),
        SymbolPriceWithTether: &types.PriceQueryResponse{
            ApiGateWay: "https://min-api.cryptocompare.com",
            PriceList:  make([]map[string]string, 3, 3),
        },
        lastQueryTime: time.Time{},
    }
}

func (spc *SymbolPriceCron) Start() {
    if (spc.lastQueryTime == time.Time{}) {
        spc.priceQuery()
    }

    spec := "0 0 0/2 * * *"
    if err := spc.cronTask.AddFunc(spec, spc.priceQuery); err != nil {
        logger.Log.Fatal("add spc.priceQuery into cron error: %s", err)
    }

    spc.cronTask.Start()
    select {}
}

type priceQueryResultWithCryptoCompare struct {
    PriceWithTether float64 `json:"USD"`
}

func (spc *SymbolPriceCron) priceQuery() {
    for _, symbol := range types.SupportPriceSymbols {
        if err := spc.symbolPriceQuery(symbol); err != nil {
            logger.Log.Error("query symbol '%s' price error: %s", symbol, err)
        }

        time.Sleep(1 * time.Second)
    }

    spc.lastQueryTime = time.Now()
    logger.Log.Debug("symbol price: %+v", spc.SymbolPriceWithTether)
}

func (spc *SymbolPriceCron) symbolPriceQuery(symbol string) error {
    url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=USD", strings.ToUpper(symbol))
    client := http.Client{Timeout: 20 * time.Second}
    getReq, err := utils.NewCommonRequest(http.MethodGet, url, nil)
    if err != nil {
        return err
    }

    resp, err := client.Do(getReq)
    if err != nil {
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return errors.New(string(body))
    }

    logger.Log.Debug("%s: %+v", symbol, string(body))

    var priceDetail priceQueryResultWithCryptoCompare
    if err := json.Unmarshal(body, &priceDetail); err != nil {
        return err
    }

    switch strings.ToUpper(symbol) {
    case utils.ETH:
        priceEth := map[string]string{strings.ToLower(utils.ETH): strconv.FormatFloat(priceDetail.PriceWithTether, 'f', 4, 64)}
        spc.SymbolPriceWithTether.PriceList[0] = priceEth
    case utils.EOS:
        priceEos := map[string]string{strings.ToLower(utils.EOS): strconv.FormatFloat(priceDetail.PriceWithTether, 'f', 4, 64)}
        spc.SymbolPriceWithTether.PriceList[1] = priceEos
    case utils.BTC:
        priceBtc := map[string]string{strings.ToLower(utils.BTC): strconv.FormatFloat(priceDetail.PriceWithTether, 'f', 4, 64)}
        spc.SymbolPriceWithTether.PriceList[2] = priceBtc
    }

    spc.SymbolPriceWithTether.Timestamp = time.Now().Unix()

    return nil
}
