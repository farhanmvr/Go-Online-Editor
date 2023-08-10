CREATE TABLE code_snippets (
   id varchar(48) PRIMARY KEY,
   name varchar(64) NOT NULL,
   code TEXT NOT NULL,
   status varchar(16) NOT NULL,
   date_created datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);
