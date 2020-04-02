## 查询EOS交易详情

### 通过`https://eos.greymass.com`查询

#### 查询语法

```sh
curl --request POST --url https://eos.greymass.com/v1/history/get_transaction --header 'content-type: application/json' -d '{"id": "2D192BDEA6C52EF0997EB756D300AD6CE02DFD5299A5B58478252A7AC92506E7"}'
```

#### 查询结果

```json
{
	"id": "2d192bdea6c52ef0997eb756d300ad6ce02dfd5299a5b58478252a7ac92506e7",
	"trx": {
		"receipt": {
			"status": "executed",
			"cpu_usage_us": 337,
			"net_usage_words": 20,
			"trx": [1, {
				"signatures": ["SIG_K1_KZvELL6uoGBx6v8aYjU852MCj1UVBf3hmknmv1xfLVbjHwVQDq4p3ohck81E8NxTkEgZDekxv7X3bYRPCVWLwzBmk86xii", "SIG_K1_KbuJmeLRDQoVmYPzgGySAnVZbHKXU1xdTbUEzLy8x2EWjm53VDjdXmh9foReEcV13Lb49kPWGiDD7Yhd1LB4DKR3XUwJaT"],
				"compression": "none",
				"packed_context_free_data": "",
				"packed_trx": "e88d055ef68725252e8100000000019091b97952a47075000000572d3ccdcd020040cd204677320e00000000a8ed323250c81041614cbb4200000000a8ed32322750c81041614cbb425033202308365def20f40e000000000004504f535400000006e88bb1e88bb100"
			}]
		},
		"trx": {
			"expiration": "2019-12-27T04:51:52",
			"ref_block_num": 34806,
			"ref_block_prefix": 2167285029,
			"max_net_usage_words": 0,
			"max_cpu_usage_ms": 0,
			"delay_sec": 0,
			"context_free_actions": [],
			"actions": [{
				"account": "ipsecontract",
				"name": "transfer",
				"authorization": [{
					"actor": "1stbill.tp",
					"permission": "active"
				}, {
					"actor": "cexosse12345",
					"permission": "active"
				}],
				"data": {
					"from": "cexosse12345",
					"to": "xxing2134.tp",
					"quantity": "98.0000 POST",
					"memo": "英英"
				},
				"hex_data": "50c81041614cbb425033202308365def20f40e000000000004504f535400000006e88bb1e88bb1"
			}],
			"transaction_extensions": [],
			"signatures": ["SIG_K1_KZvELL6uoGBx6v8aYjU852MCj1UVBf3hmknmv1xfLVbjHwVQDq4p3ohck81E8NxTkEgZDekxv7X3bYRPCVWLwzBmk86xii", "SIG_K1_KbuJmeLRDQoVmYPzgGySAnVZbHKXU1xdTbUEzLy8x2EWjm53VDjdXmh9foReEcV13Lb49kPWGiDD7Yhd1LB4DKR3XUwJaT"],
			"context_free_data": []
		}
	},
	"block_time": "2019-12-27T04:46:55.000",
	"block_num": 97093957,
	"last_irreversible_block": 97103561,
	"traces": [{
		"action_ordinal": 1,
		"creator_action_ordinal": 0,
		"closest_unnotified_ancestor_action_ordinal": 0,
		"receipt": {
			"receiver": "ipsecontract",
			"act_digest": "bfb0d55dd147738a47a4eb6df48a03ae917ee5249a5d9eb57a7dd77deeecffd9",
			"global_sequence": "28566469766",
			"recv_sequence": 26795957,
			"auth_sequence": [
				["1stbill.tp", 4817603],
				["cexosse12345", 977]
			],
			"code_sequence": 1,
			"abi_sequence": 1
		},
		"receiver": "ipsecontract",
		"act": {
			"account": "ipsecontract",
			"name": "transfer",
			"authorization": [{
				"actor": "1stbill.tp",
				"permission": "active"
			}, {
				"actor": "cexosse12345",
				"permission": "active"
			}],
			"data": {
				"from": "cexosse12345",
				"to": "xxing2134.tp",
				"quantity": "98.0000 POST",
				"memo": "英英"
			},
			"hex_data": "50c81041614cbb425033202308365def20f40e000000000004504f535400000006e88bb1e88bb1"
		},
		"context_free": false,
		"elapsed": 593,
		"console": "",
		"trx_id": "2d192bdea6c52ef0997eb756d300ad6ce02dfd5299a5b58478252a7ac92506e7",
		"block_num": 97093957,
		"block_time": "2019-12-27T04:46:55.000",
		"producer_block_id": "05c98945ac885d8050f15ede7c83df70af61f9ef5aa4fe61c152b8477bc6839e",
		"account_ram_deltas": [{
			"account": "cexosse12345",
			"delta": 128
		}, {
			"account": "ipsecontract",
			"delta": -128
		}],
		"except": null,
		"error_code": null
	}, {
		"action_ordinal": 2,
		"creator_action_ordinal": 1,
		"closest_unnotified_ancestor_action_ordinal": 1,
		"receipt": {
			"receiver": "cexosse12345",
			"act_digest": "bfb0d55dd147738a47a4eb6df48a03ae917ee5249a5d9eb57a7dd77deeecffd9",
			"global_sequence": "28566469767",
			"recv_sequence": 2533,
			"auth_sequence": [
				["1stbill.tp", 4817604],
				["cexosse12345", 978]
			],
			"code_sequence": 1,
			"abi_sequence": 1
		},
		"receiver": "cexosse12345",
		"act": {
			"account": "ipsecontract",
			"name": "transfer",
			"authorization": [{
				"actor": "1stbill.tp",
				"permission": "active"
			}, {
				"actor": "cexosse12345",
				"permission": "active"
			}],
			"data": {
				"from": "cexosse12345",
				"to": "xxing2134.tp",
				"quantity": "98.0000 POST",
				"memo": "英英"
			},
			"hex_data": "50c81041614cbb425033202308365def20f40e000000000004504f535400000006e88bb1e88bb1"
		},
		"context_free": false,
		"elapsed": 11,
		"console": "",
		"trx_id": "2d192bdea6c52ef0997eb756d300ad6ce02dfd5299a5b58478252a7ac92506e7",
		"block_num": 97093957,
		"block_time": "2019-12-27T04:46:55.000",
		"producer_block_id": "05c98945ac885d8050f15ede7c83df70af61f9ef5aa4fe61c152b8477bc6839e",
		"account_ram_deltas": [],
		"except": null,
		"error_code": null
	}, {
		"action_ordinal": 3,
		"creator_action_ordinal": 1,
		"closest_unnotified_ancestor_action_ordinal": 1,
		"receipt": {
			"receiver": "xxing2134.tp",
			"act_digest": "bfb0d55dd147738a47a4eb6df48a03ae917ee5249a5d9eb57a7dd77deeecffd9",
			"global_sequence": "28566469768",
			"recv_sequence": 646,
			"auth_sequence": [
				["1stbill.tp", 4817605],
				["cexosse12345", 979]
			],
			"code_sequence": 1,
			"abi_sequence": 1
		},
		"receiver": "xxing2134.tp",
		"act": {
			"account": "ipsecontract",
			"name": "transfer",
			"authorization": [{
				"actor": "1stbill.tp",
				"permission": "active"
			}, {
				"actor": "cexosse12345",
				"permission": "active"
			}],
			"data": {
				"from": "cexosse12345",
				"to": "xxing2134.tp",
				"quantity": "98.0000 POST",
				"memo": "英英"
			},
			"hex_data": "50c81041614cbb425033202308365def20f40e000000000004504f535400000006e88bb1e88bb1"
		},
		"context_free": false,
		"elapsed": 18,
		"console": "",
		"trx_id": "2d192bdea6c52ef0997eb756d300ad6ce02dfd5299a5b58478252a7ac92506e7",
		"block_num": 97093957,
		"block_time": "2019-12-27T04:46:55.000",
		"producer_block_id": "05c98945ac885d8050f15ede7c83df70af61f9ef5aa4fe61c152b8477bc6839e",
		"account_ram_deltas": [],
		"except": null,
		"error_code": null
	}]
}
```

