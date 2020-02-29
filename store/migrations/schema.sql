CREATE TABLE IF NOT EXISTS shopping (
	id	INTEGER NOT NULL PRIMARY KEY,
	date	text,
	sum	INTEGER,
	shop_id	INTEGER,
	complete	INTEGER,
	time	TEXT,
	owner_id	INTEGER
);
CREATE TABLE IF NOT EXISTS shop_list (
	id	INTEGER NOT NULL PRIMARY KEY,
	product_name	TEXT,
	quantity	INTEGER,
	category_id	INTEGER,
	complete	INTEGER,
	list_id	INTEGER
);
CREATE TABLE IF NOT EXISTS shop (
	id	INTEGER NOT NULL PRIMARY KEY,
	name	TEXT
);
CREATE TABLE IF NOT EXISTS owner (
	id	INTEGER NOT NULL PRIMARY KEY,
	telegram_id INTEGER,
	first_name	TEXT,
	last_name TEXT,
	user_name TEXT,
	token TEXT
);