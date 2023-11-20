insert into usuarios (nome,nick,email,senha)
values
("User1","user1","user1@gmail.com", "$2a$10$dIktY21ypfnIrD9aiOj2vOvrSXnPkDSaBuUe9T6zd5PA/xm5k57NS"),
("User2","user2","user2@gmail.com", "$2a$10$dIktY21ypfnIrD9aiOj2vOvrSXnPkDSaBuUe9T6zd5PA/xm5k57NS"),
("User3","user3","user3@gmail.com", "$2a$10$dIktY21ypfnIrD9aiOj2vOvrSXnPkDSaBuUe9T6zd5PA/xm5k57NS"),
("User4","user4","user4@gmail.com", "$2a$10$dIktY21ypfnIrD9aiOj2vOvrSXnPkDSaBuUe9T6zd5PA/xm5k57NS");


insert into seguidores (usuario_id, seguidor_id)
values
(1,2),
(3,1),
(1,3);