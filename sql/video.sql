CREATE TABLE `test`.`video` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `uid` INT NOT NULL,
    `play_url` VARCHAR(255) NOT NULL,
    `cover_url` VARCHAR(255) NOT NULL,
    `favorite_count` INT NOT NULL DEFAULT 0,
    `comment_count` INT NOT NULL DEFAULT 0,
    `title` VARCHAR(50) NOT NULL,
    `create_time` INT NOT NULL,
    `update_time` INT NOT NULL,
    `delete_time` INT NULL,
    PRIMARY KEY (`id`),
    INDEX `index_create_time` (`create_time` DESC),
    INDEX `index_uid` (`uid` ASC));