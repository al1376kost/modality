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
	id_old INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	comment TEXT, 
	id INTEGER, 
	active INTEGER
);

INSERT INTO simple_types (name,comment,id,active) VALUES
	 ('вопрос','%вопрос%',1,1),
	 ('риторический вопрос','%риторический вопрос%',2,1),
	 ('повелительное наклонение','%повелител%наклон%',3,1),
	 ('согласие','%соглас%',4,1),
	 ('несогласие','%несоглас%',5,1),
	 ('деонтическая модальность','%диантич%модаль%',6,1),
	 ('аксиологическая модальность','%аксиологич%мод%',7,1),
	 ('эпистемическая модальность','%эпистемич%',8,1),
	 ('положительный речевой акт','%положитель%р%а%',9,1),
	 ('отрицательный речевой акт','%отрицатель%р%а%',10,1);
INSERT INTO simple_types (name,comment,id,active) VALUES
	 ('метаязыковой речевой акт','%метаязыков%р%а%',11,1),
	 ('метаязыковое вводное слово','%метаязыков%в%с%',12,1),
	 ('алетическая модальность','%алетич%модаль%',13,1),
	 ('модальность желания','%модаль%желан%',14,1),
	 ('модальность акцентирования','%акцентир%',15,1),
	 ('абивалентный речевой акт','%абивален%р%а%',16,1),
	 ('модальность акцентирования','%усил%',15,0),
	 ('модальность акцентирования','%подч_р%',15,0),
	 ('аксиологическая модальность','%оценоч%',7,0),
	 ('аксиологическая модальность','%эмоцион%',7,0);
INSERT INTO simple_types (name,comment,id,active) VALUES
	 ('сослагательное наклонение','%сослагат%',17,1),
	 ('деонтическая модальность','%возможност%',6,0),
	 ('речевой акт побуждения','речевой акт побуждения',18,1),
	 ('деонтическая модальность','%рекоменд%',6,0),
	 ('деонтическая модальность','%допущен%',6,0),
	 ('модальность акцентирования','%выделен%',15,0),
	 ('метаязыковое вводное слово','%ист%инф%',12,0),
	 ('деонтическая модальность','%должен%',6,0),
	 ('деонтическая модальность','%необходим%',6,0);