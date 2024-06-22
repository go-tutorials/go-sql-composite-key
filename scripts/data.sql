create table if not exists company_users (
  company_id varchar(40) not null,
  user_id varchar(40) not null,
  username varchar(120),
  email varchar(120),
  phone varchar(45),
  date_of_birth date,
  primary key (id)
);

insert into company_users (company_id, user_id, username, email, phone, date_of_birth) values ('c1', 'ironman', 'tony.stark', 'tony.stark@gmail.com', '0987654321', '1963-03-25');
insert into company_users (company_id, user_id, username, email, phone, date_of_birth) values ('c2', 'spiderman', 'peter.parker', 'peter.parker@gmail.com', '0987654321', '1962-08-25');
insert into company_users (company_id, user_id, username, email, phone, date_of_birth) values ('c1', 'wolverine', 'james.howlett', 'james.howlett@gmail.com', '0987654321', '1974-11-16');
insert into company_users (company_id, user_id, username, email, phone, date_of_birth) values ('c2', 'wolverine', 'james.howlett', 'james.howlett@gmail.com', '0987654321', '1974-11-16');
