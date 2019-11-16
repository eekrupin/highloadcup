CREATE TABLE IF NOT EXISTS `travels`.`user` (
      `id` INT NOT NULL,
      `email` VARCHAR(200) NOT NULL,
      `first_name` VARCHAR(200) NOT NULL,
      `last_name` VARCHAR(200) NOT NULL,
      `gender` NCHAR(1) NOT NULL,
      `birth_date` DATETIME NOT NULL,
      `age` INT NOT NULL,
      PRIMARY KEY (`id`),
      UNIQUE INDEX `id_UNIQUE` (`id` ASC),
      INDEX `gender` (`gender` ASC, `age` ASC),
      INDEX `age` (`age` ASC, `gender` ASC))
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;