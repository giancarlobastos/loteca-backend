SELECT lp.lottery_id, lp.match_id, h.name, a.name,
	count(case when lp.result = 'H' then 1 else null end) as home_votes,
	count(case when lp.result = 'D' then 1 else null end) as draw_votes,
	count(case when lp.result = 'A' then 1 else null end) as away_votes,
	count(DISTINCT lp.user_id) total_votes
FROM lottery_poll lp
JOIN lottery_match lm ON lp.lottery_id = lm.lottery_id AND lp.match_id = lm.match_id
JOIN loteca.match m ON m.id = lp.match_id
JOIN team h on m.home_id = h.id
JOIN team a on m.away_id = a.id
WHERE lp.lottery_id = 1007 AND lp.user_id NOT IN (1, 5, 6, 33) 
GROUP BY lp.lottery_id, lp.match_id, h.name, a.name
ORDER BY lm.order