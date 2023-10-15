package ru.sinforge.practice5.Repository;

import org.springframework.data.jpa.repository.JpaRepository;
import ru.sinforge.practice5.Entity.CartItem;

public interface CartItemRepository extends JpaRepository<CartItem, Long> {
    public Iterable<CartItem> findCartItemByUserId(int clientId);
    public void deleteAllByUserId(int clientId);
}
