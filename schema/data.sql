INSERT INTO manager(username, password, first_name, last_name)
VALUES ('supernova', 'secret', 'Super', 'Nova');

INSERT INTO store (title, manager_id)
VALUES ('Buger King', 1);

INSERT INTO category(title, image, store_id)
VALUES ('Завтраки', 'empty', 1),
       ('Комбо', 'empty', 1),
       ('Бакеты', 'empty', 1),
       ('Баскеты', 'empty', 1),
       ('Снэки', 'empty', 1),
       ('Бургеры', 'empty', 1),
       ('Напитки', 'empty', 1),
       ('Десерты', 'empty', 1),
       ('Детское меню', 'empty', 1);

INSERT INTO product(title, description, price, image, category_id)
VALUES ('Воппер',
        'Ароматная булочка с черным и белым кунжутом, салат Айсберг, плавленный сыр Гауда, ветчина из индейки и котлета из мраморной говядины. Бургер заправляем сладкой горчицей',
        1690, 'empty', 6),
       ('Биг Кинг',
        'Ароматная булочка с черным и белым кунжутом, салат Айсберг, плавленный сыр Гауда, ветчина из индейки и котлета из мраморной говядины. Бургер заправляем сладкой горчицей',
        1200, 'empty', 6),
       ('Кола', '', 1200, 'empty', 7);

INSERT INTO customer(external_id, username, phone, first_name, last_name)
VALUES (123456789, 'supernova', '+77777777777', 'Super', 'Nova');

INSERT INTO orders(customer_id, store_id, address)
values (1, 1, 'улица Кабанбай батыра, 110');

INSERT INTO order_item(order_id, product_id, amount)
VALUES (1, 1, 2),
       (1, 3, 1);