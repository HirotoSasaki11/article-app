drop table if exists users;
CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  first_name VARCHAR(45) NULL,
  last_name VARCHAR(45) NULL,
  email VARCHAR(45) NOT NULL,
  PRIMARY KEY (id),
    UNIQUE INDEX email_UNIQUE (email ASC));
insert into users values(1,'sasa','hiro','kd@co.jp');