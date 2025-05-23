CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE roles(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE user_roles(
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE films(
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    genre VARCHAR(50) NOT NULL,
    duration INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE halls(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    capacity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions(
    id SERIAL PRIMARY KEY,
    start_time TIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- фильм + зал + сеанс
CREATE TABLE film_sessions(
    id SERIAL PRIMARY KEY,
    film_id INT REFERENCES films(id) ON DELETE CASCADE,
    hall_id INT REFERENCES halls(id) ON DELETE CASCADE,
    session_id INT REFERENCES sessions(id) ON DELETE CASCADE,
    session_date DATE NOT NULL,
    UNIQUE (film_id, hall_id, session_id, session_date),
    ticket_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE seats(
    id SERIAL PRIMARY KEY,
    hall_id INT REFERENCES halls(id) ON DELETE CASCADE,
    row_num INT NOT NULL,
    seat_num INT NOT NULL,
    UNIQUE (hall_id, row_num, seat_num),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    film_session_id INT REFERENCES film_sessions(id) ON DELETE CASCADE,
    seat_id INT REFERENCES seats(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL DEFAULT 'confirmed', -- может быть еще cancelled и pending  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (film_session_id, seat_id) 
);




