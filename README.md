## token-query

几个常见币种交易验证和查询。


此工具主要用于提供给链上Offchain Worker请求验证使用。

比如一笔比特币转账，需要验证其合法性，需要验证的有：

- sender
- receiver
- quantity
- usdt quantity
- timeout

需求一：目前还差将币额兑换成usdt数量，比对usdt数量是否合法。为防止币价剧烈波动造成的误判，计算出来的usdt数量相差只要不超过50%即可。

需求二：新增其他币种。
