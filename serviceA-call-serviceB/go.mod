module serviceA-call-serviceB

go 1.15

require (
	github.com/arl/assertgo v0.0.0-20180702120748-a1be5afdc871 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/giant-tech/go-service v0.0.9
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jmoiron/sqlx v1.2.0 // indirect
	go.uber.org/atomic v1.7.0
)

replace github.com/giant-tech/go-service v0.0.9 => ../../go-service
