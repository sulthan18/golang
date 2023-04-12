func main() {
	// Configure the PostgreSQL connection.
	config, err := pgx.ParseConfig("postgresql://maxroach@localhost:26257/bank?sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.maxroach.key&sslcert=certs/client.maxroach.crt")
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	// Set the server name for the TLS config.
	config.TLSConfig.ServerName = "localhost"

	// Connect to the "bank" database.
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer func() {
		// Close the database connection.
		if err := conn.Close(context.Background()); err != nil {
			log.Fatal("error closing the database connection: ", err)
		}
	}()

	// Create the "accounts" table if it doesn't exist.
	if _, err := conn.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal("error creating the accounts table: ", err)
	}

	// Insert two rows into the "accounts" table.
	if _, err := conn.Exec(context.Background(),
		"INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 250)"); err != nil {
		log.Fatal("error inserting data into the accounts table: ", err)
	}
}
