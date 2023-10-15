package ru.sinforge.practice5.Controller;

import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.User;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import ru.sinforge.practice5.Entity.Good;
import ru.sinforge.practice5.Repository.GoodRepository;
import ru.sinforge.practice5.Service.VerifyService;

import java.util.Arrays;
import java.util.List;

@RestController
@RequestMapping("/api/good")
public class GoodController {
    private final GoodRepository _goodRepository;
    private final VerifyService _verifyService;
    public GoodController(GoodRepository goodRepository, VerifyService verifyService) {
        _goodRepository = goodRepository;
        _verifyService = verifyService;
    }
    @DeleteMapping
    @PreAuthorize("hasAnyAuthority('SELLER', 'ADMIN')")
    public ResponseEntity<String> deleteGoodById(@RequestParam int id, @AuthenticationPrincipal User user) {
        Good good =_goodRepository.findById(id).get();
        if(!_verifyService.verifyRole(good.sellerId)) {
            return new ResponseEntity<String>("You not owner of this good", HttpStatusCode.valueOf(401));
        }
        _goodRepository.deleteAllByIdInBatch(Arrays.asList(id));
        return new ResponseEntity<>("Successful delete", HttpStatusCode.valueOf(200));
    }
}
