package postgres

func Request() string {
	var request = `
	CREATE TABLE requests (
		id serial primary key,
		request varchar
	)`
	return request
}
