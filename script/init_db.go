package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

// 初始化postgres数据库表

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := "host=localhost user=srbbs password=123456 dbname=srbbs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	if gormdb, err := gorm.Open(postgres.Open(dsn)); err != nil {
		log.Println("postgresql connecting error")
		panic(err)
	}
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		g.GenerateModelAs("user", "User"),
		g.GenerateModelAs("community", "Community"),
		g.GenerateModelAs("post", "Post"),
		g.GenerateModelAs("comment", "Comment"),
	)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
