create schema if not exists `trade-shop`;

create table if not exists `trade-shop`.inventory
(
  user_id char(36) null,
  item_id char(36) null,
  count   int      null
);

create table if not exists `trade-shop`.item_sale
(
  sale_id char(36)        not null,
  item_id char(36)        not null,
  count   int             null,
  price   double unsigned not null
);

create index item_sale_items_id_fk
  on `trade-shop`.item_sale (item_id);

create index item_sale_sales_id_fk
  on `trade-shop`.item_sale (sale_id);

create table if not exists `trade-shop`.items
(
  id   char(36)     null,
  name varchar(255) null,
  constraint items_id_uindex
  unique (id)
);

create table if not exists `trade-shop`.sales
(
  id      char(36) null,
  user_id char(36) null,
  constraint sales_id_uindex
  unique (id)
);

create table if not exists `trade-shop`.users
(
  id       char(36)                  not null,
  email    varchar(255) charset utf8 not null,
  password varchar(255) charset utf8 null,
  bill     double                    not null,
  constraint users_id_uindex
  unique (id)
);

alter table `trade-shop`.users
  add primary key (id);


