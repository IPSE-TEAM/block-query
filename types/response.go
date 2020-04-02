package types

type TxQueryResponse struct {
    Symbol    string `json:"symbol"`
    Timestamp int64  `json:"timestamp"`
    From      string `json:"from"`
    To        string `json:"to"`
    Quantity  uint64 `json:"quantity"`
    Decimal   uint32 `json:"decimal"`
}

type PriceQueryResponse struct {
    ApiGateWay string              `json:"api_gate_way"`
    PriceList  []map[string]string `json:"price"`
    Timestamp  int64               `json:"timestamp"`
}

type TransactionVerifyResponse struct {
    TransactionVerifyRequest
    VerifyCode int    `json:"verify_status"`
    VerifyMsg  string `json:"verify_msg"`
}
