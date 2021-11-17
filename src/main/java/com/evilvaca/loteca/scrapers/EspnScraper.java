package com.evilvaca.loteca.scrapers;

import com.evilvaca.loteca.domain.Match;
import com.evilvaca.loteca.scrapers.domain.EspnData;
import lombok.RequiredArgsConstructor;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

import java.time.ZonedDateTime;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

@Component
@RequiredArgsConstructor(onConstructor_ = @Autowired)
public class EspnScraper {

    private static final String TOURNAMENT_URL = "https://secure.espn.com/core/futebol/fixtures/_/data/{date}/liga/{tournament}?table=true&device=desktop&country=br&lang=pt&region=br&site=espn&edition-host=espn.com.br";

    private static final Pattern ID_PATTERN = Pattern.compile("[^0-9]*(\\d+)[^0-9]*");

    private static final Pattern MATCH_RESULT_PATTERN = Pattern.compile("(\\d+)[^0-9]*(\\d+)");

    private static final DateTimeFormatter DATE_TIME_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mmz");

    private final RestTemplate restTemplate;

    public List<Match> getMatchList(String tournament, String date) {
        EspnData espnData = restTemplate.getForObject(TOURNAMENT_URL, EspnData.class, date, tournament);
        Document doc = Jsoup.parse(espnData.getTable().replace('\\', ' '));
        Elements tableCaption = doc.select("h2.table-caption");
        List<Match> matches = new ArrayList<>();
        tableCaption.forEach(matchDayElement -> {
            String matchDate = tableCaption.first().text();
            matchDayElement.nextElementSibling().select("tbody tr").forEach(match -> {
                String homeTeamLink = ((Element) match.childNode(0)).select("a.team-name").attr("href");


                Matcher matcher = ID_PATTERN.matcher(homeTeamLink);
                Long homeTeamId = matcher.matches() ? Long.parseLong(matcher.group(1)) : null;
                String awayTeamLink = ((Element) match.childNode(1)).select("a.team-name").attr("href");

                matcher = ID_PATTERN.matcher(awayTeamLink);
                Long awayTeamId = matcher.matches() ? Long.parseLong(matcher.group(1)) : null;
                Elements resultElement = ((Element) match.childNode(0)).select(".record a");
                String matchLink = resultElement.attr("href");

                matcher = ID_PATTERN.matcher(matchLink);
                Long matchId = matcher.matches() ? Long.parseLong(matcher.group(1)) : null;

                String result = resultElement.first().text();
                Matcher resultMatcher = MATCH_RESULT_PATTERN.matcher(result);
                boolean hasResult = resultMatcher.matches();
                Integer homeScore = hasResult ? Integer.parseInt(resultMatcher.group(1)) : null;
                Integer awayScore = hasResult ? Integer.parseInt(resultMatcher.group(2)) : null;
                String stadium = match.getElementsByClass("schedule-location").text();
                String matchTime = match.getElementsByAttribute("data-date").size() == 0 ? null : match.getElementsByAttribute("data-date").attr("data-date");
                matches.add(Match.builder()
                        .id(matchId)
                        .homeId(homeTeamId)
                        .awayId(awayTeamId)
                        .homeScore(homeScore)
                        .awayScore(awayScore)
                        .startAt(ZonedDateTime.parse(matchTime, DATE_TIME_FORMATTER).toLocalDateTime())
                        .build());
            });
        });
        return matches;
    }
}
