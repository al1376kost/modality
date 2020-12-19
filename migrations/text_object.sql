-- input_texts definition

CREATE TABLE input_texts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	object_text TEXT,
	lang_id INTEGER DEFAULT 1,
	add_date_time TEXT,
	active INTEGER DEFAULT 1,
	url TEXT NOT NULL DEFAULT ''
);

-- languages definition

CREATE TABLE languages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT
);

-- modalities definition

CREATE TABLE modalities (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	modality_text TEXT,
	type_id INTEGER,
	text_id INTEGER,
	start_symbol INTEGER,
	add_date_time TEXT,
	active INTEGER DEFAULT 1
);

-- simple_types definition

CREATE TABLE "simple_types" (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	comment TEXT
);

ALTER TABLE modalities RENAME COLUMN start_date_time TO add_date_time;
ALTER TABLE input_texts ADD COLUMN url TEXT NOT NULL DEFAULT '';