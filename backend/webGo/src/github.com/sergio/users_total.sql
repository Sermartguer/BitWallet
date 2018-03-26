CREATE TABLE `users_total` (
  `id_user_curr` INT(50) NOT NULL AUTO_INCREMENT,
  `id_user` VARCHAR(50) NOT NULL,
  `amount` INT(30) NOT NULL,
  `currency` VARCHAR(30) NOT NULL,
  `create_at` VARCHAR(30) NOT NULL,
  PRIMARY KEY (`id_user_curr`),
  FOREIGN KEY (`id_user`) REFERENCES users(`id`) ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=latin1;