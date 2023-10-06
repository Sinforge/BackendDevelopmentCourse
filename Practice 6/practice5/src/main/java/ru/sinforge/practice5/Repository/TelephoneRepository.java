package ru.sinforge.practice5.Repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import ru.sinforge.practice5.Entity.Telephone;

@Repository
public interface TelephoneRepository extends JpaRepository<Telephone, Integer> {
}
