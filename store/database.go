"database operations(CRUD)"

package store

import(
	"database/sql"
	"log"
	"user-profile-service/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct{
	conn *sql.DB
}

func NewDataBase(dsn string) *Database{
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil{
		log.Fatalf("Failed to connect to database :%v", err)
	}

	_, err = conn.Exec(`
	CREATE TABLE IF NOT EXISTS profiles(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE);
	`)
	if err != nil{
		log.Fatalf("Failed to run migrations: %v",err)
	}
	return &Database{conn: conn}
}

func (db *Database) Close(){
	db.conn.Close()
}

//creating profile to insert into database

func (db *Database)CreateProfile(profile models.Profile)(models.Profile, error){
	result,err := db.conn.Exec("INSERT INTO profiles (name, email) VALUES(? , ?)", profile.Name, profile.Email)
	if err != nil{
		return models.Profile{}, err
	}

	id, _ := result.LastInsertId()
	profile.ID = int(id)
	return profile, nil

}