CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY, 
    username VARCHAR(255), 
    phone VARCHAR(255), 
    email VARCHAR(255), 
    balance NUMERIC 
);