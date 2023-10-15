package ru.sinforge.practice5.Controller;

import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.User;
import org.springframework.web.bind.annotation.*;
import ru.sinforge.practice5.Entity.WashingMachine;
import ru.sinforge.practice5.Repository.WashingMachineRepository;
import ru.sinforge.practice5.Service.VerifyService;

import java.util.List;

@RestController
@RequestMapping("/api/washing_machine")
public class WashingMachineController {
    private final VerifyService _verifyService;
    private final WashingMachineRepository _washingMachineRepository;
    public WashingMachineController(VerifyService verifyService, WashingMachineRepository washingMachineRepository) {
        _verifyService = verifyService;
        _washingMachineRepository = washingMachineRepository;
    }
    @GetMapping
    public Iterable<WashingMachine> getAllWashingMachines() {
        return _washingMachineRepository.findAll();
    }
    @GetMapping("/{id}")
    public WashingMachine getWashingMachineById(@PathVariable int id) {
        return _washingMachineRepository.findById(id).get();
    }

    @PutMapping
    @PreAuthorize("hasAnyAuthority('SELLER', 'ADMIN')")
    public ResponseEntity<String> updateWashingMachine(@RequestBody WashingMachine wm) {
        WashingMachine oldWM = _washingMachineRepository.getById(wm.id);
        if(!_verifyService.verifyRole(oldWM.sellerId)) {
            return new ResponseEntity<>("You not owner of this washing machine", HttpStatusCode.valueOf(403));
        }
        oldWM.setName(wm.name);
        oldWM.setPrice(wm.price);
        oldWM.setSellerId(wm.sellerId);
        oldWM.setProducer(wm.getProducer());
        oldWM.setTankCapacity(wm.getTankCapacity());
        oldWM.setProductType(wm.productType);
        _washingMachineRepository.saveAndFlush(wm);
        return new ResponseEntity<>("Successful updated", HttpStatusCode.valueOf(200));
    }
    @PostMapping
    @PreAuthorize("hasAnyAuthority('SELLER', 'ADMIN')")
    public ResponseEntity<String> createBook(@RequestBody WashingMachine wm)
    {

        if(!_verifyService.verifyRole(wm.sellerId)) {
            return new ResponseEntity<>("You not owner of this washin machine", HttpStatusCode.valueOf(403));
        }
        _washingMachineRepository.saveAndFlush(wm);
        return new ResponseEntity<>("Successful added", HttpStatusCode.valueOf(200));
    }
}
