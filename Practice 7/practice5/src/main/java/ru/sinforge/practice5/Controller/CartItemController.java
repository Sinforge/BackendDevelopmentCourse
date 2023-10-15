package ru.sinforge.practice5.Controller;

import jakarta.transaction.Transactional;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.User;
import org.springframework.web.bind.annotation.*;
import ru.sinforge.practice5.DTO.CartItemDTO;
import ru.sinforge.practice5.DTO.UpdateCartItemDTO;
import ru.sinforge.practice5.Entity.CartItem;
import ru.sinforge.practice5.Entity.Good;
import ru.sinforge.practice5.Repository.CartItemRepository;
import ru.sinforge.practice5.Repository.GoodRepository;
import ru.sinforge.practice5.Service.VerifyService;

import java.util.List;

@RestController
@RequestMapping("/api/cart")
public class CartItemController {
    private final VerifyService _verifyService;
    private final CartItemRepository _cartItemRepository;
    private final GoodRepository _goodRepository;
    public CartItemController(VerifyService verifyService, CartItemRepository cartItemRepository, GoodRepository goodRepository) {
        _verifyService = verifyService;
        _cartItemRepository = cartItemRepository;
        _goodRepository = goodRepository;
    }

    //1.	Добавить товар в корзину.
//        2.	Удалить товар из корзины.
//        3.	Изменить количество товара в корзине.
//        4.	Посмотреть всю корзину.
//        5.	Оформить заказ и очистить корзину.

    @PostMapping
    @PreAuthorize("isAuthenticated()")
    public ResponseEntity<String> AddNewGoodInCart(@RequestBody CartItemDTO cartItemDTO) {
        if(!_verifyService.verifyUserId(cartItemDTO.clientId)) {
            return new ResponseEntity<>(HttpStatusCode.valueOf(403));
        }

        Good good = _goodRepository.getById(cartItemDTO.goodId);
        CartItem cartItem = new CartItem();
        cartItem.setGood(good);
        cartItem.setUserId(cartItem.userId);
        cartItem.setCount(cartItemDTO.countOfGood);
        _cartItemRepository.save(cartItem);
        return new ResponseEntity<>(HttpStatusCode.valueOf(201));
    }
    @DeleteMapping
    @PreAuthorize("isAuthenticated()")
    public ResponseEntity<String> DeleteGoodFromCart(@RequestParam long cartItemId) {
        CartItem cartItem =  _cartItemRepository.findById(cartItemId).get();
        if(!_verifyService.verifyUserId(cartItem.userId)) {
            return new ResponseEntity<>(HttpStatusCode.valueOf(403));
        }
        _cartItemRepository.deleteById(cartItemId);
        return new ResponseEntity<>(HttpStatusCode.valueOf(200));
    }
    @PutMapping
    public ResponseEntity<String> UpdateGoodsCountInCart(@RequestBody UpdateCartItemDTO updateCartItemDTO) {
        CartItem cartItem =  _cartItemRepository.findById(updateCartItemDTO.id).get();
        if(!_verifyService.verifyUserId(cartItem.userId)) {
            return new ResponseEntity<>(HttpStatusCode.valueOf(403));
        }
        CartItem item = _cartItemRepository.getById(updateCartItemDTO.id);
        item.setCount(updateCartItemDTO.count);
        System.out.println(item.count + " " + updateCartItemDTO.count + " " + updateCartItemDTO.id);
        _cartItemRepository.saveAndFlush(item);
        return new ResponseEntity<>(HttpStatusCode.valueOf(200));
    }
    @GetMapping
    @PreAuthorize("isAuthenticated()")
    public ResponseEntity<List<CartItem>> GetAllGoodsInCart(@RequestParam int userId) {
        if(!_verifyService.verifyUserId(userId)) {
            return new ResponseEntity<>(HttpStatusCode.valueOf(403));
        }
        return new ResponseEntity<List<CartItem>>(((List<CartItem>)_cartItemRepository.findCartItemByUserId(userId)), HttpStatusCode.valueOf(200));
    }
    @PostMapping("/pay")
    @Transactional
    @PreAuthorize("isAuthenticated()")
    public ResponseEntity<String> PayGoodsInCart(@RequestParam int userId) {
        if(_verifyService.verifyUserId(userId)) {
            return new ResponseEntity<>(HttpStatusCode.valueOf(403));
        }
        _cartItemRepository.deleteAllByUserId(userId);
        return new ResponseEntity<>(HttpStatusCode.valueOf(200));

    }





}
