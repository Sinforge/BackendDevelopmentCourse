package ru.sinforge.practice5.Repository;

import org.springframework.data.jpa.repository.JpaRepository;
import ru.sinforge.practice5.Entity.CartItem;
import ru.sinforge.practice5.Entity.Client;

public interface CartItemRepository extends JpaRepository<CartItem, Long> {
    public Iterable<CartItem> findCartItemByUser(Client client);
    public void deleteAllByUser(Client client);
}
