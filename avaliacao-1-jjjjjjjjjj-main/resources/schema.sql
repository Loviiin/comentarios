CREATE DATABASE IF NOT EXISTS `avaliacao1`;

USE `avaliacao1`;

CREATE TABLE IF NOT EXISTS `products`
(
    `product_id`  INT            NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(45)    NOT NULL,
    `price`       DECIMAL(10, 2) NOT NULL,
    `description` VARCHAR(255)   NOT NULL,
    PRIMARY KEY (`product_id`)
);

CREATE TABLE IF NOT EXISTS `categories`
(
    `category_id` INT         NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(45) NOT NULL,
    PRIMARY KEY (`category_id`)
);

CREATE TABLE IF NOT EXISTS `products_categories`
(
    `product_id`  INT NOT NULL,
    `category_id` INT NOT NULL,
    PRIMARY KEY (`product_id`, `category_id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`),
    FOREIGN KEY (`category_id`) REFERENCES `categories` (`category_id`)
);
