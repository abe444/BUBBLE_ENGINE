CREATE TABLE access (
    pw VARCHAR(90), 
    creation timestamp
)

CREATE TABLE entries (
    id BIGSERIAL PRIMARY KEY,
    user VARCHAR(150),
    filepath TEXT, 
    filename TEXT, 
    creation timestamp
)

