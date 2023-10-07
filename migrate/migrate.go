package main

import (
	"github.com/AnotherFurakam/samamander-api/internal"
	"github.com/AnotherFurakam/samamander-api/pkg/database"
)

func init() {
	database.ConnectToDatabase()
}

func main() {
	internal.Migration(database.DB)
}
