### teams
https://loteca.click/manager/Bolivia/teams
https://loteca.click/manager/Argentina/teams
https://loteca.click/manager/England/teams
https://loteca.click/manager/Spain/teams
https://loteca.click/manager/Paraguay/teams
https://loteca.click/manager/Colombia/teams
https://loteca.click/manager/Uruguay/teams
https://loteca.click/manager/Ecuador/teams
https://loteca.click/manager/Peru/teams
https://loteca.click/manager/Chile/teams
https://loteca.click/manager/Venezuela/teams<br>
`SELECT * FROM loteca.team WHERE country = 'Ecuador'`

### competition, season
https://loteca.click/manager/World/competitions
https://loteca.click/manager/Brazil/competitions<br>
```
 SELECT * 
 FROM loteca.season s
 JOIN loteca.competition c ON s.competition_id = c.id
 WHERE year = 2021 AND country = 'England'
 ```

### matches
#### libertadores
https://loteca.click/manager/World/competitions/13/2022/matches

#### sul-americana
https://loteca.click/manager/World/competitions/11/2022/matches

#### champions league
https://loteca.click/manager/World/competitions/2/2021/matches

#### europe league
https://loteca.click/manager/World/competitions/3/2021/matches

#### europe conference league
https://loteca.click/manager/World/competitions/848/2021/matches

#### uefa nations league
https://loteca.click/manager/World/competitions/5/2022/matches

#### serie a
https://loteca.click/manager/Brazil/competitions/71/2022/matches

#### serie b
https://loteca.click/manager/Brazil/competitions/72/2022/matches

#### serie c
https://loteca.click/manager/Brazil/competitions/75/2022/matches

#### copa brasil
https://loteca.click/manager/Brazil/competitions/73/2022/matches

#### la liga
https://loteca.click/manager/Spain/competitions/140/2021/matches

### get matches
https://loteca.click/manager/World/competitions/2/2021/matches

#### premier league
https://loteca.click/manager/England/competitions/39/2021/matches

#### ligue 1 - france
https://loteca.click/manager/France/competitions/61/2021/matches

#### copa alemanha
https://loteca.click/manager/Germany/competitions/81/2021/matches

#### serie A - italy
https://loteca.click/manager/Italy/competitions/135/2021/matches







```
select count(*)
FROM (SELECT h.name home, a.name away, lp.result vote, 
	case when m.home_score > m.away_score then 'H'
	when m.home_score = m.away_score then 'D'
	else 'A' end result
 FROM lottery_poll lp
 JOIN lottery_match lm ON lp.match_id = lm.match_id
 JOIN loteca.match m ON m.id = lp.match_id
 JOIN team h on m.home_id = h.id
 JOIN team a on m.away_id = a.id
 WHERE lp.lottery_id = 1006 AND lp.user_id = 13
 ORDER BY lm.order) c
 WHERE vote = result
```