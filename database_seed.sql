CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "reffered_user_id" int,
  "email" varchar(256) NOT NULL,
  "password" varchar(256) NOT NULL,
  "long_name" varchar(256) NOT NULL,
  "phone" varchar(256) NOT NULL,
  "role" varchar(20) NOT NULL DEFAULT (USER),
  "balance" int NOT NULL,
  "photo" varchar(256),
  "refferal_code" varchar NOT NULL,
  "is_complete_bonus" boolean NOT NULL DEFAULT (false),
  "is_complete_bonus_reff" boolean NOT NULL DEFAULT (false),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "addresses" (
  "id" SERIAL PRIMARY KEY,
  "recipient_name" varchar(256) NOT NULL,
  "full_address" text NOT NULL,
  "recipient_phone" varchar(50) NOT NULL,
  "user_id" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(256) NOT NULL,
  "description" text NOT NULL,
  "price" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "shippings" (
  "id" SERIAL PRIMARY KEY,
  "size_id" int NOT NULL,
  "category_id" int NOT NULL,
  "address_id" int NOT NULL,
  "payment_id" int NOT NULL,
  "status_shipping" varchar(256) NOT NULL,
  "review" text,
  "is_play_game" boolean NOT NULL DEFAULT (false),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "sizes" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(256) NOT NULL,
  "description" text NOT NULL,
  "price" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "payments" (
  "id" SERIAL PRIMARY KEY,
  "payment_status" varchar(50) NOT NULL,
  "total_cost" int NOT NULL,
  "promo_id" int,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "promos" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(256) NOT NULL,
  "min_fee" int NOT NULL,
  "discount" int NOT NULL,
  "max_discount" int NOT NULL,
  "quota" int NOT NULL,
  "expire_date" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "add_ons" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(256) NOT NULL,
  "description" text NOT NULL,
  "price" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "add_on_shippings" (
  "id" SERIAL PRIMARY KEY,
  "shipping_id" int NOT NULL,
  "add_on_id" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "promo_users" (
  "id" SERIAL PRIMARY KEY,
  "promo_id" int NOT NULL,
  "user_id" int NOT NULL,
  "is_used" boolean NOT NULL DEFAULT (false),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

ALTER TABLE "users" ADD FOREIGN KEY ("reffered_user_id") REFERENCES "users" ("id");

ALTER TABLE "addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "shippings" ADD FOREIGN KEY ("size_id") REFERENCES "sizes" ("id");

ALTER TABLE "shippings" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "shippings" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");

ALTER TABLE "shippings" ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("promo_id") REFERENCES "promos" ("id");

ALTER TABLE "add_on_shippings" ADD FOREIGN KEY ("shipping_id") REFERENCES "shippings" ("id");

ALTER TABLE "add_on_shippings" ADD FOREIGN KEY ("add_on_id") REFERENCES "add_ons" ("id");

ALTER TABLE "promo_users" ADD FOREIGN KEY ("promo_id") REFERENCES "promos" ("id");

ALTER TABLE "promo_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");



insert into users (email,"password", long_name, phone, "role", balance, photo, refferal_code, reffered_user_id) values 
('admin@gmail.com', '$2a$04$ZFNhfT03PiRSI92uZ2A2FehiT3EonpyVx5GqDGbzQfx5qBImgzq4m', 'admin', '08912345671', 'ADMIN', 0, '', '', null),
('user1@gmail.com', '$2a$04$ZFNhfT03PiRSI92uZ2A2FehiT3EonpyVx5GqDGbzQfx5qBImgzq4m', 'user1', '08912345672', 'USER', 500000, '', 'qwerty12', null),
('user2@gmail.com', '$2a$04$ZFNhfT03PiRSI92uZ2A2FehiT3EonpyVx5GqDGbzQfx5qBImgzq4m', 'user2', '08912345673', 'USER', 500000, '', 'qwerty34', null),
('user3@gmail.com', '$2a$04$ZFNhfT03PiRSI92uZ2A2FehiT3EonpyVx5GqDGbzQfx5qBImgzq4m', 'user3', '08912345672', 'USER', 500000, '', 'qwerty56', null);

insert into addresses (recipient_name, full_address, recipient_phone, user_id) values 
('budi', 'kota Surabaya jawa timur', '0897654321', 2),
('susi', 'kota Gresik jawa timur', '0897654322', 2),
('tio', 'kota Malang jawa timur', '0897654323', 2),
('mukti', 'Madura jawa timur', '0897654325', 2),
('kuki', 'denpasar bali', '0897654326', 2),
('dana', 'surabaya jawa timur', '0897654327', 2),
('aris', 'surabaya jawa timur', '0897654328', 2),
('dudung', 'jakarta jawa barat', '0897654328', 2),
('putri', 'lamongan jawa timur', '0897654329', 2),
('siti', 'gresik jawa timur', '0897654339', 2),
('bintang', 'jakarta jawa barat', '0897654339', 2),
('sinta', 'jakarta jawa barat', '0897654339', 2),
('adi', 'surabaya jawa timur', '0897654339', 2);

insert into sizes(name, description, price) values 
('Large', 'VOLUME > 50x50x50 cm OR WEIGHT > 3 kg', 150000),
('Medium', '25x25x25 cm < VOLUME <= 50x50x50 cm OR 2 kg < WEIGHT <= 3kg', 100000),
('Small', 'VOLUME <= 25x25x25 cm OR WEIGHT <= 2 kg', 75000);

insert into add_ons (name, description, price) values
('Safe package', 'Add more protection for your package', 50000),
('Cooler', 'Keep your food or drink cool', 20000),
('Heatkeeper', 'Keep your food or drink warm', 15000);

insert into categories (name, description, price) values
('Food and Beverages', 'Send your food or beverages safely with us', 30000),
('Fragile', 'Send your fragile package safely with us', 25000);

insert into promos ("name", min_fee, discount, max_discount, quota, expire_date) values 
('Promo 1', 20000, 40, 15000, 4, '2022-10-13 13:59:53.298'),
('Promo 2', 20000, 60, 15000, 6, '2022-12-13 15:59:53.298'),
('Promo 3', 20000, 80, 15000, 8, '2022-12-13 20:59:53.298');

insert into promo_users (promo_id, user_id, is_used) values 
(1, 2, false),
(2, 2, false),
(3, 2, false);

insert into payments (payment_status, total_cost, promo_id) values
('PENDING', 245000, NULL),
('PENDING', 177000, NULL),
('PENDING', 123000, NULL),
('PENDING', 145000, NULL),
('PENDING', 195000, NULL),
('SUCCESS', 200000, NULL),
('SUCCESS', 230000, NULL),
('SUCCESS', 230000, NULL),
('SUCCESS', 200000, NULL),
('SUCCESS', 195000, NULL),
('SUCCESS', 230000, NULL);

insert into shippings (size_id, category_id, address_id, payment_id, status_shipping, review, is_play_game) values
(1, 1, 1, 1, 'PROCESS', null, false),
(1, 2, 2, 2, 'PROCESS', null, false),
(2, 1, 3, 3, 'PROCESS', null, false),
(2, 2, 4, 4, 'PROCESS', null, false),
(1, 1, 5, 5, 'PROCESS', null, false),
(1, 1, 6, 6, 'PICKUP', null, false),
(1, 1, 7, 7, 'DELIVERY', null, false),
(1, 1, 8, 8, 'DELIVERED', null, false),
(1, 1, 9, 9, 'DELIVERED', 'Barang Diterima dengan baik', false),
(1, 1, 10, 10, 'DELIVERED', null, true),
(1, 1, 11, 11, 'DELIVERED', 'Barang bagus', true);

insert into add_on_shippings (shipping_id, add_on_id) values
(1, 1),
(1, 2),
(2, 2),
(3, 2),
(4, 2),
(5, 3),
(6, 2),
(7, 1),
(8, 1),
(9, 2),
(10, 3),
(11, 1);