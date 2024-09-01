-- Keyspace for follow service
CREATE KEYSPACE IF NOT EXISTS follow
    WITH REPLICATION = {
        'class': 'SimpleStrategy', -- for dev, local
        'replication_factor': 1
        };

-- user A 關注 B（但 B 沒有關注 A）
-- user B 關注 C（但 C 沒有關注 B）
-- 則資料為以下
-- follower_id A, followee_id B
-- follower_id B, followee_id C
CREATE TABLE follow.follows
(
    follower_id UUID,   -- 關注者
    followee_id UUID,   -- 被專注者
    created_at  timestamp,
    PRIMARY KEY (follower_id, followee_id)
);

-- user_id A 擁有了擁有 friend_id 個 friend
-- user A 跟 B 是好友（互相關注）
-- 則資料為以下
-- user_id A, friend_id B
-- user_id B, friend_id A
CREATE TABLE follow.friends
(
    user_id UUID,
    friend_id UUID,
    created_at  timestamp,
    PRIMARY KEY (user_id, friend_id)
);