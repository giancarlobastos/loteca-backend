package com.evilvaca.loteca.scrapers.domain;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;

@Data
@JsonIgnoreProperties(ignoreUnknown = true)
public class EspnData {

    private String table;

    private String activeDate;
}
