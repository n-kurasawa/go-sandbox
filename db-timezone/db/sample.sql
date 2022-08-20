CREATE TABLE IF NOT EXISTS sample_table(
    id INT(11) AUTO_INCREMENT NOT NULL, 
    datetime DATETIME NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO sample_table (datetime) VALUES ('2022-08-20 00:00:00');