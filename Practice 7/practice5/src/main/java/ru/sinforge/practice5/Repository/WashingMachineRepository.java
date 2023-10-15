package ru.sinforge.practice5.Repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import ru.sinforge.practice5.Entity.WashingMachine;

@Repository
public interface WashingMachineRepository extends JpaRepository<WashingMachine, Integer> {
}
