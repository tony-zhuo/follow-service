# Follow Service

## 功能
1. 關注和取消關注功能
2. 採取分頁方式返回粉絲清單、關注列表、好友列表(互相關注)。

## 系統架構
- 流程1（下方圖示）：application 接到 follow 或是 unfollow 時，會將資訊先存在 cache， 存到 cache 後會將 message 丟到 kafka，再由 worker 將 message 接出來進行持久化儲存。
- 流程2：application 接到 followers, followees, friends 的列表請求時，會先從 cache 確認有無資料，有則回傳資料，無則才進 Cassandra 取資料，從 Cassandra 取完資料後，再將資料存進 cache 內。

![image](./doc/arch.jpg)


## 程式碼架構
```text
.
├── cmd/
│   ├── rpc/
│   └── worker/
├── database/
├── delivery/
│   ├── rpc/
│   └── worker/
├── doc/
├── pkg/
├── protos/
└── service/
    ├── model/
    ├── repository/
    └── usecase
```

- entry point(rpc)：rpc 接口的進入點，提供給其他服務使用（proto 檔尚未建置）。
- entry point(worker)：worker 的進入點，作為 kafka 的 consumer。
- database：db 的 init schema 以及 migration 放置處。
- delivery：整合 service 的層級，以及對外的交互介面會放在此處。
- doc：放置文檔，包含 API 文檔。
- pkg：專門放置一些內部套件。
- protos：為放置 proto 檔的地方（實際應該要放在外部）。
- service：專門用來放置核心邏輯的地方。

## 專案說明
詳情可看 [doc](./doc/README.md)
