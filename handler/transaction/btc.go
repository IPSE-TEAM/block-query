package transaction

import (
    "encoding/json"
    "errors"
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

type BTCBlockChainTransactionResponse struct {
    TxId  string `json:"txid"`
    Time  int64  `json:"time"`
    Block struct {
        MemPool int64 `json:"mempool"`
    } `json:"block"`
    Deleted bool  `json:"deleted"`
    Fee     int64 `json:"fee"`
    Inputs  []struct {
        Address  string `json:"address"`
        CoinBase bool   `json:"coinbase"`
        PkScript string `json:"pkscript"`
        Value    int64  `json:"value"`
    } `json:"inputs"`
    Outputs []struct {
        Address  string `json:"address"`
        PkScript string `json:"pkscript"`
        Spent    bool   `json:"spent"`
        Value    int64  `json:"value"`
    } `json:"outputs"`
}

func (bbctr *BTCBlockChainTransactionResponse) getSenderReceiverValue() (sender, receiver string, value int64) {
    for _, output := range bbctr.Outputs {
        for _, input := range bbctr.Inputs {
            if input.Address != output.Address {
                if receiver != "" {
                    receiver = strings.Join([]string{receiver, output.Address}, ",")
                } else {
                    receiver = output.Address
                }
                value += output.Value
            } else {
                if sender != "" {
                    sender = strings.Join([]string{sender, input.Address}, ",")
                } else {
                    sender = input.Address
                }
            }
        }
    }
    return
}

func NewTxQueryResponseWithBlockChain(bbctr *BTCBlockChainTransactionResponse) (*types.TxQueryResponse, error) {
    sender, receiver, value := bbctr.getSenderReceiverValue()

    // 将比特币的数量转换为以个为单位
    //quantity := float64(value) / 1e8

    return &types.TxQueryResponse{
        Symbol:    strings.ToLower(utils.BTC),
        Timestamp: time.Now().Unix(),
        From:      sender,
        To:        receiver,
        Quantity:  uint64(value),
        Decimal:   uint32(utils.DecimalBTC),
    }, nil
}

func fetchBTCTxWithBlockChain(url string) (*BTCBlockChainTransactionResponse, error) {
    client := http.Client{Timeout: 20 * time.Second}
    postReq, err := utils.NewCommonRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := client.Do(postReq)
    if err != nil {
        return nil, errors.New("fetchEOSTxWithBlockChain request error: " + err.Error())
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    logger.Log.Debug("body: %+v", string(body))

    if resp.StatusCode != 200 {
        return nil, errors.New(string(body))
    }

    var txDetail BTCBlockChainTransactionResponse
    if err := json.Unmarshal(body, &txDetail); err != nil {
        return nil, err
    }

    return &txDetail, nil
}

func BTCTxQueryWithBlockChain(tx string) (*BTCBlockChainTransactionResponse, error) {
    url := strings.Join([]string{"https://api.blockchain.info/haskoin-store/btc/transaction", tx}, "/")
    return fetchBTCTxWithBlockChain(url)
}

// 校验Tx
func NewBTCTxVerifyResponseWithBlockChain(bbctr *BTCBlockChainTransactionResponse, txVerify *types.TransactionVerifyRequest) (int, error) {
    // 将比特币的数量转换为以个为单位
    sender, receiver, value := bbctr.getSenderReceiverValue()
    quantity := float64(value) / 1e8
    quantityStr := strconv.FormatFloat(quantity, 'f', 8, 64)
    verifyQuantifyStr := strconv.FormatFloat(float64(txVerify.Quantity)/math.Pow10(int(txVerify.Decimal)), 'f', 8, 64)

    timeDelta := int64(math.Abs(float64(txVerify.Timestamp - bbctr.Time)))

    if verifyQuantifyStr != quantityStr {
        return 20001, errors.New("transaction quantity mismatch")
    }

    if txVerify.From != sender {
        return 20002, errors.New("transaction sender hash mismatch")
    }

    if txVerify.To != receiver {
        return 20003, errors.New("transaction receiver hash mismatch")
    }

    if timeDelta > int64(time.Hour*24) {
        return 20004, errors.New("transaction timeout, no longer verify")
    }

    return 20000, nil
}
