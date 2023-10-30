# msgpack-json-converter

思維 JSON -> struct -> MessagePack

## 系統需求

tested @ go:1.21

## 實作內容

主要開發模組 `pkg/msgpack`  

※ 僅在測試驗證時引入他人實作的 msgpack 模組  

### 測試

* 測試 `pkg/msgpack`

```shell
go test ./pkg/msgpack/... -v
```

* 測試 JSON -> MessagePack

```shell
go test ./main_test.go
```

## 尚未實作的部分

因時間關係，尚未實作部分延伸內容

* 僅開發 `Marshal` ，Unmarshal 未實作
* msgpack 的 struct 型別的 struct tag 功能尚未實作，暫以 field name 代替
* 部分 Golang 型別因為不會出現在 JSON unmarshaled struct 內，故未實作
* 因時間關係，沒有測試較為複雜的物件狀態，可能有部分細節疏漏，請見諒
