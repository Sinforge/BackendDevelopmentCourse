package ru.sinforge.practice5.Config;

import org.springframework.context.annotation.Configuration;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@Configuration
@EnableJpaRepositories(basePackages = {
        "ru.sinforge.practice5.Repository"
})
public class DbConfig {
}