module client-call-serviceA

go 1.12

require (
	github.com/arl/assertgo v0.0.0-20180702120748-a1be5afdc871 // indirect
	github.com/aurelien-rainone/assertgo v0.0.0-20180702120748-a1be5afdc871
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/giant-tech/go-service v0.0.9
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gogo/protobuf v1.2.1
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	go.uber.org/atomic v1.4.0
)

replace github.com/giant-tech/go-service v0.0.9 => ../../go-service
