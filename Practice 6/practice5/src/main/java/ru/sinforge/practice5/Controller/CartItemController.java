package ru.sinforge.practice5.Controller;

import jakarta.transaction.Transactional;
import org.springframework.web.bind.annotation.*;
import ru.sinforge.practice5.DTO.CartItemDTO;
import ru.sinforge.practice5.DTO.UpdateCartItemDTO;
import ru.sinforge.practice5.Entity.CartItem;
import ru.sinforge.practice5.Entity.Client;
import ru.sinforge.practice5.Entity.Good;
import ru.sinforge.practice5.Repository.CartItemRepository;
import ru.sinforge.practice5.Repository.ClientRepository;
import ru.sinforge.practice5.Repository.GoodRepository;

import java.util.List;

@RestController
@RequestMapping("/api/cart")
public class CartItemController {
    private final CartItemRepository _cartItemRepository;
    private final GoodRepository _goodRepository;
    private final ClientRepository _clientRepository;
    public CartItemController(CartItemRepository cartItemRepository, ClientRepository clientRepository, GoodRepository goodRepository) {
        _cartItemRepository = cartItemRepository;
        _goodRepository = goodRepository;
        _clientRepository = clientRepository;
    }

    //1.	Добавить товар в корзину.
//        2.	Удалить товар из корзины.
//        3.	Изменить количество товара в корзине.
//        4.	Посмотреть всю корзину.
//        5.	Оформить заказ и очистить корзину.

    @PostMapping
    public void AddNewGoodInCart(@RequestBody CartItemDTO cartItemDTO) {
        Client client = _clientRepository.getById(cartItemDTO.clientId);
        Good good = _goodRepository.getById(cartItemDTO.goodId);
        CartItem cartItem = new CartItem();
        cartItem.setGood(good);
        cartItem.setUser(client);
        cartItem.setCount(cartItemDTO.countOfGood);
        _cartItemRepository.save(cartItem);
    }
    @DeleteMapping
    public void DeleteGoodFromCart(@RequestParam long cartItemId) {
        _cartItemRepository.deleteById(cartItemId);
    }
    @PutMapping
    public void UpdateGoodsCountInCart(@RequestBody UpdateCartItemDTO updateCartItemDTO) {
        CartItem item = _cartItemRepository.getById(updateCartItemDTO.id);
        item.setCount(updateCartItemDTO.count);
        System.out.println(item.count + " " + updateCartItemDTO.count + " " + updateCartItemDTO.id);
        _cartItemRepository.saveAndFlush(item);
    }
    @GetMapping
    public List<CartItem> GetAllGoodsInCart(@RequestParam int userId) {
        return (List<CartItem>)_cartItemRepository.findCartItemByUser(_clientRepository.getById(userId));
    }
    @PostMapping("/pay")
    @Transactional
    public void PayGoodsInCart(@RequestParam int userId) {
        _cartItemRepository.deleteAllByUser(_clientRepository.getById(userId));
    }





}
