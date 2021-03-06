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
    time text,
	id int,
	doctorid text,
	patientid text,
	duration int,
	timestart text,
	timeend text

);

CREATE TABLE IF NOT EXISTS users (
    id text,
	name text,
	email text,
	password text,
	role text

);


	`

func GiveSchema() *string {
	return &schema
}
