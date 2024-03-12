INSERT INTO orders VALUES (10, 'order10'),
                          (11, 'order11'),
                          (14, 'order14'),
                          (15, 'order15');
INSERT INTO order_details (order_id, order_info)
VALUES (10, '{
    "product1": {"id": 1, "count": 2},
    "product2": {"id": 3, "count": 1},
    "product3": {"id": 6, "count": 1}
}'::jsonb),
    (11, '{
    "product1": {"id": 2, "count": 3}
}'::jsonb),
    (14, '{
    "product1": {"id": 1, "count": 3},
    "product2": {"id": 4, "count": 4}
}'::jsonb),
    (15, '{
    "product1": {"id": 5, "count": 1}
}'::jsonb);

INSERT INTO shelving VALUES (10, 'А'),
                            (20, 'Б'),
                            (30, 'В'),
                            (40, 'З'),
                            (50, 'Ж');

INSERT INTO products VALUES (1, 10, true, 6, 'Ноутбук'),
                            (2, 10, true, 4, 'Телевизор'),
                            (3, 20, true, 2, 'Телефон'),
                            (3, 40, false, 1, 'Телефон'),
                            (3, 30, false, 1, 'Телефон'),
                            (4, 50, true, 6, 'Системный блок'),
                            (5, 50, true, 1, 'Часы'),
                            (5, 10, false, 1, 'Часы'),
                            (6, 50, true, 1, 'Микрофон');
