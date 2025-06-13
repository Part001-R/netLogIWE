package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	_ "modernc.org/sqlite"
)

type MessageT struct {
	TypeMessage   string // type of message - I, W, E
	NameProject   string
	LocationEvent string
	BodyMessage   string
}

// Connect DB
func ConDb(typeDB, nameDB string) (*sql.DB, func() error, error) {

	db, err := sql.Open("sqlite", "iwe.db")
	if err != nil {
		return nil, nil, fmt.Errorf("error connect DB: %v", err)
	}

	closeDB := func() error {
		err := db.Close()
		if err != nil {
			return fmt.Errorf("fault close connect DB: %v", err)
		}
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("fault ping DB: %v", err)
	}

	return db, closeDB, nil
}

// Create tables
func Tables(db *sql.DB, nameTableMain, nameTableI, nameTableW, nameTableE string) error {

	// main table - stores information about the log tables
	if !isValidString(nameTableMain) {
		log.Fatalf("Not Allowed content name table: %s", nameTableMain)
	}
	q := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY,
	nameTableI string UNIQUE,
	nameTableW string UNIQUE,
	nameTableE string UNIQUE,
	timestamp TEXT DEFAULT CURRENT_TIMESTAMP);
	`, nameTableMain)

	_, err := db.Exec(q)
	if err != nil {
		return fmt.Errorf("main table is not created: %v", err)
	}

	// logi table - table for store type I message
	if !isValidString(nameTableI) {
		log.Fatalf("Not Allowed content name table: %s", nameTableI)
	}
	q = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY,
	nameProject string NOT NULL,
	locationEvent string NOT NULL,
	bodyMessage string NOT NULL,
	timestamp TEXT DEFAULT CURRENT_TIMESTAMP);
	`, nameTableI)

	_, err = db.Exec(q)
	if err != nil {
		return fmt.Errorf("logi table is not created: %v", err)
	}

	// logw table - table for store type W message
	if !isValidString(nameTableW) {
		log.Fatalf("Not Allowed content name table: %s", nameTableW)
	}
	q = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY,
	nameProject string NOT NULL,
	locationEvent string NOT NULL,
	bodyMessage string NOT NULL,
	timestamp TEXT DEFAULT CURRENT_TIMESTAMP);
	`, nameTableW)

	_, err = db.Exec(q)
	if err != nil {
		return fmt.Errorf("logw table is not created: %v", err)
	}

	// loge table - table for store type E message
	if !isValidString(nameTableE) {
		log.Fatalf("Not Allowed content name table: %s", nameTableE)
	}
	q = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY,
	nameProject string NOT NULL,
	locationEvent string NOT NULL,
	bodyMessage string NOT NULL,
	timestamp TEXT DEFAULT CURRENT_TIMESTAMP);
	`, nameTableE)

	_, err = db.Exec(q)
	if err != nil {
		return fmt.Errorf("loge table is not created: %v", err)
	}

	return nil
}

// Check string. Return true if valid
func isValidString(tableName string) bool {

	validNameRegex := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

	if !validNameRegex.MatchString(tableName) {
		return false
	}
	maxLength := 63
	if len(tableName) > maxLength {
		return false
	}
	if strings.Contains(tableName, "--") || strings.Contains(tableName, ";") {
		return false
	}

	return true
}

// Saving the received message in the database. Return error
func StoreMessage(db *sql.DB, msg MessageT) error {

	// Validation
	if !isValidString(msg.TypeMessage) {
		log.Fatalf("Not Allowed content in TypeMessage: %s", msg.TypeMessage)
	}
	if !isValidString(msg.NameProject) {
		log.Fatalf("Not Allowed content in NameProject: %s", msg.NameProject)
	}
	if !isValidString(msg.LocationEvent) {
		log.Fatalf("Not Allowed content in LocationEvent: %s", msg.LocationEvent)
	}
	if !isValidString(msg.BodyMessage) {
		log.Fatalf("Not Allowed content in BodyMessage: %s", msg.BodyMessage)
	}

	// Saving
	switch msg.TypeMessage {
	case "I":
		_, err := db.Exec("INSERT INTO logi (nameProject, locationEvent, bodyMessage) VALUES (:project, :location, :body)",
			sql.Named("project", msg.NameProject),
			sql.Named("location", msg.LocationEvent),
			sql.Named("body", msg.BodyMessage))
		if err != nil {
			return nil
		}
	case "W":
		_, err := db.Exec("INSERT INTO logw (nameProject, locationEvent, bodyMessage) VALUES (:project, :location, :body)",
			sql.Named("project", msg.NameProject),
			sql.Named("location", msg.LocationEvent),
			sql.Named("body", msg.BodyMessage))
		if err != nil {
			return nil
		}
	case "E":
		_, err := db.Exec("INSERT INTO loge (nameProject, locationEvent, bodyMessage) VALUES (:project, :location, :body)",
			sql.Named("project", msg.NameProject),
			sql.Named("location", msg.LocationEvent),
			sql.Named("body", msg.BodyMessage))
		if err != nil {
			return nil
		}
	default:
		return errors.New("not allowed type of message")

	}

	return nil
}
