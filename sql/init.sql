create table if not exists users(
id bigint primary key auto_increment,
full_name varchar(256) not null unique,
phone varchar(256) not null unique,
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp,
on delete cascade on update cascade
)

engine = InnoDB
default charset = utf8mb4;

create table if not exists locations(
id bigint primary key auto_increment,
location_detail varchar(256) not null unique,
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp,
on delete cascade on update cascade
)

engine = InnoDB
default charset = utf8mb4;

create table if not exists location_users(
id bigint primary key auto_increment,
location_id bigint not null,
users_id bigint not null,
constraint location_users_location_id foreign key(location_id) references locations(id),
constraint location_users_user_id foreign key(user_id) references users(id),
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp,
on delete cascade on update cascade
)

engine = InnoDB
default charset = utf8mb4;

create table if not exists products(
id bigint primary key auto_increment,
file varchar(256) not null unique,
cost decimal(10,2) not null default 0.0,
constraint products_productorder_id foreign key(productorder_id) references productsorder(id),
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp,
on delete cascade on update cascade
)

engine = InnoDB
default charset = utf8mb4;

create table if not exists productsorder(
id bigint primary key auto_increment,
detail varchar(512) not null unique,
sumcost decimal(10,2) not null default 0.0,
quantity int(10) unsigned default 0,
status char(1) default 0,
-- product_id bigint not null,
users_id bigint not null,
location_id bigint not null,
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp,
-- constraint productsorder_users_id foreign key(user_id) references users(id),
-- constraint productsorder_location_id foreign key(location_id) references locations(id),
constraint productsorder_product_id foreign key(product_id) references products(id)
on delete cascade on update cascade
)

engine = InnoDB
default charset = utf8mb4;


-- CREATE TABLE `product_order` (
--   `id` int PRIMARY KEY AUTO_INCREMENT,
--   `users_id` int,
--   `detail` varchar(255),
--   `location_id` int,
--   `status` varchar(255),
--   `sum_cost` float,
--   `quantity` int,
--   `created_at` timestamp NOT NULL DEFAULT (now()),
--   `updated_at` timestamp NOT NULL DEFAULT (now())
-- );

-- CREATE TABLE `product` (
--   `id` int PRIMARY KEY AUTO_INCREMENT,
--   `product_order_id` int,
--   `file` varchar(255) NOT NULL,
--   `cost` float NOT NULL,
--   `created_at` timestamp NOT NULL DEFAULT (now()),
--   `updated_at` timestamp NOT NULL DEFAULT (now())
-- );

-- CREATE TABLE `location` (
--   `id` int PRIMARY KEY AUTO_INCREMENT,
--   `location_detail` varchar(255),
--   `created_at` timestamp NOT NULL DEFAULT (now()),
--   `updated_at` timestamp NOT NULL DEFAULT (now())
-- );

-- CREATE TABLE `location_users` (
--   `id` int PRIMARY KEY AUTO_INCREMENT,
--   `location_id` int,
--   `users_id` int,
--   `created_at` timestamp NOT NULL DEFAULT (now()),
--   `updated_at` timestamp NOT NULL DEFAULT (now())
-- );

-- ALTER TABLE `product_order` ADD FOREIGN KEY (`users_id`) REFERENCES `users` (`id`);

-- ALTER TABLE `product` ADD FOREIGN KEY (`product_order_id`) REFERENCES `product_order` (`id`);

-- ALTER TABLE `location_users` ADD FOREIGN KEY (`users_id`) REFERENCES `users` (`id`);

-- ALTER TABLE `location_users` ADD FOREIGN KEY (`location_id`) REFERENCES `location` (`id`);
