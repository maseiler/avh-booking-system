package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	loginInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv("AVHBS_DB_USER"), os.Getenv("AVHBS_DB_PASS"), os.Getenv("AVHBS_DB_IP"), os.Getenv("AVHBS_DB_PORT")) + os.Getenv("AVHBS_DB_NAME") + "?parseTime=true"
	fmt.Println("Database Login Info:")
	fmt.Println(loginInfo)
	//loginInfo = loginInfo + os.Getenv("AVHBS_DB_NAME") + "?parseTime=true"
	db, _ = sql.Open("mysql", loginInfo)

	insertClients()
	insertPasswords()
	insertUsers()
	insertItems()
	insertBookings()
	insertFavoriteItems()
	insertFeedback()
}

func insertClients() {
	db.Exec("INSERT INTO clients VALUES('Carmerstraße', NOW() - INTERVAL 6 DAY);")
	db.Exec("INSERT INTO clients VALUES('Pichelütte', NOW() - INTERVAL 6 DAY);")
}

func insertPasswords() {
	db.Exec("INSERT INTO passwords VALUES('secret', NOW() - INTERVAL 6 DAY);")
}

func insertUsers() {
	db.Exec("INSERT INTO users VALUES(0, NOW() - INTERVAL 6 DAY, 'Alpha', 'Ahmed', 'Angus', '', 'Aktiv B', '', '', -19, -50);")
	db.Exec("INSERT INTO users VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 2 MINUTE, 'Beta', 'Bert', '', '', 'Aktiv B', '', '', -8.05, -50);")
	db.Exec("INSERT INTO users VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 4 MINUTE, 'Gamma', '', 'Gimel', '', 'AH', '', '', 0, -100);")
	db.Exec("INSERT INTO users VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 7 MINUTE, '', 'Delta', 'Dalet', '', 'Gast', 'delta@dalet.com', '+49 123 456', 20, -50);")
	db.Exec("INSERT INTO users VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 8 MINUTE, '', 'Epsilon', 'Epikur', '', 'Gast', 'epikur@athens.gr', '', -110, -50);")
	db.Exec("INSERT INTO users VALUES(0, NOW() - INTERVAL 5 DAY, 'Zeta', '', '', '', 'AH', '', '', 0, -100);")
	db.Exec("INSERT INTO users VALUES(0, NOW() - INTERVAL 4 DAY, '', 'Emily', 'Everest', 'Eta', 'Steganleger', 'Sagarmatha@web.de', '', 0, -50);")
	db.Exec("INSERT INTO users VALUES(0, NOW() - INTERVAL 4 DAY + INTERVAL 4 MINUTE, '', 'Timothy', 'Tiger', 'Theta', 'Steganleger', '', '+41 456 789', 0, -50);")
}

func insertItems() {
	db.Exec("INSERT INTO items VALUES(0, NOW() - INTERVAL 6 DAY, 'Agua', 'non-alcoholic', 1, 'l', 1);")
	db.Exec("INSERT INTO items VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 2 MINUTE, 'Beer', 'alcoholic', 0.5, 'l', 1.5);")
	db.Exec("INSERT INTO items VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 3 MINUTE, 'Beer', 'alcoholic', 0.3, 'l', 0.9);")
	db.Exec("INSERT INTO items VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 20 MINUTE, 'Cauliflower', 'food', 1, 'piece', 2.35);")
	db.Exec("INSERT INTO items VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 25 MINUTE, 'Date Juice', 'non-alcoholic', 0.5, 'l', 5);")
	db.Exec("INSERT INTO items VALUES(0, NOW() - INTERVAL 5 DAY + INTERVAL 8 MINUTE, 'Eggnog', 'non-alcoholic', 0.2, 'l', 2);")
}

func insertBookings() {
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 6 DAY + INTERVAL 2 MINUTE, 1, 1, 1, -1, 'Part 1/1', '');")
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 5 DAY + INTERVAL 33 MINUTE, 2, 3, 2, -1.8, 'Part 1/1', '');")
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 5 DAY + INTERVAL 90 MINUTE, 1, 2, 12, -18, 'Part 1/1', '');")
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 4 DAY + INTERVAL 12 MINUTE, 3, 5, 2, -10, 'Part 1/2', '');")
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 4 DAY + INTERVAL 12 MINUTE, 3, 6, 1, -2, 'Part 2/2', '');")
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 4 DAY + INTERVAL 55 MINUTE, 3, 0, 1, 12, 'Payment Full', 'Cash');")
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 2 DAY + INTERVAL 44 MINUTE, 4, 0, 1, 20, 'Payment Part', 'Cash');")
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 1 DAY + INTERVAL 82 MINUTE, 5, 5, 22, -110, 'Part 1/1', '');")
	db.Exec("INSERT INTO bookings VALUES(0, NOW() - INTERVAL 30 MINUTE, 2, 4, 3, -7.05, 'Part 1/1', '');")
}

func insertFavoriteItems() {
	db.Exec("INSERT INTO favorite_items VALUES(1, 1, 1);")
	db.Exec("INSERT INTO favorite_items VALUES(2, 3, 2);")
	db.Exec("INSERT INTO favorite_items VALUES(1, 2, 12);")
	db.Exec("INSERT INTO favorite_items VALUES(3, 5, 2);")
	db.Exec("INSERT INTO favorite_items VALUES(3, 6, 1);")
	db.Exec("INSERT INTO favorite_items VALUES(5, 5, 22);")
	db.Exec("INSERT INTO favorite_items VALUES(2, 4, 3);")
}

func insertFeedback() {
	db.Exec("INSERT INTO feedback VALUES(0, NOW() - INTERVAL 4 DAY + INTERVAL 2 MINUTE, 'AVHBS = AVH Bullshit. Lol.', 'Troll');")
	db.Exec("INSERT INTO feedback VALUES(0, NOW() - INTERVAL 1 DAY + INTERVAL 100 MINUTE, 'Great work!', '');")
}
