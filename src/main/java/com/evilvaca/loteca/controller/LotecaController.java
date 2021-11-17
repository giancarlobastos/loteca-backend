package com.evilvaca.loteca.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
public class LotecaController {

    @PostMapping("/accounts")
    public ResponseEntity<?> createAccount() {
        return ResponseEntity.ok()
                .build();
    }
}
