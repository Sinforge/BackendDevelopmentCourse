package ru.sinforge.practice5.Controller;

import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.User;
import org.springframework.web.bind.annotation.*;
import ru.sinforge.practice5.Entity.Telephone;
import ru.sinforge.practice5.Repository.TelephoneRepository;
import ru.sinforge.practice5.Service.VerifyService;

import java.util.List;

@RestController
@RequestMapping("/api/telephone")
public class TelephoneController {
    private final TelephoneRepository _telephoneRepository;
    private final VerifyService _verifyService;
    public TelephoneController(TelephoneRepository telephoneRepository, VerifyService verifyService) {
        _telephoneRepository = telephoneRepository;
        _verifyService = verifyService;
    }
    @GetMapping
    public Iterable<Telephone> getAllTelephones() {
        return _telephoneRepository.findAll();
    }
    @GetMapping("/{id}")
    public Telephone getTelephoneById(@PathVariable int id) {
        return _telephoneRepository.findById(id).get();
    }

    @PutMapping
    @PreAuthorize("hasAnyAuthority('SELLER', 'ADMIN')")
    public ResponseEntity<String> updateTelephone(@RequestBody Telephone telephone) {
        Telephone oldTelephone = _telephoneRepository.getById(telephone.id);
        if(!_verifyService.verifyRole(oldTelephone.sellerId)) {
            return new ResponseEntity<>("You not owner of this telephone", HttpStatusCode.valueOf(401));
        }

        oldTelephone.setProducer(telephone.getProducer());
        oldTelephone.setName(telephone.name);
        oldTelephone.setPrice(telephone.price);
        oldTelephone.setSellerId(telephone.sellerId);
        oldTelephone.setProductType(telephone.productType);
        oldTelephone.setBatteryCapacity(telephone.batteryCapacity);
        _telephoneRepository.saveAndFlush(oldTelephone);
        return new ResponseEntity<>("Successful update", HttpStatusCode.valueOf(200));
    }
    @PostMapping
    @PreAuthorize("hasAnyAuthority('SELLER', 'ADMIN')")
    public ResponseEntity<String> createTelephone(@RequestBody Telephone telephone) {
        if(!_verifyService.verifyRole(telephone.sellerId)) {
            return new ResponseEntity<>("You not owner of this telephone", HttpStatusCode.valueOf(401));
        }
        _telephoneRepository.saveAndFlush(telephone);
        return new ResponseEntity<>("Successful added", HttpStatusCode.valueOf(201));
    }
}
