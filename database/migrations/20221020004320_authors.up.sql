BEGIN;

CREATE TABLE authors (
    name TEXT NOT NULL,
    book_name TEXT REFERENCES books(name)
);

COMMIT;
