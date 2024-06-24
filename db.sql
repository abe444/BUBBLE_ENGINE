CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR, 
    password VARCHAR,
    creation timestamp
)

CREATE TABLE entries (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200)
    author VARCHAR(150)
    article INT, 
    creation timestamp
)
