CREATE TABLE `user` (
    `user_id` VARCHAR(36) PRIMARY KEY,
    `email` VARCHAR(225) UNIQUE NOT NULL,
    `password` VARCHAR(225) NOT NULL
);
CREATE TABLE `customer` (
    `customer_id` VARCHAR(36) PRIMARY KEY,
    `user_id` VARCHAR(36) UNIQUE NOT NULL,
    `name` VARCHAR(225) NOT NULL,
    `address` VARCHAR(225) NOT NULL,
    `date_of_birth` DATE NOT NULL,
    `phone_number` VARCHAR(20) UNIQUE NOT NULL,
    `gender` ENUM('Male', 'Female')
);
CREATE TABLE `seller` (
    `seller_id` VARCHAR(36) PRIMARY KEY,
    `user_id` VARCHAR(36) UNIQUE NOT NULL,
    `name` VARCHAR(225) NOT NULL,
    `balance` FLOAT NOT NULL,
    `address` VARCHAR(225) NOT NULL,
    `date_of_birth` DATE NOT NULL,
    `phone_number` VARCHAR(20) UNIQUE NOT NULL,
    `gender` ENUM('Male', 'Female') NOT NULL,
    `identity_number` VARCHAR(225) NOT NULL,
    `bank_name` VARCHAR(225) NOT NULL,
    `debit_number` INT NOT NULL
);
CREATE TABLE `product` (
    `product_id` VARCHAR(36) PRIMARY KEY,
    `seller_id` VARCHAR(36) NOT NULL,
    `name` VARCHAR(225) NOT NULL,
    `category` VARCHAR(225) NOT NULL,
    `description` TEXT NOT NULL,
    `price` FLOAT NOT NULL,
    `stock` INT NOT NULL,
    `weight` FLOAT NOT NULL
);
CREATE TABLE `order` (
    `order_id` VARCHAR(36) PRIMARY KEY,
    `customer_id` VARCHAR(36) NOT NULL,
    `seller_id` VARCHAR(36) NOT NULL,
    `transaction_date` DATETIME NOT NULL,
    `status` VARCHAR(225) NOT NULL
);
CREATE TABLE `order_product` (
    `order_product_id` VARCHAR(36) PRIMARY KEY,
    `product_id` VARCHAR(36) NOT NULL,
    `order_id` VARCHAR(36) NOT NULL,
    `qty` INT NOT NULL
);
CREATE TABLE `payment` (
    `payment_id` VARCHAR(36) PRIMARY KEY,
    `order_id` VARCHAR(36) UNIQUE NOT NULL,
    `payment_date` DATE NOT NULL,
    `total_payment` FLOAT NOT NULL,
    `payment_method` VARCHAR(20) NOT NULL
);

ALTER TABLE `customer`
ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);
ALTER TABLE `seller`
ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);
ALTER TABLE `product`
ADD FOREIGN KEY (`seller_id`) REFERENCES `seller` (`seller_id`);
ALTER TABLE `order`
ADD FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`);
ALTER TABLE `order`
ADD FOREIGN KEY (`seller_id`) REFERENCES `seller` (`seller_id`);
ALTER TABLE `order_product`
ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`);
ALTER TABLE `order_product`
ADD FOREIGN KEY (`order_id`) REFERENCES `order` (`order_id`);
ALTER TABLE `payment`
ADD FOREIGN KEY (`order_id`) REFERENCES `order` (`order_id`);