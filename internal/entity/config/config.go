package entityconfig

const (
	Port      = "7540"
	WebDir    = "./web"
	DBName    = "scheduler.db"
	DBDriver  = "sqlite"
	DBDir     = "./db"
	RowsLimit = 50
	DBCreate  = `CREATE TABLE scheduler (
					id integer PRIMARY KEY,
					date VARCHAR(8) NOT NULL,
					title text NOT NULL,
					comment text,
					repeat VARCHAR(128)
				);
				CREATE INDEX scheduler_date ON "scheduler"("date");`

	Format1 = "20060102"
	Format2 = "02.01.2006"
)

type HTTPConfig struct {
	Port string
}

type DBConfig struct {
	Driver string
	File   string
	Create string
}

type Authentication struct {
	Password string
}
