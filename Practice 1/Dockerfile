FROM php:7.4-apache

RUN apt-get update && apt-get install -y default-mysql-client
RUN docker-php-ext-install mysqli && docker-php-ext-enable mysqli
RUN a2enmod rewrite

COPY index.php /var/www/html
COPY style.css /var/www/html

COPY init.sql /docker-entrypoint-initdb.d/init.sql

EXPOSE 80

CMD ["apache2-foreground"]