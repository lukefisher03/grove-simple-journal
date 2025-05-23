CREATE TABLE IF NOT EXISTS notes(
    id                  serial PRIMARY KEY,
    created             timestamp NOT NULL DEFAULT NOW(),
    last_updated        timestamp NOT NULL DEFAULT NOW(),
    content             text,
    title               text NOT NULL
);

CREATE TABLE IF NOT EXISTS notes_tags(
    id          serial PRIMARY KEY,
    note_id     int NOT NULL,
    tag         text NOT NULL,
    FOREIGN KEY (note_id) REFERENCES notes(id)
);