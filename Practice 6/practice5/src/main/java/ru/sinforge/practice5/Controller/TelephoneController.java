package ru.sinforge.practice5.Controller;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import ru.sinforge.practice5.Entity.Book;
import ru.sinforge.practice5.Entity.Telephone;
import ru.sinforge.practice5.Repository.TelephoneRepository;

import java.util.Arrays;

@RestController
@RequestMapping("/api/telephone")
public class TelephoneController {
    private final TelephoneRepository _telephoneRepository;
    public TelephoneController(TelephoneRepository telephoneRepository) {
        _telephoneRepository = telephoneRepository;
    }
    @GetMapping
    public Iterable<Telephone> getAllTelephones() {
        return _telephoneRepository.findAll();
    }
    @GetMapping("/{id}")
    public Telephone getTelephoneById(@PathVariable int id) {
        return _telephoneRepository.findById(id).get();
    }
    @DeleteMapping
    public void deleteTelephoneById(@RequestParam int id) {
        _telephoneRepository.deleteAllByIdInBatch(Arrays.asList(id));
    }
    @PutMapping
    public void updateTelephone(@RequestBody Telephone telephone) {
        Telephone oldTelephone = _telephoneRepository.getById(telephone.id);
        oldTelephone.setProducer(telephone.getProducer());
        oldTelephone.setName(telephone.name);
        oldTelephone.setPrice(telephone.price);
        oldTelephone.setSellerId(telephone.sellerId);
        oldTelephone.setProductType(telephone.productType);
        oldTelephone.setBatteryCapacity(telephone.batteryCapacity);
        _telephoneRepository.saveAndFlush(oldTelephone);
    }
    @PostMapping
    public void createTelephone(@RequestBody Telephone telephone) {
        _telephoneRepository.saveAndFlush(telephone);
    }
}
