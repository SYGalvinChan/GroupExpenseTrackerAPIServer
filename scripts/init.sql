DROP DATABASE IF EXISTS `group_expense_tracker_db`;

CREATE DATABASE `group_expense_tracker_db`;

USE `group_expense_tracker_db`;

CREATE TABLE `users` (
    `user_id` BIGINT PRIMARY KEY,
    `telegram_user_id` BIGINT UNIQUE
);

CREATE TABLE `groups` (
    `group_id` BIGINT PRIMARY KEY,
    `telegram_chat_id` BIGINT UNIQUE
);

CREATE TABLE `users_in_groups` (
    `group_id` BIGINT,
    `user_id` BIGINT,

    PRIMARY KEY (`group_id`, `user_id`),
    FOREIGN KEY (`group_id`) REFERENCES `groups`(`group_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `expenses` (
    `expense_id` BIGINT,
    `group_id` BIGINT,
    `expense_name` VARCHAR(255),
    `amount` DECIMAL(10, 2),

    PRIMARY KEY (`expense_id`),
    FOREIGN KEY (`group_id`) REFERENCES `groups`(`group_id`)
);

CREATE TABLE `paid_by` (
    `expense_id` BIGINT,
    `user_id` BIGINT,
    `amount` DECIMAL(10, 2),

    PRIMARY KEY (`expense_id`, `user_id`),
    FOREIGN KEY (`expense_id`) REFERENCES `expenses`(`expense_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `split_between` (
    `expense_id` BIGINT,
    `user_id` BIGINT,
    `amount` DECIMAL(10, 2),

    PRIMARY KEY (`expense_id`, `user_id`),
    FOREIGN KEY (`expense_id`) REFERENCES `expenses`(`expense_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
)