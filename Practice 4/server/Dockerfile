FROM php:7.4-apache

RUN apt-get update && apt-get install -y default-mysql-client
RUN docker-php-ext-install mysqli && docker-php-ext-enable mysqli
RUN a2enmod headers && sed -ri -e 's/^([ \t]*)(<\/VirtualHost>)/\1\tHeader set Access-Control-Allow-Origin "*"\n\1\2/g' /etc/apache2/sites-available/*.conf




COPY group.php /var/www/html
COPY student.php /var/www/html


    
EXPOSE 80

CMD ["apache2-foreground"]