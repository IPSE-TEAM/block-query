package transaction

import (
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

type EtherScanTransactionResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Result  []struct {
        BlockNumber     string `json:"blockNumber"`
        TimeStamp       string `json:"timeStamp"`
        From            string `json:"from"`
        To              string `json:"to"`
        Value           string `json:"value"`
        ContractAddress string `json:"contractAddress"`
        Input           string `json:"input"`
        Type            string `json:"type"`
        Gas             string `json:"gas"`
        GasUsed         string `json:"gasUsed"`
        IsError         string `json:"isError"`
        ErrCode         string `json:"errCode"`
    } `json:"result"`
}

func NewTxQueryResponseWithEtherScan(estr *EtherScanTransactionResponse) *types.TxQueryResponse {
    return &types.TxQueryResponse{
        Symbol:    strings.ToLower(utils.ETH),
        Timestamp: time.Now().Unix(),
        From:      estr.Result[0].From,
        To:        estr.Result[0].To,
        //Quantity:  estr.Result[0].Value,
        Quantity: uint64(0),
        Decimal:  uint32(utils.DecimalETH),
    }
}

func fetchETHTxWithEtherScan(url string) (*EtherScanTransactionResponse, error) {
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

    var txDetail EtherScanTransactionResponse
    if err := json.Unmarshal(body, &txDetail); err != nil {
        return nil, err
    }

    if txDetail.Status != "1" && txDetail.Message != "OK" {
        return nil, errors.New(txDetail.Message)
    }

    return &txDetail, nil
}

func ETHTxQueryWithEtherScan(tx string) (*EtherScanTransactionResponse, error) {
    domain := "https://etherscan.io"
    url := fmt.Sprintf("%s/api?module=account&action=txlistinternal&txhash=%s", domain, tx)
    return fetchETHTxWithEtherScan(url)
}

// ---------------------------------------------------------------------------------------------------

type EthBtcTransactionResponse struct {
    ErrNo  int    `json:"err_no"`
    ErrMsg string `json:"err_msg"`
    Data   struct {
        TxHash       string `json:"tx_hash"`
        BlockHeight  int64  `json:"block_height"`
        CreatedTs    int64  `json:"created_ts"`
        TimeInSec    int64  `json:"time_in_sec"`
        SenderHash   string `json:"sender_hash"`
        ReceiverHash string `json:"receiver_hash"`
        Amount       string `json:"amount"`
    } `json:"data"`
}

func NewTxQueryResponseWithEthBtc(ebtr *EthBtcTransactionResponse) (*types.TxQueryResponse, error) {
    // ebtr.Data.Amount的单位就是个，所以只需要转换精度
    amount, err := strconv.ParseFloat(ebtr.Data.Amount, 64)
    if err != nil {
        return nil, errors.New("ETH tx amount convert into Float failed")
    }

    quantity := amount * math.Pow10(utils.DecimalETH)
    return &types.TxQueryResponse{
        Symbol:    strings.ToLower(utils.ETH),
        Timestamp: time.Now().Unix(),
        From:      ebtr.Data.SenderHash,
        To:        ebtr.Data.ReceiverHash,
        Quantity:  uint64(quantity),
        Decimal:   uint32(utils.DecimalETH),
    }, nil
}

func fetchETHTxWithEthBtc(url string) (*EthBtcTransactionResponse, error) {
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

    var txDetail EthBtcTransactionResponse
    if err := json.Unmarshal(body, &txDetail); err != nil {
        return nil, err
    }

    // ebtr.Data.Amount的单位就是个，所以只需要转换精度
    if _, err := strconv.ParseFloat(txDetail.Data.Amount, 64); err != nil {
        return &txDetail, errors.New("can't convert amount into f64")
    }

    return &txDetail, nil
}

func ETHTxQueryWithEthBtc(tx string) (*EthBtcTransactionResponse, error) {
    url := strings.Join([]string{"https://explorer-web.api.btc.com/v1/eth/txns", tx}, "/")
    return fetchETHTxWithEthBtc(url)
}

// 校验Tx
func NewETHTxVerifyResponseWithEthBtc(ebtr *EthBtcTransactionResponse, txVerify *types.TransactionVerifyRequest) (int, error) {
    // ebtr.Data.Amount的单位就是个，所以只需要转换精度
    amount, _ := strconv.ParseFloat(ebtr.Data.Amount, 64)
    quantityStr := strconv.FormatFloat(amount, 'f', 6, 64)
    verifyQuantifyStr := strconv.FormatFloat(float64(txVerify.Quantity)/math.Pow10(int(txVerify.Decimal)), 'f', 6, 64)

    timeDelta := int64(math.Abs(float64(txVerify.Timestamp - ebtr.Data.CreatedTs)))

    if verifyQuantifyStr != quantityStr {
        return 20001, errors.New("transaction quantity mismatch")
    }

    if txVerify.From != ebtr.Data.SenderHash {
        return 20002, errors.New("transaction sender hash mismatch")
    }

    if txVerify.To != ebtr.Data.ReceiverHash {
        return 20003, errors.New("transaction receiver hash mismatch")
    }

    if timeDelta > int64(time.Hour*24) {
        return 20004, errors.New("transaction timeout, no longer verify")
    }

    return 20000, nil
}
