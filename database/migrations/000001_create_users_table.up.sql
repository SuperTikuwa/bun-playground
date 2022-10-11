create table if not exists users (
  id int not null auto_increment primary key,
  name varchar(255) not null,
  email varchar(255) not null,
  cognito_uuid varchar(255),
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp
);