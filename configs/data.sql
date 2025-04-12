-- Create tables

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT
);

CREATE TABLE IF NOT EXISTS user_animals (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    animal_type TEXT NOT NULL,
    count INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS animal_duplicate_requests (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    event_code TEXT NOT NULL,
    institution_or_name TEXT NOT NULL,
    fiscal_or_personal_code TEXT NOT NULL,
    farm_name TEXT NOT NULL,
    address TEXT NOT NULL,
    locality TEXT NOT NULL,
    representative_name TEXT NOT NULL,
    phone TEXT,
    email TEXT,
    full_name TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS animal_duplicates (
   id SERIAL PRIMARY KEY,
   request_id INTEGER NOT NULL REFERENCES animal_duplicate_requests(id),
   tag_number TEXT NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert mock data

INSERT INTO users (name, email) VALUES
    ('Ion Popescu', 'ion.popescu@example.com');

INSERT INTO user_animals (user_id, animal_type, count) VALUES
    (1, 'bovine', 3);
