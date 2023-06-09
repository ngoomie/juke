BEGIN TRANSACTION;
DROP TABLE IF EXISTS "directories";
CREATE TABLE "directories" (
	"id"	INTEGER NOT NULL UNIQUE,
	"path"	TEXT NOT NULL UNIQUE,
	"type"	INTEGER NOT NULL,
	"depth"	INTEGER,
	PRIMARY KEY("id" AUTOINCREMENT)
);
DROP TABLE IF EXISTS "library_items";
CREATE TABLE "library_items" (
	"item_file"	TEXT NOT NULL,
	"directory"	TEXT NOT NULL,
	PRIMARY KEY("item_file","directory")
);
DROP TABLE IF EXISTS "meta";
CREATE TABLE "meta" (
	"name"	TEXT,
	"type"	NUMERIC
);
COMMIT;
