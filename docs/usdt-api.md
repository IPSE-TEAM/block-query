### USDT的三种形态分别是：
- 基于比特币网络的Omni-USDT，充币地址是BTC地址，充提币走BTC网络；
- 基于以太坊ERC20协议的ERC20-USDT，充币地址是ETH地址，充提币走ETH网络；
- 基于波场TRC20协议的TRC20-USDT，充币地址是TRON地址，充提币走TRON网络；

### 三者最明显的区别：

#### 根据地址区分：
- Omni-USDT的地址是以1开头
- ERC20-USDT的地址是以0x开头
- TRC20-USDT的地址是以T开头

#### 根据交易HASH区分
- ERC20-USDT：0xdaaf39ec1b91ea82b9f2ef438eff43009ba7ba960845df93a6b75d875b80ffa6

### 查询ERC20-USDT的交易记录
```sh
curl https://explorer-web.api.btc.com/v1/eth/tokentxns/0xdaaf39ec1b91ea82b9f2ef438eff43009ba7ba960845df93a6b75d875b80ffa6
```

### 查询结果
```json
{
	"err_no": 0,
	"data": {
		"page": 1,
		"pagesize": 10,
		"total_count": 1,
		"list": [{
			"id": 1,
			"tx_hash": "0xdaaf39ec1b91ea82b9f2ef438eff43009ba7ba960845df93a6b75d875b80ffa6",
			"block_height": 9169916,
			"created_ts": 1577435173,
			"time_in_sec": 1031055,
			"sender_hash": "0x53bbf920a9cd498172ed2f069fa046a0dc4a0109",
			"receiver_hash": "0x02055623b3b4cf1dcc16963804c6bb1a834d5fd5",
			"amount": "62000000000",
			"token_hash": "0xdac17f958d2ee523a2206206994597c13d831ec7",
			"token_name": "Tether",
			"token_decimal": 6,
			"unit_name": "USDT",
			"sender_name": "",
			"receiver_name": "",
			"token_url": null,
			"token_icon_url": null,
			"token_found": true
		}]
	},
	"err_msg": null
}
```



