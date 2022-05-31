CREATE USER 'exchange'@'localhost' IDENTIFIED BY 'exchange';

GRANT ALL PRIVILEGES ON * . * TO 'exchange'@'localhost';


# See the MySQL Reference Manual for details: 
# https://dev.mysql.com/doc/refman/8.0/en/caching-sha2-pluggable-authentication.html

ALTER USER 'exchange'@'localhost' IDENTIFIED WITH mysql_native_password BY 'exchange';