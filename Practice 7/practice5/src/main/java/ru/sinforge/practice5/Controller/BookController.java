package ru.sinforge.practice5.Controller;

import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;
import ru.sinforge.practice5.Entity.Book;
import ru.sinforge.practice5.Repository.BookRepository;
import ru.sinforge.practice5.Service.VerifyService;

import java.util.List;

@RestController
@RequestMapping("/api/book")
public class BookController {
    private final VerifyService _verifyService;
    private final BookRepository _bookRepository;
    public BookController(VerifyService verifyService, BookRepository bookRepository) {
        _verifyService = verifyService;
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

    @PutMapping
    @PreAuthorize("hasAnyAuthority('SELLER', 'ADMIN')")
    public ResponseEntity<String> updateBook(@RequestBody Book book) {
        Book oldBook = _bookRepository.getById(book.id);

        if(!_verifyService.verifyRole(oldBook.sellerId))
            return new ResponseEntity<>("You not owner of this book", HttpStatusCode.valueOf(401));

        oldBook.setAuthor(book.author);
        oldBook.setName(book.name);
        oldBook.setPrice(book.price);
        oldBook.setSellerId(book.sellerId);
        oldBook.setProductType(book.productType);
        _bookRepository.saveAndFlush(oldBook);
        return new ResponseEntity<>( HttpStatusCode.valueOf(200));

    }
    @PostMapping
    @PreAuthorize("hasAnyAuthority('SELLER', 'ADMIN')")
    public ResponseEntity<String> createBook(@RequestBody Book book) {
        if(!_verifyService.verifyRole(book.sellerId))
            return new ResponseEntity<>("You not owner of this book", HttpStatusCode.valueOf(401));

        _bookRepository.saveAndFlush(book);
        return new ResponseEntity<>("Succesful added", HttpStatusCode.valueOf(201));
    }

}
