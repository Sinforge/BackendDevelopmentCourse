package ru.sinforge.practice5.Controller;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import ru.sinforge.practice5.Entity.Book;
import ru.sinforge.practice5.Repository.BookRepository;

import java.util.Arrays;
import java.util.List;

@RestController
@RequestMapping("/api/book")
public class BookController {
    private final BookRepository _bookRepository;
    public BookController(BookRepository bookRepository) {
        _bookRepository = bookRepository;
    }
    @GetMapping
    public Iterable<Book> getAllBooks() {
        return _bookRepository.findAll();
    }
    @GetMapping("/{id}")
    public Book getBookById(@PathVariable int id) {
        return _bookRepository.findById(id).get();
    }
    @DeleteMapping
    public void deleteBookById(@RequestParam int id) {
        _bookRepository.deleteAllByIdInBatch(Arrays.asList(id));
    }
    @PutMapping
    public void updateBook(@RequestBody Book book) {
        Book oldBook = _bookRepository.getById(book.id);
        oldBook.setAuthor(book.author);
        oldBook.setName(book.name);
        oldBook.setPrice(book.price);
        oldBook.setSellerId(book.sellerId);
        oldBook.setProductType(book.productType);
        _bookRepository.saveAndFlush(oldBook);
    }
    @PostMapping
    public void createBook(@RequestBody Book book) {
        _bookRepository.saveAndFlush(book);
    }
}
