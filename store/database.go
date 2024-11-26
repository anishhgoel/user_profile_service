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

//Retrieve all profiles from the database

func (db *Database) GetProfiles()([]models.Profile, error){
	rows, err = db.conn.Query("SELECT id, name , email FROM profiles")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var profiles []models.Profile
	for rows.Next(){
		var profile models.Profile
		if err := rows.Scan(&profile.ID, profile.Name, &profile.Email); err != nil{
			return nil, err
		}
		profiles.append(profiles, profile)
	}
	return profiles, nil
}