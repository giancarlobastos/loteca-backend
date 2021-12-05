INSERT INTO competition(name, division, code, code_name, rounds)
VALUES ('Campeonato Brasileiro', 'A', 'BRA1', 'campeonato-brasileiro-serie-a', 38);

INSERT INTO season(name, code, competition_id)
VALUES ('2021', 2021, LAST_INSERT_ID());