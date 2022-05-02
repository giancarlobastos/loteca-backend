INSERT INTO lottery_poll
SELECT lm.lottery_id, lm.match_id, u.id, elt(0.5 + rand() * 3, 'H', 'D', 'A'), now() 
FROM lottery_match lm, user u;

SELECT * FROM `stadium` WHERE country = 'Brazil'
ORDER BY rand()
LIMIT 1;

SELECT (@val := @val + 1) AS value 
FROM team, (SELECT @val := 0) AS tt
LIMIT 100;

-- random date
FROM_UNIXTIME(UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)));

select t1.id, t2.id
from team t1, team t2
where 
	t1.id between 33 and 53 and 
	t2.id between 33 and 52 and 
    t1.id != t2.id
order by 1, 2

-- https://stackoverflow.com/questions/6648512/scheduling-algorithm-for-a-round-robin-tournament

SET FOREIGN_KEY_CHECKS=0;    
TRUNCATE TABLE loteca.betting_platform;
TRUNCATE TABLE loteca.competition;
TRUNCATE TABLE loteca.group;
TRUNCATE TABLE loteca.lottery;
TRUNCATE TABLE loteca.lottery_match;
TRUNCATE TABLE loteca.lottery_poll;
TRUNCATE TABLE loteca.match;
TRUNCATE TABLE loteca.match_odds;
TRUNCATE TABLE loteca.round;
TRUNCATE TABLE loteca.season;
TRUNCATE TABLE loteca.stadium;
TRUNCATE TABLE loteca.team;
TRUNCATE TABLE loteca.team_group;
TRUNCATE TABLE loteca.user;
SET FOREIGN_KEY_CHECKS=1;