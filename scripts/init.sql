CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE game (
    id INT AUTO_INCREMENT PRIMARY KEY,
    date DATE,
    time TIME
);

CREATE TABLE seat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    col INT NOT NULL,
    `row` INT NOT NULL,
    sec INT NOT NULL
);

CREATE TABLE ticket (
    id INT AUTO_INCREMENT PRIMARY KEY,
    price DECIMAL(10, 2) NOT NULL,
    user_id INT NOT NULL,
    game_id INT NOT NULL,
    seat_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (seat_id) REFERENCES seat(id)
);