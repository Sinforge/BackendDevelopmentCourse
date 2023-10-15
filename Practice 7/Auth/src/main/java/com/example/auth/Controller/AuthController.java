package com.example.auth.Controller;

import com.example.auth.Entity.Client;
import com.example.auth.Entity.JwtResponse;
import com.example.auth.Repository.AuthRepository;
import com.example.auth.Service.JwtProvider;
import io.jsonwebtoken.Claims;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/auth")
public class AuthController {
    private final JwtProvider _jwtProvider;
    private final AuthRepository _authRepository;
    @GetMapping("/login")
    public ResponseEntity<String> login(@RequestParam String username, @RequestParam String password) {
        Client client =_authRepository.getClientByLoginAndPassword(username, password);
        if(client != null) {
            return new ResponseEntity<>(_jwtProvider.generateToken(client), HttpStatusCode.valueOf(200));
        }
        return new ResponseEntity<>(HttpStatusCode.valueOf(400));

    }
    @PostMapping("/reg")
    public String register(@RequestBody Client client) {
        _authRepository.save(client);
        return "SUCCESS";
    }
    @GetMapping("/validate")
    public JwtResponse validateToken(@RequestParam String jwt) {
        Claims claims = _jwtProvider.validateToken(jwt);
        JwtResponse jwtResponse = new JwtResponse();
        jwtResponse.id = claims.get("id", Integer.class);
        jwtResponse.role = claims.get("role", String.class);
        jwtResponse.username = claims.get("username", String.class);
        return jwtResponse;
    }
    @PatchMapping("/change-role")
    public String changeUserRoleToSeller(@RequestParam int userId) {
        System.out.println("I was here");
        Client client = _authRepository.findById(userId).get();
        client.role = "SELLER";
        _authRepository.saveAndFlush(client);
        return "SUCCESS";
    }



}