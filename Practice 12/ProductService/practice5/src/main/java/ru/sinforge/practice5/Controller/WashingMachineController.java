package ru.sinforge.practice5.Controller;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import ru.sinforge.practice5.Entity.Book;
import ru.sinforge.practice5.Entity.WashingMachine;
import ru.sinforge.practice5.Repository.WashingMachineRepository;

import java.util.Arrays;

@RestController
@RequestMapping("/api/washing_machine")
public class WashingMachineController {
    private final WashingMachineRepository _washingMachineRepository;
    public WashingMachineController(WashingMachineRepository washingMachineRepository) {
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
    @DeleteMapping
    public void deleteWashingMachineById(@RequestParam int id) {
        _washingMachineRepository.deleteAllByIdInBatch(Arrays.asList(id));
    }
    @PutMapping
    public void updateWashingMachine(@RequestBody WashingMachine wm) {
        WashingMachine oldWM = _washingMachineRepository.getById(wm.id);
        oldWM.setName(wm.name);
        oldWM.setPrice(wm.price);
        oldWM.setSellerId(wm.sellerId);
        oldWM.setProducer(wm.getProducer());
        oldWM.setTankCapacity(wm.getTankCapacity());
        oldWM.setProductType(wm.productType);
        _washingMachineRepository.saveAndFlush(wm);
    }
    @PostMapping
    public void createBook(@RequestBody WashingMachine wm) {
        _washingMachineRepository.saveAndFlush(wm);
    }
}
