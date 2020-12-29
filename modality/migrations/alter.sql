ALTER TABLE modalities RENAME COLUMN start_date_time TO add_date_time;
ALTER TABLE input_texts ADD COLUMN url TEXT NOT NULL DEFAULT '';
