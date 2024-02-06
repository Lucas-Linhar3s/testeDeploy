CREATE DATABASE IF NOT EXISTS proje;

USE proje;

DROP TABLE IF EXISTS produtos;

CREATE TABLE produtos(
    id int auto_increment primary key,
    produto varchar(255) not null unique,
    valor float not null,
    quantidade int not null,
    dt_cadastro date,
    dt_atualizacao TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=INNODB;