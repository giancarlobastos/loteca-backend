package com.evilvaca.loteca.domain;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Tournament {

    private long id;

    private String name;

    private String code;

    private String codeDescription;

    private String division;

    private String logo;
}
