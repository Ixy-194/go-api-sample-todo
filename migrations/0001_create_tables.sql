CREATE TABLE `todo` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT comment 'ID',
    `task` VARCHAR (128) NOT NULL comment 'タスク',
    `status` VARCHAR(20) NOT NULL comment 'タスクステータス',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '作成日時',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;