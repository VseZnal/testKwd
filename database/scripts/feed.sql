INSERT INTO books (
    name
)
VALUES
('testBook1'),
('testBook2'),
('testBook3'),
('testBook4'),
('testBook5'),
('testBook6');


INSERT INTO authors (
    name,
    book_name
)
VALUES
('testAuthor1', 'testBook1'),
('testAuthor1', 'testBook2'),
('testAuthor2', 'testBook3'),
('testAuthor3', 'testBook3'),
('testAuthor4', 'testBook4'),
('testAuthor4', 'testBook5'),
('testAuthor4', 'testBook6');
