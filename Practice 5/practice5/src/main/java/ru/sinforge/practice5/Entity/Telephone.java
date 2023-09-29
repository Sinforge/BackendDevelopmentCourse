package ru.sinforge.practice5.Entity;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import lombok.*;

@Getter
@Setter
@Entity
public class Telephone {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    public int id;
    public String producer;
    public double batteryCapacity;
    public int sellerId;
    public String productType;
    public int price;
    public String name;
}
