#!/bin/bash

#DataBase Deploy

echo "Enter Database name:"
read dbname
echo "Enter MySqlUser"
read mysqluser

mysql -u$mysqluser -p << EOF
create database $dbname;
use $dbname;
CREATE TABLE accounts (
  id VARCHAR(50) NOT NULL,
  username VARCHAR(30) NOT NULL,
  email VARCHAR(30) NOT NULL,
  password VARCHAR(100) NOT NULL,
  firstname VARCHAR(30) DEFAULT NULL,
  surname VARCHAR(30) DEFAULT NULL,
  picture VARCHAR(30) DEFAULT NULL,
  acc_type VARCHAR(30) NOT NULL,
  create_at VARCHAR(30) NOT NULL,
  update_at VARCHAR(30) NOT NULL,
  active VARCHAR(30) DEFAULT NULL,
  pin VARCHAR(5) DEFAULT NULL,
  mobile_hash VARCHAR(50) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE addresses (
  id_address INT(50) NOT NULL AUTO_INCREMENT,
  id_user VARCHAR(50) NOT NULL,
  address VARCHAR(50) DEFAULT "",
  label VARCHAR(50) DEFAULT "",
  currency VARCHAR(30) DEFAULT "",
  amount DECIMAL(30,25) DEFAULT NULL,
  create_at VARCHAR(30) DEFAULT NULL,
  active VARCHAR(10) NOT NULL,
  PRIMARY KEY (id_address),
  FOREIGN KEY (id_user) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE users_total (
  id_user_curr INT(50) NOT NULL AUTO_INCREMENT,
  id_user VARCHAR(50) NOT NULL,
  amount DECIMAL(30,25) NOT NULL,
  currency VARCHAR(30) NOT NULL,
  create_at VARCHAR(30) NOT NULL,
  PRIMARY KEY (id_user_curr),
  FOREIGN KEY (id_user) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE transactions (
  id_transactions INT(50) NOT NULL AUTO_INCREMENT,
  id_account VARCHAR(50) NOT NULL,
  send_to VARCHAR(100) NOT NULL,
  hash_id VARCHAR(100) NOT NULL,
  amount DECIMAL(30,25) NOT NULL,
  currency VARCHAR(30) NOT NULL,
  trans_type VARCHAR(10) NOT NULL,
  trans_time VARCHAR(30) NOT NULL,
  PRIMARY KEY (id_transactions),
  FOREIGN KEY (id_account) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB AUTO_INCREMENT=00001 DEFAULT CHARSET=latin1;

CREATE TABLE orders (
  id_order INT(50) NOT NULL AUTO_INCREMENT,
  id_account VARCHAR(50) NOT NULL,
  amount DECIMAL(30,25) NOT NULL,
  currency VARCHAR(30) NOT NULL,
  currency_to VARCHAR(30) NOT NULL,
  price DECIMAL(30,25) NOT NULL,
  create_at VARCHAR(30) NOT NULL,
  PRIMARY KEY (id_order),
  FOREIGN KEY (id_account) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=latin1;

CREATE TABLE activity_login (
  id_login INT(50) NOT NULL AUTO_INCREMENT,
  id_account VARCHAR(50) NOT NULL,
  ip VARCHAR(50) NOT NULL,
  time VARCHAR(50) NOT NULL,
  success VARCHAR(50) NOT NULL,
  PRIMARY KEY (id_login),
  FOREIGN KEY (id_account) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=latin1;

CREATE TABLE activity_address (
  id_actAddress INT(50) NOT NULL AUTO_INCREMENT,
  id_account VARCHAR(50) NOT NULL,
  address VARCHAR(50) NOT NULL,
  label VARCHAR(50) NOT NULL,
  time VARCHAR(50) NOT NULL,
  PRIMARY KEY (id_actAddress),
  FOREIGN KEY (id_account) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=latin1;

CREATE TABLE activity_actions (
  id_action INT(50) NOT NULL AUTO_INCREMENT,
  id_account VARCHAR(50) NOT NULL,
  amount DECIMAL(30,25) NOT NULL,
  address_local VARCHAR(50) NOT NULL,
  address_external VARCHAR(50) NOT NULL,
  currency VARCHAR(50) NOT NULL,
  time VARCHAR(50) NOT NULL,
  action VARCHAR(50) NOT NULL,
  PRIMARY KEY (id_action),
  FOREIGN KEY (id_account) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=latin1;

CREATE TABLE activity_orders (
  id_order INT(50) NOT NULL AUTO_INCREMENT,
  id_account VARCHAR(50) NOT NULL,
  currency VARCHAR(50),
  amount DECIMAL(30,25) NOT NULL,
  price DECIMAL(30,25) NOT NULL,
  action VARCHAR(10) NOT NULL,
  time VARCHAR(50) NOT NULL,
  PRIMARY KEY (id_order),
  FOREIGN KEY (id_account) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=latin1;

DELIMITER $$

CREATE TRIGGER add_addreses AFTER INSERT
    ON accounts
    FOR EACH ROW BEGIN
      INSERT INTO addresses (id_user,currency,amount,create_at,active) VALUES (NEW.id,'BTC',0,NOW(),'0');
      INSERT INTO addresses (id_user,currency,amount,create_at,active) VALUES (NEW.id,'DOGE',0,NOW(),'0');
      INSERT INTO addresses (id_user,currency,amount,create_at,active) VALUES (NEW.id,'LTC',0,NOW(),'0');
    END$$

DELIMITER ;

DELIMITER $$
CREATE TRIGGER addresses_activity AFTER UPDATE
    ON addresses
    FOR EACH ROW BEGIN
	    INSERT INTO activity_address (id_account,address,label,time) VALUES (NEW.id_user,NEW.address,NEW.label,NOW());
    END$$
DELIMITER ;

DELIMITER $$
CREATE TRIGGER order_publish_activity AFTER INSERT
    ON orders
    FOR EACH ROW BEGIN
	    INSERT INTO activity_orders (id_account,currency,amount,price,time,action) VALUES (NEW.id_account,NEW.currency,NEW.amount,NEW.price,NOW(),'Publish');
    END$$
DELIMITER ;

DELIMITER $$
CREATE TRIGGER activity_transaction_activity AFTER INSERT
    ON transactions
    FOR EACH ROW BEGIN
	    INSERT INTO activity_actions (id_account,amount,address_local,address_external,currency,TIME,ACTION) VALUES (NEW.id_account,NEW.amount,NEW.send_to,NEW.send_to,NEW.currency,NOW(),NEW.trans_type);
    END$$
DELIMITER ;


DELIMITER $$
CREATE PROCEDURE get_accounts (IN id_account VARCHAR(50))
BEGIN
 SELECT address,currency,label FROM addresses WHERE id_user = id_account;
END$$

DELIMITER ;
EOF