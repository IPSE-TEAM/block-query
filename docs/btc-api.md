## 查询BTC交易详情

### 查询网站
```url
https://www.blockchain.com/btc/tx/ceee32a2b528591aa92376812dfea6f1c714243387dce19f224cd38405cbc37e
```

### 通过`https://www.blockchain.com`查询
```sh
curl https://api.blockchain.info/haskoin-store/btc/transaction/ceee32a2b528591aa92376812dfea6f1c714243387dce19f224cd38405cbc37e
```

#### 查询结果
```json
{
	"txid": "ceee32a2b528591aa92376812dfea6f1c714243387dce19f224cd38405cbc37e",
	"size": 225,
	"version": 1,
	"locktime": 0,
	"fee": 2000,
	"inputs": [{
		"pkscript": "76a91407907044b9b1e14146be55c4deb2259f0e0acb2988ac",
		"value": 2127004528,
		"address": "1gzm7L4GNbNDUkfWZxQJdZt6b7tjoWzRb",
		"witness": null,
		"sequence": 4294967295,
		"output": 1,
		"sigscript": "47304402203733e9cfb8184990a712e324f88e93f9a477570a04c6a29986408697a501462302207d28979a5679604b6fefd86265fdebf1024538ca5d27357932a02fe00f2f6710012103666ad8b082896a655ea44787e75ba36e969999dbd2a26cb874d8aa6cf32cbacd",
		"coinbase": false,
		"txid": "9d1b64ad701f6fb57122a42fd514f46aac44d53494c3f0e444118c774d20d57e"
	}],
	"outputs": [{
		"spent": true,
		"pkscript": "76a9142f3d70599ca8ee9ce0054c092ab66999cfa33de088ac",
		"value": 999900000,
		"address": "15JnMs5PLUxWXMqBhEkpYaE9ib6mxPc6qK",
		"spender": {
			"input": 4,
			"txid": "1bfa8b6fb0c93e6f45f4ffc1e6c594b1e8612825b443846b26709ae8b2d67cb1"
		}
	}, {
		"spent": true,
		"pkscript": "76a91407907044b9b1e14146be55c4deb2259f0e0acb2988ac",
		"value": 1127102528,
		"address": "1gzm7L4GNbNDUkfWZxQJdZt6b7tjoWzRb",
		"spender": {
			"input": 0,
			"txid": "aa922ff5e81b024b20af24e0cba7c71e06887c625282bae57ebd9b3db14c746e"
		}
	}],
	"block": {
		"height": 611064,
		"position": 1030
	},
	"deleted": false,
	"time": 1578038263,
	"rbf": false,
	"weight": 900
}
```

