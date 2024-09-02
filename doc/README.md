# rpc
proto 檔案位置 [follow.proto](../protos/follow.proto)，各個接口說明為以下：
## Follow 關注
### request
| field       | type   | description | required |
|-------------|--------|-------------|----------| 
| follower_id | string | 為關注者 id     | false    |
| followee_id | string | 為被關注者的 id   | false    |

### response
| field   | type   | discription | 
|---------|--------|-------------|
| code    | int    | 錯誤碼         |
| message | string | 錯誤訊息        |


***

## UnFollow 取消關注
### request
| field       | type   | description | required |
|-------------|--------|-------------|----------| 
| follower_id | string | 為關注者 id     | false    |
| followee_id | string | 為被關注者的 id   | false    |

### response
| field   | type   | discription | 
|---------|--------|-------------|
| code    | int    | 錯誤碼         |
| message | string | 錯誤訊息        |


***


## Followers 關注者列表
### request
| field          | type   | discription                    | required |
|----------------|--------|--------------------------------|----------| 
| user_id        | string | 被關注者的 id                       | false    |
| next_user_id   | string | 上一頁的最後一個 follower 的 id         | true     |
| next_timestamp | string | 上一頁的最後一個 follower 的 created_at | true     |
| limit          | string | 每頁數量                           | false    |


### response
| field   | type   | discription         | 
|---------|--------|---------------------|
| code    | int    | 錯誤碼                 |
| message | string | 錯誤訊息                |
| data    | array  | schema 為 follow 的結構 |

***

## Followees 粉絲列表
### request
| field          | type   | discription                    | required |
|----------------|--------|--------------------------------|----------| 
| user_id        | string | 被關注者的 id                       | false    |
| next_user_id   | string | 上一頁的最後一個 followee 的 id         | true     |
| next_timestamp | string | 上一頁的最後一個 followee 的 created_at | true     |
| limit          | string | 每頁數量                           | false    |

### response
| field   | type   | discription         | 
|---------|--------|---------------------|
| code    | int    | 錯誤碼                 |
| message | string | 錯誤訊息                |
| data    | array  | schema 為 follow 的結構 |

***

## Friends 好友列表
| field          | type   | discription                  | required |
|----------------|--------|------------------------------|----------| 
| user_id        | string | 被關注者的 id                     | false    |
| next_friend_id | string | 上一頁的最後一個 friend 的 id         | true     |
| next_timestamp | string | 上一頁的最後一個 friend 的 created_at | true     |
| limit          | string | 每頁數量                         | false    |

### response
| field   | type   | discription         | 
|---------|--------|---------------------|
| code    | int    | 錯誤碼                 |
| message | string | 錯誤訊息                |
| data    | array  | schema 為 friend 的結構 |


# Schema

## Follow
| field       | type      | discription   | 
|-------------|-----------|---------------|
| follower_id | string    | 關注者的 user_id  |
| followee_id | string    | 被關注者的 user_id |
| created_at  | timestamp | 關注時間          |

## Friend
| field      | type      | discription          | 
|------------|-----------|----------------------|
| user_id    | string    | user_id 擁有的好友        |
| friend_id  | string    | 與 user_id 互相關注的 user |
| created_at | timestamp | 成為好友時間               |