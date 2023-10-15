package ru.sinforge.practice5.Service;

import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class VerifyService {

    public boolean verifyRole(int sellerId) {
        Authentication auth = getAuth();
        String authority = "";
        List<? extends GrantedAuthority> authorities =  auth.getAuthorities().stream().toList();
        for(int i=0 ;i < authorities.size(); i++) {
            authority = authorities.get(i).getAuthority();
        }
        if(sellerId != Integer.parseInt(auth.getName()) && !authority.equals("ADMIN")) {
            return false;
        }
        return true;
    }
    public boolean verifyUserId(int userId) {
        Authentication auth = getAuth();
        if(userId != Integer.parseInt(auth.getName())) {
            return false;
        }
        return true;
    }
    private Authentication getAuth() {
        return SecurityContextHolder.getContext().getAuthentication();
    }
}
