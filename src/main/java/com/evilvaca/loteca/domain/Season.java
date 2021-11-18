package com.evilvaca.loteca.domain;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Season {

    private long id;

    private String name;

    private String code;

    private Competition competition;
}
