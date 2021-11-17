package com.evilvaca.loteca.domain;

import lombok.Builder;
import lombok.Data;

import java.time.LocalDateTime;

@Data
@Builder
public class Match {

    private long id;

    private String season;

    private String round;

    private long homeId;

    private long awayId;

    private String stadium;

    private LocalDateTime startAt;

    private Integer homeScore;

    private Integer awayScore;
}
