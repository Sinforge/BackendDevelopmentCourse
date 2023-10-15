package ru.sinforge.practice5.Service;

import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;
import ru.sinforge.practice5.DTO.JwtResponse;

import java.util.List;

@Component
class UserService implements UserDetailsService {

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        RestTemplate restTemplate = new RestTemplate();
        System.out.println("Was in user service");
        JwtResponse user = restTemplate.getForObject("http://auth:8081/api/auth/validate?jwt=" + username, JwtResponse.class);
        System.out.println(user.role + " " + user.id + " " + user.username + " ");
        if (user == null || user.id == 0 || user.role == null || user.role.equals("") || user.username == null || user.username.equals("")) {
            throw new UsernameNotFoundException("Jwt token expired or incorrect");
        }
        User userDetails = new User();
        userDetails.username = String.valueOf(user.id);
        userDetails.authorities = List.of(user.role);
        userDetails.password = null;
        return userDetails;
    }
}