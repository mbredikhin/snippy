CREATE TABLE users (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name varchar(255),
  username varchar(255) NOT NULL UNIQUE,
  password_hash varchar(255) NOT NULL UNIQUE
);

CREATE TABLE languages (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name varchar(255) NOT NULL
);

CREATE TABLE lists (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  user_id int REFERENCES users (id) ON DELETE CASCADE NOT NULL,
  name varchar(255) NOT NULL
);

CREATE TABLE snippets (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  list_id int REFERENCES lists (id) ON DELETE CASCADE NOT NULL,
  name varchar(255) NOT NULL,
  language_id int NOT NULL REFERENCES languages (id),
  content text
);

CREATE TABLE favourite_snippets (
  snippet_id int REFERENCES snippets (id) NOT NULL,
  user_id int REFERENCES users (id) NOT NULL
);

CREATE TABLE tags (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name varchar(255) NOT NULL,
  user_id int REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE snippets_tags (
  snippet_id int REFERENCES snippets (id) ON DELETE CASCADE NOT NULL,
  tag_id int REFERENCES tags (id) ON DELETE CASCADE NOT NULL
);
