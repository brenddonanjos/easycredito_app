package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DB_NAME = "ve51yyhyfzso671r"
	DB_USER = "cdiafhy8genyu1rc"
	DB_PASS = "ial7a6yjsw34im3y"
	DB_HOST = "wb39lt71kvkgdmw0.cbetxkdyhwsb.us-east-1.rds.amazonaws.com"
	//DB_HOST = "mysql_app"
)

func Open() *sqlx.DB {
	db, err := sqlx.Connect("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":3306)/"+DB_NAME+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}

func Migrations() {
	db := Open()
	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS "+DB_NAME)
	exec(db, "use "+DB_NAME)

	exec(db, `create table if not exists articles (
		id integer auto_increment,
		featured TINYINT(1),
		title varchar(255),
		url TEXT,
		imageUrl TEXT,
		newsSite varchar(255),
		summary TEXT, 		
		spaceFlightId integer,	
		publishedAt DATETIME,
		createdAt DATETIME,
		updatedAt DATETIME,
		PRIMARY KEY (id)
	) ENGINE = innodb`)

	exec(db, `create table if not exists events (
		id integer auto_increment,
		articleId integer,	
		spaceFlightId integer,	
		provider varchar(255),
		createdAt DATETIME,
		updatedAt DATETIME,
		PRIMARY KEY (id)
	) ENGINE = innodb`)

	exec(db, `create table if not exists launches (
		id integer auto_increment,
		articleId integer,	
		spaceFlightId varchar(255),	
		provider varchar(255),
		createdAt DATETIME,
		updatedAt DATETIME,
		PRIMARY KEY (id)
	) ENGINE = innodb`)

	exec(db, `create table if not exists logs (
		id integer auto_increment,
		type varchar(255),	
		message TEXT,
		createdAt DATETIME,
		updatedAt DATETIME,
		PRIMARY KEY (id)
	) ENGINE = innodb`)

}

func exec(db *sqlx.DB, sql string) {
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

}
