use tgbot;
create table users (
user_id  integer (10),
user_name varchar (30),
id_chat integer (10),
subscription date,
subscription_longtime integer (2),
primary key (user_id)

);
create table tests (
user_id  integer (10),
testNow integer (2),
testEnd BIT (1),
primary key (user_id)

);