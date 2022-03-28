package datab

var schema = `

CREATE TABLE IF NOT EXISTS doctor (
    id text,
	name text,
	role text
);

CREATE TABLE IF NOT EXISTS patient (
    id text,
	name text,
	role text
);

CREATE TABLE IF NOT EXISTS appointment (
    id int,
	doctorid text,
	patientid text,
	duration int,
	timestart text,
	timeend text

);

CREATE TABLE IF NOT EXISTS users (
    id int,
	name text,
	email text,
	password int,
	role text

);


	`
