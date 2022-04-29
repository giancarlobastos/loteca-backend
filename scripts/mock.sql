INSERT INTO lottery_poll
SELECT lm.lottery_id, lm.match_id, u.id, elt(0.5 + rand() * 3, 'H', 'D', 'A'), now() 
FROM lottery_match lm, user u;

