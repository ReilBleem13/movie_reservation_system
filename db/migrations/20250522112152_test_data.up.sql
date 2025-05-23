-- roles
INSERT INTO roles (name) VALUES 
('admin'), 
('user');

-- users
INSERT INTO users (name, email, password_hash) VALUES
('Alice Admin', 'alice@example.com', 'hashed_pass_1'),
('Bob User', 'bob@example.com', 'hashed_pass_2');

-- user_roles
INSERT INTO user_roles (user_id, role_id) VALUES
(1, 1),  -- Alice → admin
(2, 2); -- Bob → user

INSERT INTO films (title, description, genre, duration) VALUES
('The Matrix', 'A hacker discovers the truth about reality.', 'Sci-Fi', 136),
('Inception', 'A thief enters dreams to steal secrets.', 'Action', 148);

INSERT INTO halls (name, capacity) VALUES
('Main Hall', 100),
('VIP Hall', 50);

INSERT INTO sessions (start_time) VALUES
('10:00:00'),
('14:30:00'),
('19:00:00');

INSERT INTO film_sessions (film_id, hall_id, session_id, session_date, ticket_price) VALUES
(1, 1, 1, '2025-05-15', 9.99),  -- Matrix, Main Hall, 10:00
(2, 2, 3, '2025-05-15', 14.99); -- Inception, VIP Hall, 19:00

-- Seats for Main Hall (id = 1)
INSERT INTO seats (hall_id, row_num, seat_num) VALUES
(1, 1, 1), (1, 1, 2), (1, 1, 3),
(1, 2, 1), (1, 2, 2), (1, 2, 3),
(1, 3, 1), (1, 3, 2), (1, 3, 3);

-- Seats for VIP Hall (id = 2)
INSERT INTO seats (hall_id, row_num, seat_num) VALUES
(2, 1, 1), (2, 1, 2), (2, 1, 3);

-- Bob бронирует место в Matrix (Main Hall)
INSERT INTO reservations (user_id, film_session_id, seat_id, status) VALUES
(2, 1, 1, 'confirmed'),
(2, 1, 2, 'confirmed');

-- Alice бронирует место в Inception (VIP Hall)
INSERT INTO reservations (user_id, film_session_id, seat_id, status) VALUES
(1, 2, 10, 'pending');  -- seat_id 10 = hall 2, row 1, seat 1
