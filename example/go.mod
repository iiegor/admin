module iegor/example

require (
	github.com/go-sql-driver/mysql v1.4.1
	google.golang.org/appengine v1.4.0 // indirect
	iegor/admin v0.0.0
)

replace iegor/admin => ../admin

go 1.13
