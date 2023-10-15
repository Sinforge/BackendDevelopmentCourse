package com.example.auth.Repository;


import com.example.auth.Entity.Client;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AuthRepository extends JpaRepository<Client, Integer> {
    Client getClientByLoginAndPassword(String login, String password);
}