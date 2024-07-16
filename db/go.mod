module example/basic_api/db

go 1.22.5

require (
	example/basic_api/mocks v0.0.0-00010101000000-000000000000
	example/basic_api/models v0.0.0-00010101000000-000000000000
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.11
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace example/basic_api/mocks => ../mocks

replace example/basic_api/models => ../models
