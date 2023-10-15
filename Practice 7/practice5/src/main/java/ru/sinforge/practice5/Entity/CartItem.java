package ru.sinforge.practice5.Entity;

import com.fasterxml.jackson.annotation.JsonIgnore;
import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@Entity
public class CartItem {


    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    public Long id;
    public int userId;
    @OneToOne
    public Good good;
    public int count;




}
