package com.evilvaca.loteca.scrapers;

import com.evilvaca.loteca.domain.Match;
import com.evilvaca.loteca.domain.Round;
import lombok.RequiredArgsConstructor;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;
import java.util.stream.Collectors;

@Component
@RequiredArgsConstructor(onConstructor_ = @Autowired)
public class TransferMarktLeagueScraper {

    public static final String HOST_URL = "https://www.transfermarkt.com.br";

    private static final String TOURNAMENT_URL = HOST_URL + "/{tournamentCodeDescription}/spieltag/wettbewerb/{tournamentCode}/saison_id/{seasonCode}/spieltag/{roundCode}";

    private static final DateTimeFormatter DATE_TIME_FORMATTER = DateTimeFormatter.ofPattern("ddMMyyHHmm");

    private final RestTemplate restTemplate;

    public List<Match> getMatchList(Round round) {
        return getDocument(TOURNAMENT_URL,
                round.getSeason().getTournament().getCodeDescription(),
                round.getSeason().getTournament().getCode(),
                round.getSeason().getCode(),
                round.getCode())
                .select("table tbody tr.table-grosse-schrift").stream()
                .map(match -> getMatchDetails(match.parents().get(2).select("div.footer a").attr("href")))
                .collect(Collectors.toList());
    }

    private Match getMatchDetails(String url) {
        String matchId = url.substring(url.lastIndexOf('/') + 1);
        Document doc = getDocument(HOST_URL + url);

        String homeTeam = doc.select(" div.sb-team a.sb-vereinslink").get(0).attr("href")
                .replaceFirst("[^0-9]*", "").split("/")[0];

        String awayTeam = doc.select(" div.sb-team a.sb-vereinslink").get(1).attr("href")
                .replaceFirst("[^0-9]*", "").split("/")[0];

        boolean hasScore = doc.select(".sb-endstand").text().matches(".*[0-9]:[0-9].*");

        String homeScore = hasScore ? doc.select(".sb-endstand").text().split("[: ]")[0] : null;
        String awayScore = hasScore ? doc.select(".sb-endstand").text().split("[: ]")[1] : null;
        LocalDateTime matchDateTime = LocalDateTime.parse(doc.select(".sb-datum").text()
                .replaceFirst(".+?\\|", "")
                .replaceAll("[^0-9]", ""), DATE_TIME_FORMATTER);
        String stadium = doc.select(".sb-zusatzinfos span a").text();

        return Match.builder()
                .id(Long.parseLong(matchId))
                .homeId(Long.parseLong(homeTeam))
                .awayId(Long.parseLong(awayTeam))
                .homeScore(hasScore ? Integer.parseInt(homeScore) : null)
                .awayScore(hasScore ? Integer.parseInt(awayScore) : null)
                .startAt(matchDateTime)
                .stadium(stadium)
                .build();
    }

    private Document getDocument(String url, Object... uriVariables) {
        final HttpHeaders headers = new HttpHeaders();
        headers.set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36");
        ResponseEntity<String> response = restTemplate.exchange(url, HttpMethod.GET, new HttpEntity<String>(headers), String.class, uriVariables);
        return Jsoup.parse(response.toString());
    }
}
