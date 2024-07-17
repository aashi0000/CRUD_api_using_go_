module example/basic_api/services

go 1.22.5

replace example/basic_api/models => ../models

replace example/basic_api/db => ../db

require (
	example/basic_api/db v0.0.0-00010101000000-000000000000
	example/basic_api/models v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.16.0 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	gorm.io/gorm v1.25.11 // indirect
)
