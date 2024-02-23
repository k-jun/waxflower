CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE games (
    id CHAR(36) PRIMARY KEY,
    date DATE,
    time TIME
);

CREATE TABLE seats (
    id CHAR(36) PRIMARY KEY,
    sec INT NOT NULL,
    `row` CHAR(2) NOT NULL,
    col INT NOT NULL
);

CREATE TABLE tickets (
    id CHAR(36) PRIMARY KEY,
    price INT NOT NULL,
    user_id CHAR(36) NULL,
    game_id CHAR(36) NOT NULL,
    seat_id CHAR(36) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (game_id) REFERENCES games(id),
    FOREIGN KEY (seat_id) REFERENCES seats(id)
);