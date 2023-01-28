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
create table if not exists wallet(
    id bigint primary key auto_increment,
    user_id bigint not null,
    value decimal(12,2) default 0.00,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    unique(id)
)
engine = InnoDB
default charset = utf8;