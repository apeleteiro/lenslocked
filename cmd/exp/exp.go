package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database!")

	// Create some tables...
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
  			id SERIAL PRIMARY KEY,
  			name TEXT,
  			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
  			id SERIAL PRIMARY KEY,
  			user_id INT NOT NULL,
  			amount INT,
  			description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created tables successfully!")

	//// Insert some user...
	//name := "Tono Peleteiro"
	//email := "tono4@peleteiro.eu"
	//row := db.QueryRow(`
	//	INSERT INTO users (name, email)
	//	VALUES ($1, $2) RETURNING id;`, name, email)
	//var id int
	//err = row.Scan(&id)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Created a user with ID:", id)

	// Query one row...
	id := 1
	row := db.QueryRow(`
		SELECT id, name, email
		FROM users
		WHERE id = $1;`, id)
	var name, email string
	err = row.Scan(&id, &name, &email)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("No user found with ID:", id)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("Retrieved user with ID: %d, name: %s, email: %s\n", id, name, email)

	//// Insert some orders...
	//userID := 1
	//for i := 1; i <= 5; i++ {
	//	amount := i * 100
	//	description := fmt.Sprintf("Fake order #%d", i)
	//	_, err = db.Exec(`
	//		INSERT INTO orders (user_id, amount, description)
	//		VALUES ($1, $2, $3);`, userID, amount, description)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//fmt.Println("Inserted 5 fake orders!")

	// Query multiple rows...
	type Order struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}
	var orders []Order
	userID := 1
	rows, err := db.Query(`
		SELECT id, amount, description
		FROM orders
		WHERE user_id = $1;`, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		order.UserID = userID
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Orders: %+v\n", orders)
}
