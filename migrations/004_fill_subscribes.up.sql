-- Подписки для Ивана Петрова
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('StreamFlix', 299, (SELECT id FROM users WHERE name = 'Иван Петров'), '2024-01-15 00:00:00', '2024-07-15 00:00:00'),
('MusicWave', 149, (SELECT id FROM users WHERE name = 'Иван Петров'), '2024-03-01 00:00:00', '2024-09-01 00:00:00'),
('CloudPlus', 199, (SELECT id FROM users WHERE name = 'Иван Петров'), '2024-02-10 00:00:00', NULL),
('GameHub Pro', 399, (SELECT id FROM users WHERE name = 'Иван Петров'), '2024-04-05 00:00:00', '2024-10-05 00:00:00');

-- Подписки для Марии Сидоровой
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('FitnessOnline', 249, (SELECT id FROM users WHERE name = 'Мария Сидорова'), '2024-02-20 00:00:00', '2024-08-20 00:00:00'),
('Readly Unlimited', 179, (SELECT id FROM users WHERE name = 'Мария Сидорова'), '2024-01-10 00:00:00', NULL),
('FoodBox Deluxe', 899, (SELECT id FROM users WHERE name = 'Мария Сидорова'), '2024-03-15 00:00:00', '2024-06-15 00:00:00');

-- Подписки для Алексея Козлова
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('StreamFlix', 299, (SELECT id FROM users WHERE name = 'Алексей Козлов'), '2024-01-05 00:00:00', NULL),
('WorkSpace Pro', 499, (SELECT id FROM users WHERE name = 'Алексей Козлов'), '2024-02-28 00:00:00', '2024-08-28 00:00:00'),
('MusicWave', 149, (SELECT id FROM users WHERE name = 'Алексей Козлов'), '2024-03-10 00:00:00', '2024-09-10 00:00:00'),
('DevTools Cloud', 699, (SELECT id FROM users WHERE name = 'Алексей Козлов'), '2024-01-20 00:00:00', NULL);

-- Подписки для Екатерины Новиковой
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('BeautyBox', 549, (SELECT id FROM users WHERE name = 'Екатерина Новикова'), '2024-03-01 00:00:00', '2024-09-01 00:00:00'),
('YogaMind', 199, (SELECT id FROM users WHERE name = 'Екатерина Новикова'), '2024-02-14 00:00:00', NULL),
('CraftMaster', 329, (SELECT id FROM users WHERE name = 'Екатерина Новикова'), '2024-04-01 00:00:00', '2024-07-01 00:00:00');

-- Подписки для Дмитрия Волкова
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('GameHub Pro', 399, (SELECT id FROM users WHERE name = 'Дмитрий Волков'), '2024-01-12 00:00:00', NULL),
('SportTV Online', 449, (SELECT id FROM users WHERE name = 'Дмитрий Волков'), '2024-02-25 00:00:00', '2024-08-25 00:00:00'),
('StreamFlix', 299, (SELECT id FROM users WHERE name = 'Дмитрий Волков'), '2024-03-18 00:00:00', '2024-09-18 00:00:00'),
('CarCare Plus', 199, (SELECT id FROM users WHERE name = 'Дмитрий Волков'), '2024-04-05 00:00:00', '2024-10-05 00:00:00');

-- Подписки для Анны Морозовой
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('MusicWave', 149, (SELECT id FROM users WHERE name = 'Анна Морозова'), '2024-01-08 00:00:00', NULL),
('PhotoMagic Pro', 279, (SELECT id FROM users WHERE name = 'Анна Морозова'), '2024-02-20 00:00:00', '2024-08-20 00:00:00'),
('Readly Unlimited', 179, (SELECT id FROM users WHERE name = 'Анна Морозова'), '2024-03-05 00:00:00', '2024-09-05 00:00:00');

-- Подписки для Сергея Орлова
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('InvestExpert', 599, (SELECT id FROM users WHERE name = 'Сергей Орлов'), '2024-01-25 00:00:00', NULL),
('BusinessInsider Pro', 799, (SELECT id FROM users WHERE name = 'Сергей Орлов'), '2024-02-10 00:00:00', '2024-08-10 00:00:00'),
('StreamFlix', 299, (SELECT id FROM users WHERE name = 'Сергей Орлов'), '2024-03-22 00:00:00', '2024-09-22 00:00:00'),
('CloudPlus', 199, (SELECT id FROM users WHERE name = 'Сергей Орлов'), '2024-04-01 00:00:00', '2024-07-01 00:00:00');

-- Подписки для Ольги Лебедевой
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('FitnessOnline', 249, (SELECT id FROM users WHERE name = 'Ольга Лебедева'), '2024-01-30 00:00:00', NULL),
('CookLikeChef', 379, (SELECT id FROM users WHERE name = 'Ольга Лебедева'), '2024-02-15 00:00:00', '2024-08-15 00:00:00'),
('MusicWave', 149, (SELECT id FROM users WHERE name = 'Ольга Лебедева'), '2024-03-08 00:00:00', '2024-09-08 00:00:00');

-- Подписки для Павла Соколова
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('GameHub Pro', 399, (SELECT id FROM users WHERE name = 'Павел Соколов'), '2024-01-18 00:00:00', '2024-07-18 00:00:00'),
('CodeMasters Pro', 899, (SELECT id FROM users WHERE name = 'Павел Соколов'), '2024-02-22 00:00:00', NULL),
('StreamFlix', 299, (SELECT id FROM users WHERE name = 'Павел Соколов'), '2024-03-12 00:00:00', '2024-09-12 00:00:00'),
('MusicWave', 149, (SELECT id FROM users WHERE name = 'Павел Соколов'), '2024-04-02 00:00:00', '2024-10-02 00:00:00');

-- Подписки для Натальи Романовой
INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES
('LuxuryFashion Club', 1299, (SELECT id FROM users WHERE name = 'Наталья Романова'), '2024-01-05 00:00:00', NULL),
('ArtGallery Online', 459, (SELECT id FROM users WHERE name = 'Наталья Романова'), '2024-02-28 00:00:00', '2024-08-28 00:00:00'),
('MusicWave', 149, (SELECT id FROM users WHERE name = 'Наталья Романова'), '2024-03-15 00:00:00', '2024-09-15 00:00:00'),
('BeautyBox', 549, (SELECT id FROM users WHERE name = 'Наталья Романова'), '2024-04-10 00:00:00', '2024-07-10 00:00:00');