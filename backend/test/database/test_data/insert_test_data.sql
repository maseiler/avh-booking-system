INSERT INTO category(name, enabled, icon, type)
VALUES ('Sailor',
        true,
        'sailboat',
        'account'),
       ('Beverage',
        true,
        'glass-water',
        'beverage'),
       ('Active',
        true,
        'person-running',
        'account'),
       ('Retiree',
        true,
        'person-cane',
        'account');

INSERT INTO account(first_name, nickname, last_name, email, phone, balance, max_debt, category, enabled)
VALUES ('Vasco',
        'Cape Conqueror',
        'da Gama',
        'indianspice@capeofgoodhope.com',
        '+351 914 97 1498',
        3355,
        100,
        1,
        true),
       ('陽',
        'Pirate Queen',
        '石',
        'z.yisao@redfleet.cn',
        '+86 138 2501 1807',
        1203,
        100,
        1,
        true),
       ('Brandon',
        'Broseidon',
        'Poseidon', 'kingofocean@mail.com',
        '+30 123 456 78',
        3355,
        50,
        3,
        true),
       ('Jeanne',
        'Mistress of Time',
        'Calment', 'jeanne.calment1875@longevity.fr',
        '+33 4 90 18 75 97',
        55,
        122,
        4,
        true);


INSERT INTO account_option(account, key, value)
VALUES (1,
        'deceased',
        'true'),
       (2,
        'deceased',
        'true'),
       (3,
        'deceased',
        'false'),
       (4,
        'deceased',
        'true');

INSERT INTO location(name)
VALUES ('Bermuda Triangle'),
       ('Atlantis');

INSERT INTO unit(name)
VALUES ('ml');

INSERT INTO vat(rate)
VALUES (19);

INSERT INTO product_group(name, parent)
VALUES ('Alcohol', NULL),
       ('Non-Alcoholic', NULL),
       ('Port Wine', 1),
       ('Liquor', 1),
       ('Water', 2);

INSERT INTO product (name, price, vat, product_group, size, unit, category)
VALUES ('Rota das Especiarias', 1800, 1, 1, 150, 1, 2),
       ('白酒 (Báijiǔ)', 6, 1, 2, 100, 1, 2),
       ('Fiji Water', 233, 1, 3, 500, 1, 2);

INSERT INTO product_visibility(category, location, product)
VALUES (1, 1, 1),
       (1, 1, 2),
       (1, 1, 3),
       (3, 1, 1),
       (3, 1, 3),
       (3, 2, 3),
       (4, 2, 3);
