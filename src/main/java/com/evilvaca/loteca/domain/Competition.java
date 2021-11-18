package com.evilvaca.loteca.domain;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Competition {

    private long id;

    private String name;

    private String code;

    private String codeName;

    private String division;

    private String logo;
}
