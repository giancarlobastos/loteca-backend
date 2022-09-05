CREATE TABLE `team` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `abbreviation` varchar(3),
  `logo` varchar(255),
  `country` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `stadium` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `city` varchar(255),
  `state` varchar(255),
  `country` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `competition` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `division` varchar(255),
  `logo` varchar(255),
  `type` varchar(255),
  `country` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `season` (
  `competition_id` int,
  `year` int,
  `name` varchar(255),
  `ended` boolean,
  PRIMARY KEY (`competition_id`, `year`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `group` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `competition_id` int,
  `year` int
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `team_group` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `group_id` int,
  `team_id` int
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `round` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `number` int,
  `ended` boolean,
  `competition_id` int,
  `year` int
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `match` (
  `id` int PRIMARY KEY,
  `round_id` int,
  `group_id` int,
  `home_id` int,
  `away_id` int,
  `stadium_id` int,
  `start_at` timestamp,
  `home_score` int,
  `away_score` int,
  `ended` boolean,
  `status` varchar(10),
  `elapsed_time` int
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `lottery` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `number` int,
  `estimated_prize` double,
  `main_prize` double,
  `main_prize_winners` int,
  `side_prize` double,
  `side_prize_winners` int,
  `special_prize` double,
  `accumulated` boolean,
  `end_at` timestamp,
  `result_at` timestamp,
  `enabled` boolean NOT NULL DEFAULT FALSE,
  `current` boolean NOT NULL DEFAULT FALSE,
  `updated_at` timestamp NULL DEFAULT NOW()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `lottery_match` (
  `lottery_id` int,
  `match_id` int,
  `order` int,
  `raffle` boolean NOT NULL DEFAULT FALSE,
  `raffle_result` char,
  PRIMARY KEY (`lottery_id`, `match_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `facebook_id` varchar(255),
  `photo` varchar(2000),
  `email` varchar(255),
  `blocked` boolean NOT NULL DEFAULT FALSE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `lottery_poll` (
  `lottery_id` int,
  `match_id` int,
  `user_id` int,
  `result` char,
  `voted_at` timestamp,
  PRIMARY KEY (`lottery_id`, `match_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `betting_platform` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `url` varchar(255),
  `preference` int
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `match_odds` (
  `platform_id` int,
  `match_id` int,
  `home_odds` double,
  `draw_odds` double,
  `away_odds` double,
  `updated_at` timestamp,
  PRIMARY KEY (`platform_id`, `match_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `season` ADD FOREIGN KEY (`competition_id`) REFERENCES `competition` (`id`);

ALTER TABLE `group` ADD FOREIGN KEY (`competition_id`, `year`) REFERENCES `season` (`competition_id`, `year`);

ALTER TABLE `team_group` ADD FOREIGN KEY (`group_id`) REFERENCES `group` (`id`);

ALTER TABLE `team_group` ADD FOREIGN KEY (`team_id`) REFERENCES `team` (`id`);

ALTER TABLE `round` ADD FOREIGN KEY (`competition_id`, `year`) REFERENCES `season` (`competition_id`, `year`);

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

CREATE UNIQUE INDEX `round_index_0` ON `round` (`name`, `competition_id`, `year`);
