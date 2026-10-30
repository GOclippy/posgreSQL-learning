package main

import (
	"context"
	"fmt"
	"postgres/postgres/simple_conection"
	"postgres/postgres/simple_sql_example"
)

func main() {
	ctx := context.Background()

	conn, err := simple_conection.CreateConnection(ctx)
	if err != nil{
		panic(err)
	}

	if err := simple_sql_example.CreateTable(ctx, conn); err != nil{
		panic(err)
	}

	fmt.Println("the table was successful created in the database")

	if err := simple_sql_example.InsertRow(ctx,conn); err != nil{
		fmt.Println("an error occured while InsertRow")
		panic(err)
	}
	fmt.Println("the new row was successful insert in the table")
}
