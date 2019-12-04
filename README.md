# UUID
### UUID 的格式：
```
123e4567-e89b-12d3-a456-426655440000
xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
```
* 一共36个字符
* 四位数字 M表示 UUID 版本，数字 N的一至三个最高有效位表示 UUID 变体。在例子中，M 是 1 而且 N 是 a(10xx)，这意味着此 UUID 是 "变体1"、"版本1" UUID；即基于时间的 DCE/RFC 4122 UUID。

**UUID record layout**

|Name|Length (bytes)|Length (hex digits)|Contents|
|---|---|---|---|
|time_low|4|8|整数：低位 32 bits 时间|
|time_mid|2|4|整数：中间位 16 bits 时间|
|time_hi_and_version|2|4|最高有效位中的 4 bits“版本”，后面是高 12 bits 的时间|
|clock_seq_hi_and_res clock_seq_low|2|4|最高有效位为 1-3 bits“变体”，后跟13-15 bits 时钟序列|
|node|6|12|48 bits 节点 ID|

### UUID的版本：
* 版本1：UUID 是根据时间和节点ID（通常是MAC地址）生成
* 版本2：UUID 是根据标识符（通常是组或者用户ID）、时间和节点ID生成
* 版本3和版本5：确定性UUID通过散列（hashing）名字空间（namespace）标识符和名称生成
* 版本4：UUID使用随机性或伪随机性生成

### 性能测试
