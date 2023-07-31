
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `accounts` (
    `id` int unsigned AUTO_INCREMENT NOT NULL comment 'アカウントID',
    `name` varchar(255) NOT NULL comment 'ログインユーザー名',
    `email` varchar(255) NOT NULL comment 'ログインユーザーメールアドレス',
    `password` varchar(255) NOT NULL comment 'ログインパスワード',
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp on update current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL comment '削除日' ,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4
comment='クライアントのログイン情報を格納';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `accounts`;
