-- Table: public.records

-- DROP TABLE IF EXISTS public.records;

CREATE TABLE IF NOT EXISTS public.records
(
    objectid integer NOT NULL,
    eventts integer NOT NULL,
    idfas integer,
    iddrv bytea,
    flg integer,
    mileage integer,
    vimp integer,
    timp integer,
    uboard integer,
    batlife integer,
    sumacc integer,
    phone bytea,
    amtrx integer,
    amtry integer,
    amtrz integer,
    tachocardid bytea,
    lat integer,
    lon integer,
    gpsvel integer,
    gpsdir integer,
    gpsnsat integer,
    gpsalt integer,
    unival0 integer,
    unival1 integer,
    unival2 integer,
    unival3 integer,
    unival4 integer,
    unival5 integer,
    spn70 integer,
    spn91 integer,
    spn100 integer,
    spn110 integer,
    spn174 integer,
    spn175 integer,
    spn182 integer,
    spn184 integer,
    spn190 integer,
    spn244 integer,
    spn245 integer,
    spn247 integer,
    spn250 integer,
    spn521 integer,
    spn522 integer,
    spn527 integer,
    spn582 bytea,
    spn597 integer,
    spn598 integer,
    spn914 integer,
    spn916 integer,
    spn928 bytea,
    spn1624 integer,
    spn1821 integer,
    spn1856 integer,
    tlls1 integer,
    clls1 integer,
    flls1 integer,
    tlls2 integer,
    clls2 integer,
    flls2 integer,
    tlls3 integer,
    clls3 integer,
    flls3 integer,
    tlls4 integer,
    clls4 integer,
    flls4 integer,
    tlls5 integer,
    clls5 integer,
    flls5 integer,
    tlls6 integer,
    clls6 integer,
    flls6 integer,
    tlls7 integer,
    clls7 integer,
    flls7 integer,
    tlls8 integer,
    clls8 integer,
    flls8 integer,
    spnid bytea,
    spnval bytea,
    tpmsp smallint,
    tpmsv smallint,
    recnum integer,
    tpms bytea
) PARTITION BY RANGE (eventts)
WITH (
    OIDS = FALSE
);




-- Partitions SQL

CREATE TABLE public.records_2020 PARTITION OF public.records
( 
    CONSTRAINT records_2020_pk PRIMARY KEY (objectid, eventts)
)
    FOR VALUES FROM (347072400) TO (378694799);
	
CREATE TABLE public.records_2021 PARTITION OF public.records
( 
    CONSTRAINT records_2021_pk PRIMARY KEY (objectid, eventts)
)
    FOR VALUES FROM (378694800) TO (410230799);


CREATE TABLE public.records_2022 PARTITION OF public.records
( 
    CONSTRAINT records_2022_pk PRIMARY KEY (objectid, eventts)
)
    FOR VALUES FROM (410230800) TO (441766799);


CREATE TABLE public.records_2023 PARTITION OF public.records
( 
    CONSTRAINT records_2023_pk PRIMARY KEY (objectid, eventts)
)
    FOR VALUES FROM (441766800) TO (473302799);


CREATE TABLE public.records_2024 PARTITION OF public.records
( 
    CONSTRAINT records_pk PRIMARY KEY (objectid, eventts)
)
    FOR VALUES FROM (473302800) TO (504925199);

