-- +goose Up
-- +goose StatementBegin
INSERT INTO vinyls (title, artist, releasedate, price, rating) VALUES
('Abbey Road', 'The Beatles', 1969, 20, 4.8),
('The Dark Side of the Moon', 'Pink Floyd', 1973, 25, 4.9),
('Thriller', 'Michael Jackson', 1982, 28, 4.7),
('Led Zeppelin IV', 'Led Zeppelin', 1971, 22, 4.6),
('Back in Black', 'AC/DC', 1980, 23, 4.5),
('Rumours', 'Fleetwood Mac', 1977, 24, 4.8),
('The Wall', 'Pink Floyd', 1979, 25, 4.7),
('Hotel California', 'Eagles', 1976, 26, 4.6),
('Sgt. Pepper''s Lonely Hearts Club Band', 'The Beatles', 1967, 21, 4.9),
('The Joshua Tree', 'U2', 1987, 23, 4.7),
('Born to Run', 'Bruce Springsteen', 1975, 26, 4.8),
('Blood on the Tracks', 'Bob Dylan', 1975, 25, 4.6),
('The Velvet Underground & Nico', 'The Velvet Underground', 1967, 27, 4.5),
('Who''s Next', 'The Who', 1971, 22, 4.7),
('The Rise and Fall of Ziggy Stardust and the Spiders from Mars', 'David Bowie', 1972, 24, 4.8),
('The Doors', 'The Doors', 1967, 30, 4.6),
('Blue', 'Joni Mitchell', 1971, 23, 4.9),
('Horses', 'Patti Smith', 1975, 25, 4.5),
('Court and Spark', 'Joni Mitchell', 1974, 29, 4.7),
('Ramones', 'Ramones', 1976, 24, 4.6),
('Pet Sounds', 'The Beach Boys', 1966, 30, 4.9),
('Kind of Blue', 'Miles Davis', 1959, 35, 4.8),
('The Queen Is Dead', 'The Smiths', 1986, 23, 4.7),
('Nevermind', 'Nirvana', 1991, 28, 4.8),
('A Love Supreme', 'John Coltrane', 1965, 32, 4.9),
('Born to Run', 'Bruce Springsteen', 1975, 26, 4.8),
('London Calling', 'The Clash', 1979, 29, 4.7),
('Electric Ladyland', 'Jimi Hendrix', 1968, 31, 4.9),
('The Clash', 'The Clash', 1977, 24, 4.6),
('The Velvet Underground & Nico', 'The Velvet Underground', 1967, 27, 4.5),
('Highway 61 Revisited', 'Bob Dylan', 1965, 33, 4.8),
('Darkness on the Edge of Town', 'Bruce Springsteen', 1978, 28, 4.7),
('The Doors', 'The Doors', 1967, 30, 4.6),
('Back to Black', 'Amy Winehouse', 2006, 27, 4.8),
('Exile on Main St.', 'The Rolling Stones', 1972, 32, 4.7),
('Court and Spark', 'Joni Mitchell', 1974, 29, 4.7),
('Disintegration', 'The Cure', 1989, 30, 4.9),
('Is This It', 'The Strokes', 2001, 26, 4.7),
('Purple Rain', 'Prince', 1984, 31, 4.8),
('The Bends', 'Radiohead', 1995, 25, 4.6);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM vinyls;
-- +goose StatementEnd
