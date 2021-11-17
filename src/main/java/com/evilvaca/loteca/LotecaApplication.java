package com.evilvaca.loteca;

import com.evilvaca.loteca.domain.Round;
import com.evilvaca.loteca.domain.Season;
import com.evilvaca.loteca.domain.Tournament;
import com.evilvaca.loteca.scrapers.TransferMarktLeagueScraper;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import javax.annotation.PostConstruct;

@Slf4j
@NoArgsConstructor
@AllArgsConstructor
@SpringBootApplication
public class LotecaApplication {

    @Autowired
    private TransferMarktLeagueScraper scraper;

    public static void main(String[] args) {
        SpringApplication.run(LotecaApplication.class, args);
    }

    @PostConstruct
    public void postConstruct() {
        Season season = Season.builder()
                .code("2020")
                .tournament(Tournament.builder()
                        .code("BRA1")
                        .codeDescription("campeonato-brasileiro-serie-a")
                        .build())
                .build();

        System.out.println(scraper.getMatchList(Round.builder().code("32").season(season).build()));
    }
}
