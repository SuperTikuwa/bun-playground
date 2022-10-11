package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type User struct {
	ID          int64  `bun:"id"`
	Name        string `bun:"name"`
	Email       string `bun:"email"`
	CognitoUUID string `bun:"cognito_uuid"`
}

func main() {
	engine, err := sql.Open("mysql", "test:test@tcp([127.0.0.1]:33062)/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(engine, mysqldialect.New())
	// if _, err := db.NewCreateTable().Model((*User)(nil)).Exec(context.Background()); err != nil {
	// 	panic(err)
	// }

	user := &User{
		Name:        "test",
		Email:       "test@example.com",
		CognitoUUID: "test",
	}
	if _, err := db.NewInsert().Model(user).Exec(context.Background()); err != nil {
		panic(err)
	}

}
