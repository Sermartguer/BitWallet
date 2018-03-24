CREATE TABLE `addrs` (
  `id_addrs` VARCHAR(50) NOT NULL,
  `id_user` VARCHAR(50) NOT NULL,
  `address` VARCHAR(30) NOT NULL,
  `amount` VARCHAR(30) NOT NULL,
  `currency` VARCHAR(30) NOT NULL,
  `create_at` VARCHAR(30) NOT NULL,
  PRIMARY KEY (`id_addrs`),
  FOREIGN KEY (`id_user`) REFERENCES users(`id`) ON DELETE CASCADE
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;