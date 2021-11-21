CREATE TABLE IF NOT EXISTS todos (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "title" TEXT,
    "body" TEXT,
    "status" TEXT,
    "owner_id" integer
);

CREATE TABLE IF NOT EXISTS users (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "first_name" TEXT,
    "last_name" TEXT,
    "username" TEXT,
    "password" TEXT
);
