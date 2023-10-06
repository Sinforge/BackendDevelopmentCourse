package ru.sinforge.practice5.Entity;


//1.	Добавить товар в корзину.
//        2.	Удалить товар из корзины.
//        3.	Изменить количество товара в корзине.
//        4.	Посмотреть всю корзину.
//        5.	Оформить заказ и очистить корзину.

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@Entity
public class Good {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    public int id;
    public int sellerId;
    public String productType;
    public int price;
    public String name;
}
