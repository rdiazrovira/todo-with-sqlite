-- 20240618165023 - create_todos migration
CREATE TABLE todos (
	id uuid PRIMARY KEY,
	title VARCHAR(155) NOT NULL,
	description TEXT NOT NULL
);

INSERT INTO todos (id, title, description)
VALUES ('8e85a4fa-4381-4e63-98f5-eb5a2ebfa0af', 'Todo #1', 'Random description');
