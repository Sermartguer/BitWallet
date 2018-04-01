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
  PRIMARY KEY (id)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE addrs (
  id_addrs INT(50) NOT NULL AUTO_INCREMENT,
  id_user VARCHAR(50) NOT NULL,
  address VARCHAR(50) NOT NULL,
  currency VARCHAR(30) NOT NULL,
  create_at VARCHAR(30) NOT NULL,
  PRIMARY KEY (id_addrs),
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
  send_to VARCHAR(30) NOT NULL,
  hash_id VARCHAR(50) NOT NULL,
  amount DECIMAL(30,25) NOT NULL,
  currency VARCHAR(30) NOT NULL,
  trans_time VARCHAR(30) NOT NULL,
  PRIMARY KEY (id_transactions),
  FOREIGN KEY (id_account) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB AUTO_INCREMENT=00001 DEFAULT CHARSET=latin1;

DELIMITER $$

CREATE TRIGGER add_addreses AFTER INSERT
    ON accounts
    FOR EACH ROW BEGIN
	INSERT INTO users_total (id_user,amount,currency,create_at) VALUES (NEW.id,0,'BTC',NOW());
	INSERT INTO users_total (id_user,amount,currency,create_at) VALUES (NEW.id,0,'DOGE',NOW());
	INSERT INTO users_total (id_user,amount,currency,create_at) VALUES (NEW.id,0,'LTC',NOW());
    END$$

DELIMITER ;

USE BitWallet;
DROP procedure IF EXISTS get_accounts;

DELIMITER $$
USE BitWallet$$
CREATE PROCEDURE get_accounts (IN id_account VARCHAR(50))
BEGIN
 SELECT address FROM addrs WHERE id_user = id_account;
END$$

DELIMITER ;
EOF