package main

import (
	"database/sql"
	"log"
  "strconv"
)

type BastardRank struct {
  name string
  count int
}

func getCount(name string, chatId int64) int {
  // open base
	db, err := sql.Open("sqlite3", SystemConfig.dbPath)
	_check(err)
	defer db.Close()

  // some -> Some
	log.Println("Get value to key: " + name)
	// query first row from db
	row := db.QueryRow("select count from bastards where name=$1 and chatID=$2", name, chatId)

  count := -1
	row.Scan(&count)

  log.Println("Found value: " + strconv.Itoa(count))
  return count
}

func addNewName(name string, count int, chatId int64) {
  // open base
	db, err := sql.Open("sqlite3", SystemConfig.dbPath)
	_check(err)
	defer db.Close()

  // some -> Some
	log.Println("Insert new row in DB: " + name + " : " + strconv.Itoa(count))
	// query first row from db
	db.Exec("insert into bastards (name, count, chatID) values ($1, $2, $3)", name, count, chatId)
}

func updateCount(name string, count int, chatId int64) {
  // open base
	db, err := sql.Open("sqlite3", SystemConfig.dbPath)
	_check(err)
	defer db.Close()

  // some -> Some
	log.Println("Update database " + name + " : " + strconv.Itoa(count))
	// query first row from db
	db.Exec("update bastards set count=$1 where name=$2 and chatID=$3", count, name, chatId)
}

func getTop3(chatId int64) []BastardRank {
  // open base
	db, err := sql.Open("sqlite3", SystemConfig.dbPath)
	_check(err)
	defer db.Close()

  // some -> Some
	log.Println("Get top 3")
	// query first row from db
	rows, err := db.Query("SELECT name, count FROM bastards WHERE chatID=$1 ORDER BY count DESC LIMIT 3", chatId)

  rating := []BastardRank{}
  for rows.Next() {
    var name string
    var count int
  	rows.Scan(&name, &count)

    rating = append(rating, BastardRank{name, count})
  }

  for _, rate := range rating {
    log.Println("Name: " + rate.name + " : " + strconv.Itoa(rate.count))
  }

  return rating
}

func newMessageFromBastard(name string, chatId int64) {
  count := getCount(name, chatId)

  if count < 0 {
    addNewName(name, 1, chatId)
    log.Println("Add new name " + name)
  } else {
    updateCount(name, count + 1, chatId)
    log.Println("Update " + name + " count (" + strconv.Itoa(count + 1) + ")")
  }
}
