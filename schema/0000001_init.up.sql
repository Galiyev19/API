CREATE TABLE users {
    id int NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL,  
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at time NOT NULL,
}

CREATE TABLE admins {
    id int NOT NULL,
    username VARCHAR(255) NOT NULL,  
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at time NOT NULL,    
}

