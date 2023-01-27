create table if not exists users(
    id bigint primary key auto_increment,
    name varchar(200) not null,
    cpf varchar(12) not null ,
    email varchar(150),
    password varchar(300),
    user_type ENUM('F', 'J'),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    unique(id, cpf, email)
)
engine = InnoDB
default charset = utf8;