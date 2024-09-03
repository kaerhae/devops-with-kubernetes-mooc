package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type PongResponse struct {
	Pongs int `json:"pongs"`
}

type DBConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

type Config struct {
	PORT      string
	FILE_PATH string
	DBConfig  DBConfig
}

type ServerRouter interface {
	Ping(w http.ResponseWriter, req *http.Request)
	Pongs(w http.ResponseWriter, req *http.Request)
}

type serverRouter struct {
	Config            Config
	CounterRepository CounterRepository
}

type CounterRepository interface {
	InitDb()
	GetCounter() int
	UpdateCounter()
}

type counterRepository struct {
	Config   Config
	Database *sql.DB
}

func InitServerRouter(repo CounterRepository) ServerRouter {
	c := Config{}
	c.GetConfigFromEnv()
	return &serverRouter{
		Config:            c,
		CounterRepository: repo,
	}
}

func InitCounterRepository(c Config, db *sql.DB) CounterRepository {
	return &counterRepository{
		Config:   c,
		Database: db,
	}
}

// Ping implements ServerRouter.
// When receives request, adds to counter, which count amount of pings made.
// Returns amount of Pongs.
func (s *serverRouter) Ping(w http.ResponseWriter, req *http.Request) {
	s.CounterRepository.UpdateCounter()
	counter := s.CounterRepository.GetCounter()
	p := PongResponse{
		Pongs: counter,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// Pongs implements ServerRouter.
// Returns amount of Pings made to application.
func (s *serverRouter) Pongs(w http.ResponseWriter, req *http.Request) {
	count := s.CounterRepository.GetCounter()
	p := PongResponse{
		Pongs: count,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (c *Config) GetConfigFromEnv() {
	c.PORT = os.Getenv("PORT")
	c.FILE_PATH = os.Getenv("PING_FILE_PATH")
	c.DBConfig.Host = os.Getenv("DB_HOST")
	c.DBConfig.Port = os.Getenv("DB_PORT")
	c.DBConfig.User = os.Getenv("DB_USER")
	c.DBConfig.Pass = os.Getenv("DB_PASS")
	c.DBConfig.DbName = os.Getenv("DB_NAME")
	if c.PORT == "" || c.FILE_PATH == "" {
		log.Fatal("Environment variables missing")
	}
}

func CheckDbConfig(c Config) bool {
	if c.DBConfig.Host == "" {
		return false
	}
	if c.DBConfig.Port == "" {
		return false
	}
	if c.DBConfig.User == "" {
		return false
	}
	if c.DBConfig.Pass == "" {
		return false
	}
	if c.DBConfig.DbName == "" {
		return false
	}

	return true
}

func ConnectToDb(c Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBConfig.Host,
		c.DBConfig.Port,
		c.DBConfig.User,
		c.DBConfig.Pass,
		c.DBConfig.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error on opening postgres connection %v", err)
	}
	return db
}

func (c *counterRepository) InitDb() {
	sql := `
		CREATE TABLE IF NOT EXISTS counter (
			id INTEGER PRIMARY KEY,
   			count INTEGER NOT NULL
		);

		INSERT INTO counter (id, count)
		VALUES
		(1, 0)
		ON CONFLICT (id)
		DO NOTHING;
	`

	_, err := c.Database.Exec(sql)
	if err != nil {
		log.Fatal("Error while initializing table", err)
	}
}

func (c *counterRepository) GetCounter() int {
	var count int
	c.Database.QueryRow("SELECT count from counter LIMIT 1").Scan(&count)
	return count
}

func (c *counterRepository) UpdateCounter() {
	sql := `
		UPDATE counter
			SET count = count + 1
		WHERE id=(
			SELECT id FROM counter 
			LIMIT 1
		);
	`

	_, err := c.Database.Exec(sql)
	if err != nil {
		log.Fatal("Error while initializing table", err)
	}
}

func main() {
	c := Config{}
	c.GetConfigFromEnv()
	dbConfigHasValues := CheckDbConfig(c)
	if !dbConfigHasValues {
		log.Fatal("Incorrect DB settings")
	}
	db := ConnectToDb(c)

	repo := InitCounterRepository(c, db)
	repo.InitDb()
	s := InitServerRouter(repo)
	mux := http.NewServeMux()
	mux.HandleFunc("/pingpong", s.Ping)
	mux.HandleFunc("/pongs", s.Pongs)

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", c.PORT),
		mux,
	)
	if err != nil {
		log.Fatal("Error on starting server in port")
	}
}

/*
*	DEPRECATED since exercise 2.07
*
 */

// func createFileIfNotExist(filename string) error {
// 	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
// 		f, err := os.Create(filename)
// 		if err != nil {
// 			return err
// 		}
// 		defer f.Close()

// 		return nil
// 	}

// 	return nil
// }

// func readFile(filename string) (int, error) {
// 	c, err := os.ReadFile(filename)
// 	if err != nil {
// 		return 0, err
// 	}

// 	/* If file is empty, return zero requests */
// 	if len(c) == 0 {
// 		return 0, nil
// 	}

// 	r, err := strconv.Atoi(string(c))
// 	if err != nil {
// 		return 0, err
// 	}
// 	return r, nil
// }

// func overwriteFile(filename string, content string) error {
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = f.Write([]byte(content))
// 	if err != nil {
// 		return err
// 	}
// 	if err := f.Close(); err != nil {
// 		return err
// 	}

// 	return nil
// }
