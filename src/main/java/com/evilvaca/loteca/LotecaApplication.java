package com.evilvaca.loteca;

import com.evilvaca.loteca.domain.Competition;
import com.evilvaca.loteca.domain.Round;
import com.evilvaca.loteca.domain.Season;
import com.evilvaca.loteca.scrapers.TransferMarktCupScraper;
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
    private TransferMarktLeagueScraper leagueScraper;

    @Autowired
    private TransferMarktCupScraper cupScraper;

    public static void main(String[] args) {
        SpringApplication.run(LotecaApplication.class, args);
    }

    @PostConstruct
    public void postConstruct() {
        Season season = Season.builder()
                .code("2020")
                .competition(Competition.builder()
                        .code("BRA1")
                        .codeName("campeonato-brasileiro-serie-a")
                        .build())
                .build();

//        Season season = Season.builder()
//                .code("2020")
//                .competition(Competition.builder()
//                        .code("CL")
//                        .codeName("uefa-champions-league")
//                        .build())
//                .build();

        System.out.println(leagueScraper.getMatchList(Round.builder().code("10").season(season).build()));
//        System.out.println(cupScraper.getMatchList(season));
    }
}
