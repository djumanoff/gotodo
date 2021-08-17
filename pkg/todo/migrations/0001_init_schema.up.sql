CREATE TABLE IF NOT EXISTS todos (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "title" TEXT,
    "body" TEXT,
    "status" TEXT,
    "owner_id" VARCHAR(255)
);
