CREATE TABLE IF NOT EXISTS book(
    id VARCHAR PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR NOT NULL,
    pages INT NOT NULL,
    writer VARCHAR NOT NULL,
    genre VARCHAR NOT NULL,
    description TEXT NOT NULL
);

INSERT INTO book (title, pages, writer, genre, description) 
VALUES ('Bumi', 500, 'Tere Liye', 'Action, Intrik, Fiksi', 'Buku ke-1'),
('Bulan', 500, 'Tere Liye', 'Action, Intrik, Fiksi', 'Buku ke-2'),
('Matahari', 500, 'Tere Liye', 'Action, Intrik, Fiksi', 'Buku ke-3'),
('Bintang', 500, 'Tere Liye', 'Action, Intrik, Fiksi', 'Buku ke-4');