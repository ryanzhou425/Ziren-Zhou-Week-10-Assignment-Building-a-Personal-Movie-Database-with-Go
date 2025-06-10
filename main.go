package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", "movies.db")
	defer db.Close()
	fmt.Println("Database connected.")

	db.Exec(`CREATE TABLE IF NOT EXISTS movies (
        id INTEGER PRIMARY KEY,
        name TEXT,
        year INTEGER,
        rank REAL
    );`)

	db.Exec(`CREATE TABLE IF NOT EXISTS movie_genres (
        movie_id INTEGER,
        genre TEXT,
        FOREIGN KEY(movie_id) REFERENCES movies(id)
    );`)

	fmt.Println("Tables created.")

	importMovies(db, "IMDB-movies.csv")
	fmt.Println("Movies imported.")

	importGenres(db, "IMDB-movies_genres.csv")
	fmt.Println("Genres imported.")

	queryGenreCounts(db)
}

func importMovies(db *sql.DB, filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rows, _ := reader.ReadAll()

	// 开启事务
	tx, _ := db.Begin()

	// 预编译 SQL
	stmt, _ := tx.Prepare(`INSERT INTO movies (id, name, year, rank) VALUES (?, ?, ?, ?)`)

	for i, row := range rows {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		name := row[1]
		year, _ := strconv.Atoi(row[2])

		var rank sql.NullFloat64
		if row[3] != "NULL" && row[3] != "" {
			r, _ := strconv.ParseFloat(row[3], 64)
			rank = sql.NullFloat64{Float64: r, Valid: true}
		}

		stmt.Exec(id, name, year, rank)

		if i%10000 == 0 {
			fmt.Printf("   ↳ Inserted %d movie records...\n", i)
		}

	}

	stmt.Close()
	tx.Commit()
}

func importGenres(db *sql.DB, filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rows, _ := reader.ReadAll()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(`INSERT INTO movie_genres (movie_id, genre) VALUES (?, ?)`)

	for i, row := range rows {
		if i == 0 {
			continue
		}
		movieID, _ := strconv.Atoi(row[0])
		genre := row[1]

		stmt.Exec(movieID, genre)

		if i%10000 == 0 {
			fmt.Printf("   ↳ Inserted %d genre records...\n", i)
		}

	}

	stmt.Close()
	tx.Commit()
}

func queryGenreCounts(db *sql.DB) {
	rows, _ := db.Query(`
        SELECT genre, COUNT(*) as total
        FROM movie_genres
        GROUP BY genre
        ORDER BY total DESC
        LIMIT 10;
    `)
	defer rows.Close()

	var genre string
	var count int
	for rows.Next() {
		rows.Scan(&genre, &count)
		fmt.Printf("Genre: %-15s Total Movies: %d\n", genre, count)
	}
}
