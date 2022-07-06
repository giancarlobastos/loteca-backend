SELECT h.name home, a.name away, lp.result vote, 
    case when m.home_score > m.away_score then 'H'
    when m.home_score = m.away_score then 'D'
    else 'A' end result
FROM lottery_poll lp
JOIN lottery_match lm ON lp.match_id = lm.match_id
JOIN loteca.match m ON m.id = lp.match_id
JOIN team h on m.home_id = h.id
JOIN team a on m.away_id = a.id
WHERE lp.lottery_id = 1007 AND lp.user_id = 35
ORDER BY lm.order
