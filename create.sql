CREATE TABLE IF NOT EXISTS dict (
    dict_id serial NOT NULL PRIMARY KEY,
    word VARCHAR (50)  NOT NULL,
    word_translate VARCHAR (100) NOT NULL
    );

CREATE TABLE IF NOT EXISTS users (
     users_id serial NOT NULL PRIMARY KEY,
     user_number VARCHAR (50)  NOT NULL
    );

CREATE TABLE IF NOT EXISTS dict_user (
    dict_id int REFERENCES dict (dict_id) ON UPDATE CASCADE ON DELETE CASCADE,
    users_id int REFERENCES users (users_id) ON UPDATE CASCADE ON DELETE CASCADE,
    done BOOLEAN DEFAULT FALSE NOT NULL
    );