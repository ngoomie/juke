BEGIN TRANSACTION;
DROP TABLE IF EXISTS "library_items";
CREATE TABLE "library_items" (
	"item_file"	TEXT NOT NULL,
	"artist"	TEXT,
	"title"	TEXT,
	"album"	TEXT,
	PRIMARY KEY("item_file")
);
COMMIT;
