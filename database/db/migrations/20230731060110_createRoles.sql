
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `roles` (
    `id` int unsigned AUTO_INCREMENT NOT NULL comment '権限ID',
    `name` varchar(255) NOT NULL comment '権限名',
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4
comment = '権限を格納';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE `roles`;