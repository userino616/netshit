CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS movies (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS user_movie_bookmarks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users (id),
    movie_id UUID NOT NULL REFERENCES movies (id),
    UNIQUE (user_id, movie_id)
);

CREATE TABLE IF NOT EXISTS user_movie_watched (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users (id),
    movie_id UUID NOT NULL REFERENCES movies (id),
    UNIQUE (user_id, movie_id)
);