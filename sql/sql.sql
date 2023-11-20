CREATE DATABASE IF NOT EXISTS devbook;
DROP TABLE IF EXISTS seguidores;
USE devbook;
DROP TABLE IF EXISTS usuarios;
CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null ,
    criadoEm TIMESTAMP DEFAULT current_timestamp()
)engine=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    Foreign Key (usuario_id)
    REFERENCES usuarios(id)
    on delete CASCADE,
    
    seguidor_id int not null,
    Foreign Key (seguidor_id) 
    REFERENCES usuarios(id)
    on delete CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)

)engine=INNODB;

