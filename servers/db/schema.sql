create table if not exists users (
    userid int not null auto_increment primary key,
    email varchar(255) not null,
    passhash binary(72) not null,
    unique (username),
    unique(email)
);