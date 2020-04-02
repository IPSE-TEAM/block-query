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

const (
    USDTBaseEthTokenUnitName string = "Tether"
    USDTBaseEthTokenHash     string = "0xdac17f958d2ee523a2206206994597c13d831ec7"
)

type EthBtcV1TransactionResponse struct {
    ErrNo  int    `json:"err_no"`
    ErrMsg string `json:"err_msg"`
    Data   struct {
        Page       int `json:"page"`
        PageSize   int `json:"pagesize"`
        TotalCount int `json:"total_count"`
        List       []struct {
            TxHash       string `json:"tx_hash"`
            BlockHeight  int64  `json:"block_height"`
            CreatedTs    int64  `json:"created_ts"`
            TimeInSec    int64  `json:"time_in_sec"`
            SenderHash   string `json:"sender_hash"`
            ReceiverHash string `json:"receiver_hash"`
            Amount       string `json:"amount"`
            TokenHash    string `json:"token_hash"`
            TokenName    string `json:"token_name"`
            TokenDecimal int    `json:"token_decimal"`
        } `json:"list"`
    } `json:"data"`
}

func NewUSDTTxQueryResponseWithEthBtc(ebv1tr *EthBtcV1TransactionResponse) (*types.TxQueryResponse, error) {
    // 将USDT的数量转换为以个为单位
    quantity, err := strconv.ParseUint(ebv1tr.Data.List[0].Amount, 10, 64)
    if err != nil {
        return nil, errors.New("USDT tx amount convert into Uint64 failed")
    }

    //quantity := floatAmount / math.Pow10(6)

    return &types.TxQueryResponse{
        Symbol:    strings.ToLower(utils.USDT),
        Timestamp: time.Now().Unix(),
        From:      ebv1tr.Data.List[0].SenderHash,
        To:        ebv1tr.Data.List[0].ReceiverHash,
        Quantity:  quantity,
        Decimal:   uint32(utils.DecimalUSDT),
    }, nil
}

var ErrNotFind = errors.New("not find transaction data")
var ErrTxType = errors.New("transaction is not erc20-usdt")

func fetchUSDTBaseEthTxWithEthBtc(url string) (*EthBtcV1TransactionResponse, error) {
    client := http.Client{Timeout: 20 * time.Second}

    getReq, err := utils.NewCommonRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := client.Do(getReq)
    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, errors.New(string(body))
    }

    logger.Log.Debug("body: %+v", string(body))

    var txDetail EthBtcV1TransactionResponse
    if err := json.Unmarshal(body, &txDetail); err != nil {
        return nil, err
    }

    if txDetail.Data.TotalCount <= 0 {
        return &txDetail, ErrNotFind
    }

    if txDetail.Data.List[0].TokenName != USDTBaseEthTokenUnitName ||
        txDetail.Data.List[0].TokenHash != USDTBaseEthTokenHash {
        return &txDetail, ErrTxType
    }

    return &txDetail, nil
}

func USDTBaseEthTxQueryWithEthBtc(tx string) (*EthBtcV1TransactionResponse, error) {
    url := strings.Join([]string{"https://explorer-web.api.btc.com/v1/eth/tokentxns", tx}, "/")
    return fetchUSDTBaseEthTxWithEthBtc(url)
}

// 校验Tx
func NewUSDTTxVerifyResponseWithEthBtc(ebv1tr *EthBtcV1TransactionResponse, txVerify *types.TransactionVerifyRequest) (int, error) {
    // 将USDT的数量转换为以个为单位
    floatAmount, err := strconv.ParseFloat(ebv1tr.Data.List[0].Amount, 64)
    if err != nil {
        return 50010, errors.New("can't convert usdt tx amount into float64")
    }

    quantity := floatAmount / math.Pow10(6)
    quantityStr := strconv.FormatFloat(quantity, 'f', 6, 64)
    verifyQuantifyStr := strconv.FormatFloat(float64(txVerify.Quantity)/math.Pow10(int(txVerify.Decimal)), 'f', 6, 64)

    timeDelta := int64(math.Abs(float64(txVerify.Timestamp - ebv1tr.Data.List[0].CreatedTs)))

    if quantityStr != verifyQuantifyStr {
        return 20001, errors.New("transaction quantity mismatch")
    }

    if txVerify.From != ebv1tr.Data.List[0].SenderHash {
        return 20002, errors.New("transaction sender hash mismatch")
    }

    if txVerify.To != ebv1tr.Data.List[0].ReceiverHash {
        return 20003, errors.New("transaction receiver hash mismatch")
    }

    if timeDelta > int64(time.Hour*24) {
        return 20004, errors.New("transaction timeout, no longer verify")
    }

    return 20000, nil
}
