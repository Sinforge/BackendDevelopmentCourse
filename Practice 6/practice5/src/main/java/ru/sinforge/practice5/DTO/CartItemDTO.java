package ru.sinforge.practice5.DTO;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CartItemDTO {
    public int clientId;
    public int goodId;
    public int countOfGood;
}