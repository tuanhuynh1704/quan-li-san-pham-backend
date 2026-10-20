package main

import (
	"os"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "1704"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin!",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"username":  "tuna",
			"full_name": "Huynh Anh Tuan",
		})
	})

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"context"
// 	"log"
// 	// "os"
// 	"quan-li-san-pham/internal/databases"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	// "errors"
// 	// "database/sql"
// )

// func main() {
// 	_ = godotenv.Load()

// 	ctx := context.Background()

// 	db, err := databases.NewPostgres(ctx)
// 	if err != nil {
// 		log.Fatalf("Database connection failed: %v", err)
// 	}
// 	defer db.Close()

// 	r := gin.Default()

// 	r.GET("/db/schema", func(c *gin.Context) {
// 		rows, err := db.QueryContext(c.Request.Context(), `
// 		SELECT
// 			table_schema,
// 			table_name,
// 			column_name,
// 			data_type,
// 			is_nullable
// 		FROM information_schema.columns
// 		WHERE table_schema NOT IN ('pg_catalog', 'information_schema')
// 		ORDER BY table_schema, table_name, ordinal_position
// 	`)
// 		if err != nil {
// 			c.JSON(500, gin.H{"error": "Failed to query database schema"})
// 			return
// 		}
// 		defer rows.Close()

// 		type Column struct {
// 			Schema   string `json:"schema"`
// 			Table    string `json:"table"`
// 			Column   string `json:"column"`
// 			DataType string `json:"data_type"`
// 			Nullable string `json:"nullable"`
// 		}

// 		columns := make([]Column, 0)

// 		for rows.Next() {
// 			var col Column
// 			if err := rows.Scan(
// 				&col.Schema,
// 				&col.Table,
// 				&col.Column,
// 				&col.DataType,
// 				&col.Nullable,
// 			); err != nil {
// 				c.JSON(500, gin.H{"error": "Failed to scan schema"})
// 				return
// 			}

// 			columns = append(columns, col)
// 		}

// 		if err := rows.Err(); err != nil {
// 			c.JSON(500, gin.H{"error": "Failed to read schema"})
// 			return
// 		}

// 		c.JSON(200, gin.H{
// 			"count": len(columns),
// 			"data":  columns,
// 		})
// 	})

// 	r.GET("/departments", func(c *gin.Context) {
// 		rows, err := db.QueryContext(c.Request.Context(), `
// 		SELECT "id", "name", "description", "createdAt", "updatedAt"
// 		FROM public."Department"
// 		ORDER BY "createdAt" DESC
// 	`)
// 		if err != nil {
// 			c.JSON(500, gin.H{"error": "Failed to query departments"})
// 			return
// 		}
// 		defer rows.Close()

// 		type Department struct {
// 			ID          string    `json:"id"`
// 			Name        string    `json:"name"`
// 			Description *string   `json:"description"`
// 			CreatedAt   time.Time `json:"createdAt"`
// 			UpdatedAt   time.Time `json:"updatedAt"`
// 		}

// 		departments := make([]Department, 0)

// 		for rows.Next() {
// 			var d Department

// 			if err := rows.Scan(
// 				&d.ID,
// 				&d.Name,
// 				&d.Description,
// 				&d.CreatedAt,
// 				&d.UpdatedAt,
// 			); err != nil {
// 				c.JSON(500, gin.H{"error": "Failed to scan department"})
// 				return
// 			}

// 			departments = append(departments, d)
// 		}

// 		if err := rows.Err(); err != nil {
// 			c.JSON(500, gin.H{"error": "Failed to read departments"})
// 			return
// 		}

// 		c.JSON(200, gin.H{
// 			"count": len(departments),
// 			"data":  departments,
// 		})
// 	})

// 	r.GET("/health", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"status": "ok",
// 		})
// 	})

// 	r.GET("/health/db", func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(
// 			c.Request.Context(),
// 			2*time.Second,
// 		)
// 		defer cancel()

// 		if err := db.PingContext(ctx); err != nil {
// 			c.JSON(503, gin.H{
// 				"status": "unavailable",
// 			})
// 			return
// 		}

// 		c.JSON(200, gin.H{
// 			"status": "connected",
// 		})
// 	})

// 	if err := r.Run(":8080"); err != nil {
// 		log.Fatal(err)
// 	}
// }
