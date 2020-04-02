## 查询ETH交易详情

### 通过`https://api.etherscan.io`查询

#### 查询语法

```sh
curl https://api.etherscan.io/api?module=account&action=txlistinternal&txhash=0x40eb908387324f2b575b4879cd9d7188f69c8fc9d87c901b9e2daaea4b442170
```

#### 查询结果

```json
{
	"status": "1",
	"message": "OK",
	"result": [{
		"blockNumber": "1743059",
		"timeStamp": "1466489498",
		"from": "0x2cac6e4b11d6b58f6d3c1c9d5fe8faa89f60e5a2",
		"to": "0x66a1c3eaf0f1ffc28d209c0763ed0ca614f3b002",
		"value": "7106740000000000",
		"contractAddress": "",
		"input": "",
		"type": "call",
		"gas": "2300",
		"gasUsed": "0",
		"isError": "0",
		"errCode": ""
	}]
}
```



### 通过`https://eth.btc.com`查询

#### 查询语法

```sh
curl https://eth.btc.com/txinfo/0x485615bff2000aa18399a0c8314239a395facf7412ee64cb57a75065f6480c84
```

#### 查询结果

```json
{
	"err_no": 0,
	"data": {
		"tx_hash": "0x485615bff2000aa18399a0c8314239a395facf7412ee64cb57a75065f6480c84",
		"status": "success",
		"block_height": 8978645,
		"created_ts": 1574400492,
		"time_in_sec": 3044312,
		"sender_hash": "0x137ad9c4777e1d36e4b605e745e8f37b2b62e9c5",
		"receiver_hash": "0x0c8df6dfb99522d70d3247c4f56358ff23c0d810",
		"tx_type": "call",
		"amount": "0.581870070000000017",
		"amount_usd": "72.704665246500002124",
		"amount_rmb": "508.600990785600014859",
		"fee": "0.000840000000000000",
		"fee_usd": "0.104958000000000000",
		"fee_rmb": "0.734227200000000000",
		"gas_used": 21000,
		"gas_limit": 90000,
		"gas_price": "0.000000040000000000",
		"nonce": 61636,
		"position": 22,
		"input": "0x",
		"error": "",
		"sender_type": 0,
		"receiver_type": 0,
		"event_logs": [],
		"total_confirmation": 191815
	},
	"err_msg": null
}
```

### 查询交易
```sh
curl https://explorer-web.api.btc.com/v1/eth/accounts/0x0c8df6dfb99522d70d3247c4f56358ff23c0d810/txns?page=1&size=25
```
### 响应值
```json
{
	"err_no": 0,
	"data": {
		"list": [{
			"id": 1,
			"tx_hash": "0x462fa7060733c79619a9c7c4a63125ef5807f17cae165285d5f434718c6551dd",
			"block_height": 9264728,
			"created_ts": 1578813672,
			"time_in_sec": 2075323,
			"sender_hash": "0x0c8df6dfb99522d70d3247c4f56358ff23c0d810",
			"receiver_hash": "0x5586039b928ddc915a1f94f649bb04bd15968ab1",
			"amount": "0.100000000000000006",
			"fee": "0.000210000000000000",
			"status": "success",
			"tx_type": "OUT",
			"sender_name": "",
			"receiver_name": "",
			"sender_type": 0,
			"receiver_type": 0
		}, {
			"id": 2,
			"tx_hash": "0x485615bff2000aa18399a0c8314239a395facf7412ee64cb57a75065f6480c84",
			"block_height": 8978645,
			"created_ts": 1574400492,
			"time_in_sec": 6488503,
			"sender_hash": "0x137ad9c4777e1d36e4b605e745e8f37b2b62e9c5",
			"receiver_hash": "0x0c8df6dfb99522d70d3247c4f56358ff23c0d810",
			"amount": "0.581870070000000017",
			"fee": "0.000840000000000000",
			"status": "success",
			"tx_type": "IN",
			"sender_name": "",
			"receiver_name": "",
			"sender_type": 0,
			"receiver_type": 0
		}],
		"page": 1,
		"pagesize": 25,
		"total_count": 2
	},
	"err_msg": null
}
```

