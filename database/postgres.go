package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
)

const connstring = "postgres://Username:password@localhost:port/DBNAME"

func Create() error {
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), `
    CREATE TABLE IF NOT EXISTS logs(
        id SERIAL PRIMARY KEY,
        prompt TEXT NOT NULL,
        response TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
    `)
	if err != nil {
		return err
	}

	return nil
}

func Save(request string, response string) (interface{}, interface{}, error) {
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close(context.Background())
	safe_response := strings.Replace(response, "'", "`", -1)
	safe_request := strings.Replace(request, "'", "`", -1)

	querie := "INSERT INTO logs(prompt, response) VALUES('" + safe_request + "', '" + safe_response + "')"
	_, err = conn.Exec(context.Background(), querie)
	if err != nil {
		fmt.Println(querie)
		return nil, nil, err
	}
	var debug []string
	var debug1 []string
	rows, err := conn.Query(context.Background(), `select prompt,response from logs`)
	if err != nil {
		return nil, nil, err
	}
	for rows.Next() {
		var prompt string
		var response string
		err := rows.Scan(&prompt, &response)
		if err != nil {
			fmt.Printf("Unable to scan row: %v\n", err)
			return nil, nil, err
		}
        debug = append(debug, prompt)
        debug1 = append(debug1, response)

	}
	logger := log.Default()
	logger.Print(debug, debug1)
	return debug1, debug, nil
}
