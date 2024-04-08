INSERT INTO users (name, username, password_hash) VALUES ('John Doe', 'username', '6a316a6e3269307366307133316e636c335baa61e4c9b93f3f0682250b6cf8331b7ee68fd8');
INSERT INTO languages (name) VALUES ('Language 1');
INSERT INTO lists (user_id, name) VALUES (1, 'List 1');
INSERT INTO tags (name, user_id) VALUES ('Tag 1', 1), ('Tag 2', 1);
INSERT INTO snippets (list_id, name, language_id, content) VALUES (1, 'Snippet 1', 1, '1st snippet content'), (1, 'Snippet 2', 1, '2nd snippet content');
INSERT INTO favourite_snippets (snippet_id, user_id) VALUES (1, 1);
INSERT INTO snippets_tags (snippet_id, tag_id) VALUES (1, 1), (1, 2), (2, 1);
