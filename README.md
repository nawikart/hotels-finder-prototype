# 1. TABLE hotels
1. prepare main CSV (download)
2. remove first row
3. create table hotels

``` bash
CREATE TABLE hotels(
    hotel_id INTEGER,
    chain_id INTEGER,
    chain_name character varying(100),
    brand_id INTEGER,
    brand_name character varying(100),
    hotel_name character varying(255),
    hotel_formerly_name character varying(255),
    hotel_translated_name character varying(255),
    addressline1 TEXT,
    addressline2 TEXT,
    zipcode character varying(25),
    city character varying(50),
    state character varying(50),
    country character varying(50),
    countryisocode character varying(2),
    star_rating character varying(22),
    longitude character varying(25),
    latitude character varying(25),
    url character varying(255),
    checkin character varying(20),
    checkout character varying(20),
    numberrooms INTEGER,
    numberfloors INTEGER,
    yearopened SMALLINT,
    yearrenovated SMALLINT,
    photo1 character varying(255),
    photo2 character varying(255),
    photo3 character varying(255),
    photo4 character varying(255),
    photo5 character varying(255),
    overview TEXT,
    rates_from real,
    continent_id INTEGER,
    continent_name character varying(50),
    city_id INTEGER,
    country_id INTEGER,
    number_of_reviews INTEGER,
    rating_average character varying(5),
    rates_currency character varying(3),
    PRIMARY KEY (hotel_id)
);
```


4. import the earlier .csv into table hotels

<br><br>
# 2. TABLE countries

1. select group from table hotels then save into table countries + add primary key
``` bash
CREATE TABLE countries AS
SELECT country_id, country, countryisocode, regexp_replace(regexp_replace(LOWER(regexp_replace(country, '[^A-Za-z0-9]', '-', 'g')), '-+', '-', 'g'), '-$', '', 'g') AS country_key, count(country_id) as count from hotels GROUP BY country_id, country, countryisocode ORDER BY count(country_id) DESC;

ALTER TABLE countries ADD PRIMARY KEY (country_id);
```

2. select top 24 of countries base on count then save into table top24countries + add primary key
``` bash
CREATE TABLE top24countries AS SELECT * FROM countries ORDER BY count DESC LIMIT 24;

ALTER TABLE top24countries ADD PRIMARY KEY (country_id);
```

<br><br>
# 2. TABLE cities

1. select group from table hotels then save into table cities + add primary key
``` bash
CREATE TABLE cities AS
SELECT city_id, city, regexp_replace(regexp_replace(LOWER(regexp_replace(city, '[^A-Za-z0-9]', '-', 'g')), '-+', '-', 'g'), '-$', '', 'g') AS city_key, count(city_id) as count from hotels GROUP BY city_id, city ORDER BY count(city_id) DESC;

ALTER TABLE cities ADD PRIMARY KEY (city_id);
```

2. select top 24 of cities base on count then save into table top24cities + add primary key
``` bash
CREATE TABLE top24cities AS SELECT * FROM cities ORDER BY count DESC LIMIT 24;

ALTER TABLE top24cities ADD PRIMARY KEY (city_id);
```


<br><br>
# 2. TABLE autocomplete

1. select group from table hotels then save into table autocomplete
``` bash
CREATE TABLE autocomplete AS
SELECT hotel_id, hotel_name, city, country, countryisocode, 
regexp_replace(regexp_replace(LOWER(regexp_replace(hotel_name, '[^A-Za-z0-9]', '-', 'g')), '-+', '-', 'g'), '-$', '', 'g') AS hotel_name_key, 
regexp_replace(regexp_replace(LOWER(regexp_replace(city, '[^A-Za-z0-9]', '-', 'g')), '-+', '-', 'g'), '-$', '', 'g') AS city_key, 
rating_average,
star_rating,
number_of_reviews
from hotels ORDER BY rating_average::float DESC, star_rating::float DESC, number_of_reviews DESC;

ALTER TABLE autocomplete ADD PRIMARY KEY (hotel_id);
```