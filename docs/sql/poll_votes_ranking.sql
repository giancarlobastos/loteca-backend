SELECT lottery_id, user_id, count(*)
FROM (
	 SELECT lp.lottery_id, lp.user_id, h.name home, a.name away, lp.result vote, 
		case when m.home_score > m.away_score then 'H'
		when m.home_score = m.away_score then 'D'
		when m.home_score < m.away_score then 'A'
		else NULL end result
	 FROM lottery_poll lp
	 JOIN lottery_match lm ON lp.match_id = lm.match_id
	 JOIN loteca.match m ON m.id = lp.match_id
	 JOIN team h on m.home_id = h.id
	 JOIN team a on m.away_id = a.id
	 WHERE lp.user_id NOT IN (5,6,33)
	 ORDER BY lm.order
) c
WHERE vote = result
GROUP BY lottery_id, user_id
HAVING count(*) >= 6
ORDER BY 1 DESC, 3 DESC