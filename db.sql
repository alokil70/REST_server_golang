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
