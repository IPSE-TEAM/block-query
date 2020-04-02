package transaction

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "math"
    "net/http"
    "strconv"
    "strings"
    "time"
    "token-query/logger"
    "token-query/types"
    "token-query/utils"
)

type EOSGreymassTransactionResponse struct {
    Id  string `json:"id"`
    Trx struct {
        Trx struct {
            Expiration       string `json:"expiration"`
            RefBlockNum      int64  `json:"ref_block_num"`
            RefBlockPrefix   int64  `json:"ref_block_prefix"`
            MaxNetUsageWords int64  `json:"max_net_usage_words"`
            MaxCpuUsageMs    int64  `json:"max_cpu_usage_ms"`
            DelaySec         int64  `json:"delay_sec"`
            Actions          []struct {
                Account string `json:"account"`
                Name    string `json:"name"`
                Data    struct {
                    From     string `json:"from"`
                    To       string `json:"to"`
                    Quantity string `json:"quantity"`
                    Memo     string `json:"memo"`
                } `json:"data"`
            } `json:"actions"`
        } `json:"trx"`
    } `json:"trx"`
    BlockTime string `json:"block_time"`
}

type queryData struct {
    Id string `json:"id"`
}

func NewTxQueryResponseWithGreymass(egtr *EOSGreymassTransactionResponse) (*types.TxQueryResponse, error) {
    // egtr的Quantity字段的格式为"0.0001 EOS"，需要去掉EOS
    amount := "0.0000"
    quantity := egtr.Trx.Trx.Actions[0].Data.Quantity
    if len(quantity) >= 4 {
        amount = quantity[0 : len(quantity)-4]
    }

    amountFloat, err := strconv.ParseFloat(amount, 64)
    if err != nil {
        return nil, errors.New("EOS tx amount convert into Float failed")
    }

    amountFloat = amountFloat * math.Pow10(utils.DecimalEOS)

    return &types.TxQueryResponse{
        Symbol:    strings.ToLower(utils.EOS),
        Timestamp: time.Now().Unix(),
        From:      egtr.Trx.Trx.Actions[0].Data.From,
        To:        egtr.Trx.Trx.Actions[0].Data.To,
        Quantity:  uint64(amountFloat),
        Decimal:   uint32(utils.DecimalEOS),
    }, nil
}

func fetchEOSTxWithGreymass(domain, tx string) (*EOSGreymassTransactionResponse, error) {
    url := fmt.Sprintf("%s/v1/history/get_transaction", domain)

    qData := queryData{Id: tx}
    jData, err := json.Marshal(qData)
    if err != nil {
        return nil, err
    }

    client := http.Client{Timeout: 20 * time.Second}
    postReq, err := utils.NewCommonRequest(http.MethodPost, url, bytes.NewReader(jData))
    if err != nil {
        return nil, err
    }

    resp, err := client.Do(postReq)
    if err != nil {
        return nil, errors.New("fetchEOSTxWithGreymass request error: " + err.Error())
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    logger.Log.Debug("body: %+v", string(body))

    if resp.StatusCode != 200 {
        return nil, errors.New(string(body))
    }

    var txDetail EOSGreymassTransactionResponse
    if err := json.Unmarshal(body, &txDetail); err != nil {
        return nil, err
    }

    if _, err := time.Parse("2006-01-02T15:04:05", txDetail.BlockTime); err != nil {
        return &txDetail, fmt.Errorf("can't parse block time, error: %s", err)
    }

    return &txDetail, nil
}

func EOSTxQueryWithGreymass(tx string) (*EOSGreymassTransactionResponse, error) {
    domain := "https://eos.greymass.com"
    return fetchEOSTxWithGreymass(domain, tx)
}

// 校验Tx
func NewEOSTxVerifyResponseWithGreymass(egtr *EOSGreymassTransactionResponse, txVerify *types.TransactionVerifyRequest) (int, error) {
    // egtr的Quantity字段的格式为"0.0001 EOS"，需要去掉EOS
    amount := "0.0000"
    quantity := egtr.Trx.Trx.Actions[0].Data.Quantity
    if len(quantity) >= 4 {
        amount = quantity[0 : len(quantity)-4]
    }
    // 2019-12-25T16:19:48.500 -> unix
    blockTimestamp, _ := time.Parse("2006-01-02T15:04:05", egtr.BlockTime)
    timeDelta := int64(math.Abs(float64(txVerify.Timestamp - blockTimestamp.Unix())))

    verifyQuantifyStr := strconv.FormatFloat(float64(txVerify.Quantity)/math.Pow10(int(txVerify.Decimal)), 'f', 4, 64)

    if verifyQuantifyStr != amount {
        return 20001, errors.New("transaction quantity mismatch")
    }

    if txVerify.From != egtr.Trx.Trx.Actions[0].Data.From {
        return 20002, errors.New("transaction sender hash mismatch")
    }

    if txVerify.To != egtr.Trx.Trx.Actions[0].Data.To {
        return 20003, errors.New("transaction receiver hash mismatch")
    }

    if timeDelta > int64(time.Hour*24) {
        return 20004, errors.New("transaction timeout, no longer verify")
    }

    return 20000, nil
}

// -----------------------------------------------------------------------------------------
