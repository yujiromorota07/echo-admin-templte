
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `account_roles` (
    `id` int unsigned NOT NULL comment 'アカウント権限ID',
    `account_id` int unsigned NOT NULL comment 'アカウントID',
    `role_id` int unsigned NOT NULL comment '権限ID'
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`)

) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4
comment = 'アカウントの権限を格納しています'

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

