
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

CREATE TABLE IF NOT EXISTS `travels`.`location` (
       `id` INT NOT NULL,
       `place` VARCHAR(200) NOT NULL,
       `country` VARCHAR(200) NOT NULL,
       `city` VARCHAR(200) NOT NULL,
       `distance` INT NOT NULL,
       PRIMARY KEY (`id`),
       UNIQUE INDEX `id_UNIQUE` (`id` ASC),
       INDEX `place` (`place` ASC),
       INDEX `country` (`country` ASC, `city` ASC),
       INDEX `distance` (`distance` ASC, `country` ASC, `city` ASC))
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS`travels`.`visit` (
       `id` INT NOT NULL,
       `location` INT NOT NULL,
       `user` INT NOT NULL,
       `visited_at` DATETIME NOT NULL,
       `mark` INT NOT NULL,
       PRIMARY KEY (`id`),
       UNIQUE INDEX `id_UNIQUE` (`id` ASC),
       INDEX `user_v_l` (`user` ASC, `visited_at` ASC, `location` ASC),
       INDEX `user_c` (`user` ASC, `location` ASC),
       INDEX `user_d` (`user` ASC, `visited_at` ASC))
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;