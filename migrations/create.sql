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

INSERT INTO languages (name) VALUES
	 ('RU'),
	 ('EN');


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

INSERT INTO simple_types (name,comment) VALUES
	 ('вопрос',NULL),
	 ('риторический вопрос',NULL),
	 ('повелительное наклонение',NULL),
	 ('согласие',NULL),
	 ('несогласие',NULL),
	 ('деонтическая модальность',NULL),
	 ('аксиологическая модальность',NULL),
	 ('эпистимическая модальность',NULL),
	 ('положительный речевой акт',NULL),
	 ('отрицательный речевой акт',NULL);
INSERT INTO simple_types (name,comment) VALUES
	 ('метаязыковой речевой акт',NULL),
	 ('метаязыковое вводное слово',NULL),
	 ('алетическая модальность',NULL),
	 ('модальность желания',NULL),
	 ('модальность акцентирования',NULL),
	 ('амбивалентный речевой акт',NULL),
	 ('сослагательное наклонение',NULL);
