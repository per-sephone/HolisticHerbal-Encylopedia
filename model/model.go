package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const DB_FILE = "entries.db"

type Model struct {
	connection *sql.DB
}

type IModel interface {
	Insert(name string, dosage int, uses string, precautions string, preparations string) int64
	SelectByName(name string) *Herb
	Select() []Herb
}

func New() *Model {
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// possibly change uses, precautions, and preparations to JSON
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS herbs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		dosage INT,
		uses TEXT,
		precautions TEXT,
		preparations TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Model{
		connection: db,
	}
}

func (m *Model) Insert(name string, dosage int, uses string, precautions string, preparations string) int64 {
	result, err := m.connection.Exec("INSERT INTO herbs (name, dosage, uses, precautions, preparations) VALUES (?, ?, ?, ?, ?)", name, dosage, uses, precautions, preparations)
	if err != nil {
		log.Println("Error creating new entry:", err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return 0
	}
	return id
}

func (m *Model) SelectByName(name string) (*Herb) {
	row := m.connection.QueryRow("SELECT * FROM herbs WHERE name = ?", name)
	var herb Herb
	err := row.Scan(&herb.Name, &herb.Dosage, &herb.Uses, &herb.Precautions, &herb.Preparations)
	if err != nil {
		log.Println("Error retrieving entry with the given name:", err)
		return nil
	}
	return &herb
}

func (m *Model) Select() ([]Herb) {
	rows, err := m.connection.Query("SELECT * from herbs")
	if err != nil {
		log.Println("Error retrieving entries:", err)
		return nil
	}
	defer rows.Close()
	var herbs []Herb
	for rows.Next() {
		var herb Herb
		err := rows.Scan(&herb.Name, &herb.Dosage, &herb.Uses, &herb.Precautions, &herb.Preparations)
		if err != nil {
			log.Panicln("Error scanning row:", err)
			return nil
		}
		herbs = append(herbs, herb)
	}
	return herbs
}