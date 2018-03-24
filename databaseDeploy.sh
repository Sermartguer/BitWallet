#!/bin/bash

#DataBase Deploy

echo "Enter Database name:"
read dbname
echo "Enter MySqlUser"
read mysquser

mysql -u$mysquser -p << EOF
create database $dbname;
use $dbname;
CREATE TABLE users (
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
  id_addrs VARCHAR(50) NOT NULL,
  id_user VARCHAR(50) NOT NULL,
  address VARCHAR(30) NOT NULL,
  amount VARCHAR(30) NOT NULL,
  currency VARCHAR(30) NOT NULL,
  create_at VARCHAR(30) NOT NULL,
  PRIMARY KEY (id_addrs),
  FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
EOF