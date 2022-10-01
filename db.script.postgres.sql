CREATE TABLE customer
(
    row integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    id character varying(50) NOT NULL,
    name character varying(50) NOT NULL,
    age integer NOT NULL,
    birthday date NOT NULL,
    remark character varying(100),
    CONSTRAINT pk__customer PRIMARY KEY (row),
    CONSTRAINT unique__customer__id UNIQUE (id)
);

CREATE TABLE person
(
    row integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    id character(10) NOT NULL,
    name character varying(50) NOT NULL,
    age integer NOT NULL,
    birthday date NOT NULL,
    remark character varying(100),
    CONSTRAINT pk__person PRIMARY KEY (row),
    CONSTRAINT unique__person__id UNIQUE (id)
);

CREATE TABLE address
(
    row integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    id character(10) NOT NULL,
    text character varying(100) NOT NULL,
    CONSTRAINT pk__address PRIMARY KEY (row),
    CONSTRAINT fk__address_id FOREIGN KEY (id)
        REFERENCES person (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
	
INSERT INTO person
(id,name,age,birthday,remark)
VALUES
('A123456789','AAA',18,CURRENT_DATE,''),
('B123456789','BBB',28,CURRENT_DATE,'')

INSERT INTO address
(id,text)
VALUES
('A123456789','A1'),
('A123456789','A2'),
('B123456789','B1'),
('B123456789','B2')
