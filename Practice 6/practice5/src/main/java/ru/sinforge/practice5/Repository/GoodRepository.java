package ru.sinforge.practice5.Repository;

import org.springframework.data.jpa.repository.JpaRepository;
import ru.sinforge.practice5.Entity.Good;

public interface GoodRepository extends JpaRepository<Good, Integer> {

}
