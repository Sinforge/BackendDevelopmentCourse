package com.example.auth.Service;

import com.example.auth.Entity.Client;
import io.jsonwebtoken.*;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.security.Keys;
import org.springframework.stereotype.Service;

import javax.crypto.SecretKey;
import java.time.Instant;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.Date;

@Service
public class JwtProvider {
    private SecretKey jwtAccessSecret;

    public JwtProvider() {
        jwtAccessSecret = Keys.hmacShaKeyFor(Decoders.BASE64.decode("verysecretsecret1secretsupersecretkeyaaaaaaaaaaaaaaaaaaaaaa"));

    }

    public String generateToken(Client client) {
        final LocalDateTime now = LocalDateTime.now();
        final Instant accessExpirationInstant = now.plusMinutes(10).atZone(ZoneId.systemDefault()).toInstant();
        final Date accessExpiration = Date.from(accessExpirationInstant);
        String str = Jwts.builder()
                .setSubject(client.login)
                .setExpiration(accessExpiration)
                .signWith(jwtAccessSecret)
                .claim("role", client.role)
                .claim("username", client.getLogin())
                .claim("id", client.id)
                .compact();
        System.out.println(Jwts.parserBuilder()
                .setSigningKey(jwtAccessSecret)
                .build()
                .parse(str)
                .getBody().toString()
        );
        return str;

    }
    public Claims validateToken(String token) {
        try {
            Object claims = Jwts.parserBuilder()
                    .setSigningKey(jwtAccessSecret)
                    .build()
                    .parse(token)
                    .getBody();
            io.jsonwebtoken.Claims cl = (io.jsonwebtoken.Claims) claims;

            return cl;
        }
        catch (ExpiredJwtException expEx) {
            System.out.println("Token expired" + expEx);
            throw expEx;
        } catch (UnsupportedJwtException unsEx) {
            System.out.println("Unsupported jwt" +unsEx);
        } catch (MalformedJwtException mjEx) {
            System.out.println("Malformed jwt" + mjEx);
        } catch (SignatureException sEx) {
            System.out.println("Invalid signature" +sEx);
        } catch (Exception e) {
            System.out.println("invalid token" + e);
        }
        return null;
    }



}