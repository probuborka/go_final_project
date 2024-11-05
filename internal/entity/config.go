package entity

const (
	Port     = "7540"
	WebDir   = "../web"
	DBName   = "scheduler.db"
	DBDriver = "sqlite"
	DBDir    = "../db"
	DBCreate = `CREATE TABLE scheduler (
					id integer PRIMARY KEY,
					date VARCHAR(8) NOT NULL,
					title text NOT NULL,
					comment text,
					repeat VARCHAR(128)
				);
				CREATE INDEX scheduler_date ON "scheduler"("date");`
)

type HTTPConfig struct {
	Port string
}

type DBConfig struct {
	Driver string
	File   string
	Create string
}
