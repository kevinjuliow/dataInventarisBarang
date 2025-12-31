package main

import (
	"fmt"

	"github.com/kevinjuliow/dataInventarisBarang/app"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
)

func main() {
	fmt.Println("--- Starting Database Connection Test ---")
	db := app.NewDb()

	err := db.Ping()
	helper.PanicIfError(err)

	fmt.Println("✅ SUCCESS: Database connected successfully!")
	fmt.Println("--- Test Finished ---")
}
