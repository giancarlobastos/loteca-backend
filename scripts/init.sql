CREATE TABLE `team` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `logo` varchar(255),
  `country` varchar(255)
);

CREATE TABLE `stadium` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `city` varchar(255),
  `state` varchar(255),
  `country` varchar(255)
);

CREATE TABLE `competition` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `division` varchar(255),
  `logo` varchar(255),
  `ended` boolean
);

CREATE TABLE `season` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `code` int
);

CREATE TABLE `group` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `competition_id` int,
  `season_id` int
);

CREATE TABLE `team_group` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `group_id` int,
  `team_id` int
);

CREATE TABLE `round` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `number` int,
  `ended` boolean,
  `competition_id` int,
  `season_id` int
);

CREATE TABLE `match` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `round_id` int,
  `group_id` int,
  `home_id` int,
  `away_id` int,
  `stadium_id` int,
  `start_at` timestamp,
  `home_score` int,
  `away_score` int
);

CREATE TABLE `lottery` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `number` int,
  `estimated_prize` double,
  `main_prize` double,
  `main_prize_winners` int,
  `side_prize` double,
  `side_prize_winners` int,
  `special_prize` double,
  `accumulated` boolean,
  `end_at` timestamp
);

CREATE TABLE `lottery_match` (
  `lottery_id` int,
  `match_id` int,
  `order` int,
  PRIMARY KEY (`lottery_id`, `match_id`)
);

CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `facebook_id` varchar(255),
  `photo` varchar(255)
);

CREATE TABLE `lottery_poll` (
  `lottery_id` int UNIQUE,
  `match_id` int UNIQUE,
  `user_id` int UNIQUE,
  `result` char,
  `voted_at` timestamp,
  PRIMARY KEY (`lottery_id`, `match_id`, `user_id`)
);

CREATE TABLE `betting_platform` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `url` varchar(255)
);

CREATE TABLE `match_odds` (
  `platform_id` int,
  `match_id` int,
  `home_odds` double,
  `away_odds` double,
  `updated_at` timestamp,
  PRIMARY KEY (`platform_id`, `match_id`)
);

ALTER TABLE `group` ADD FOREIGN KEY (`competition_id`) REFERENCES `competition` (`id`);

ALTER TABLE `group` ADD FOREIGN KEY (`season_id`) REFERENCES `season` (`id`);

ALTER TABLE `team_group` ADD FOREIGN KEY (`group_id`) REFERENCES `group` (`id`);

ALTER TABLE `team_group` ADD FOREIGN KEY (`team_id`) REFERENCES `team` (`id`);

ALTER TABLE `round` ADD FOREIGN KEY (`competition_id`) REFERENCES `competition` (`id`);

ALTER TABLE `round` ADD FOREIGN KEY (`season_id`) REFERENCES `season` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`round_id`) REFERENCES `round` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`group_id`) REFERENCES `group` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`home_id`) REFERENCES `team` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`away_id`) REFERENCES `team` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`stadium_id`) REFERENCES `stadium` (`id`);

ALTER TABLE `lottery_match` ADD FOREIGN KEY (`lottery_id`) REFERENCES `lottery` (`id`);

ALTER TABLE `lottery_match` ADD FOREIGN KEY (`match_id`) REFERENCES `match` (`id`);

ALTER TABLE `lottery_poll` ADD FOREIGN KEY (`lottery_id`) REFERENCES `lottery` (`id`);

ALTER TABLE `lottery_poll` ADD FOREIGN KEY (`match_id`) REFERENCES `match` (`id`);

ALTER TABLE `lottery_poll` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `match_odds` ADD FOREIGN KEY (`platform_id`) REFERENCES `betting_platform` (`id`);

ALTER TABLE `match_odds` ADD FOREIGN KEY (`match_id`) REFERENCES `match` (`id`);

