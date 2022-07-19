SELECT lm.order, match_id, home_id, h.name, away_id, a.name, start_at, home_score, away_score, ended, status, elapsed_time
FROM lottery_match lm
JOIN loteca.match m on m.id = lm.match_id 
JOIN team h on m.home_id = h.id
JOIN team a on m.away_id = a.id
WHERE lm.lottery_id = 1008
ORDER BY lm.order;