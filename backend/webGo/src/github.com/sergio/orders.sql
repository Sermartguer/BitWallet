CREATE TABLE `users_total` (
  `id_order` INT(50) NOT NULL AUTO_INCREMENT,
  `id_account` VARCHAR(50) NOT NULL,
  `amount` INT(30) NOT NULL,
  `currency` VARCHAR(30) NOT NULL,
  `price` INT(30) NOT NULL,
  `create_at` VARCHAR(30) NOT NULL,
  PRIMARY KEY (`id_order`),
  FOREIGN KEY (`id_account`) REFERENCES users(`id`) ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=latin1;