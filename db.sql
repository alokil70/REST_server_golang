CREATE TABLE public.author (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    author_id UUID NOT NULL,
    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author (name) VALUES ('народ');
INSERT INTO author (name) VALUES ('Пушкин');
INSERT INTO author (name) VALUES ('Булгаков');



INSERT INTO book (name, author_id) VALUES ('Колобок', '752532e6-4cad-44ae-8e5e-c8492c557674');
INSERT INTO book (name, author_id) VALUES ('Руслан и Людмила', 'f8a6a73d-b720-4249-a59f-590d254cc185');
INSERT INTO book (name, author_id) VALUES ('Собачье сердце', '7013353a-a0f6-46be-8f2a-7fae8b3268fe');




CREATE TABLE public.book (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book_authors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID NOT NULL,
    author_id UUID NOT NULL,

    CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id)
    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
    CONSTRAINT book_author_unique UNIQUE (book_id, author_id)
)


INSERT INTO book (id, name) VALUES ('752532e6-4cad-44ae-8e5e-c8492c557674', 'Колобок');
INSERT INTO book (id, name) VALUES ('f8a6a73d-b720-4249-a59f-590d254cc185', 'Руслан и Людмила');
INSERT INTO book (id, name) VALUES ('7013353a-a0f6-46be-8f2a-7fae8b3268fe', 'Собачье сердце');


INSERT INTO book_authors (book_id, author_id) VALUES ('7013353a-a0f6-46be-8f2a-7fae8b3268fe', 'Собачье сердце');
INSERT INTO book_authors (book_id, author_id) VALUES ('7013353a-a0f6-46be-8f2a-7fae8b3268fe', 'Собачье сердце');


SELECT   b.id, b.name,
         array((SELECT ba.author_id FROM book_authors ba WHERE ba.book_id = b.id)) AS authors
FROM     book b
GROUP BY b.id, b.name;

