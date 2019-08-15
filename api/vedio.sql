drop database if exists video_server;

create database video_server;

use video_server;

create TABLE users (
	`id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	`login_name` VARCHAR(64) UNIQUE KEY,
	`pwd` TEXT
) engine=innodb default charset=utf8;

create TABLE videos (
	`id` VARCHAR(64) PRIMARY KEY,
	`author_id` INT UNSIGNED,
	`name` TEXT,
	`display_ctime` varchar(50),
	`create_time` DATETIME
) engine=innodb default charset=utf8;

create TABLE comments (
	`id` VARCHAR(64) PRIMARY KEY,
	`video_id` VARCHAR(64),
	`author_id` INT UNSIGNED,
	`content` TEXT,
	`time` DATETIME
) engine=innodb default charset=utf8;

create TABLE sessions (
	`session_id` VARCHAR(200) PRIMARY KEY,
	`TTL` TINYTEXT,
	`login_name` VARCHAR(64)
) engine=innodb default charset=utf8;