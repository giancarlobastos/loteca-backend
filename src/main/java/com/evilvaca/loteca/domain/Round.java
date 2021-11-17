package com.evilvaca.loteca.domain;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Round {

    private long id;

    private String name;

    private Integer number;

    private String code;

    private Season season;
}
