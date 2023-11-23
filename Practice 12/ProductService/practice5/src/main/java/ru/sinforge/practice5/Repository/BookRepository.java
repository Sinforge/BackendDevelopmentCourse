package ru.sinforge.practice5.Repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import ru.sinforge.practice5.Entity.Book;

@Repository
public interface BookRepository extends JpaRepository<Book, Integer> {
}
