package dbconfig

import (
	"fmt"
)

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "senha"
const DbName = "Teste"
const TableName = "transactions"

// String com informações necessárias para conectar ao DB e fazer o perar
var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
