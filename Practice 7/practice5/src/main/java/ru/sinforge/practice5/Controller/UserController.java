package ru.sinforge.practice5.Controller;

import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.reactive.function.client.WebClient;

@RestController
@RequestMapping("/api/user")
public class UserController {

    @PatchMapping("/change-role")
    @PreAuthorize("hasAuthority('ADMIN')")
    public void changeUserRoleToSeller(@RequestParam int userId) {

        WebClient webClient = WebClient.create("http://auth:8081");
        webClient.patch()
                .uri("/api/auth/change-role?userId=" + userId)
                .retrieve()
                .bodyToMono(String.class)
                .block();

    }
}
