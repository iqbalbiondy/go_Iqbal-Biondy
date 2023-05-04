package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
)

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	e := echo.New()

	// Route untuk halaman utama
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat datang di Aplikasi Manajemen Proyek")
	})

	// Route untuk halaman daftar proyek
	e.GET("/projects", func(c echo.Context) error {
		projects := []Project{}
		rows, err := db.Query("SELECT id, name, description FROM projects")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		defer rows.Close()

		for rows.Next() {
			p := Project{}
			err := rows.Scan(&p.ID, &p.Name, &p.Description)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			}
			projects = append(projects, p)
		}

		return c.JSON(http.StatusOK, projects)
	})

	// Route untuk halaman detail proyek
	e.GET("/projects/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		project := Project{}

		row := db.QueryRow("SELECT id, name, description FROM projects WHERE id=?", id)
		err := row.Scan(&project.ID, &project.Name, &project.Description)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, project)
	})

	// Route untuk membuat proyek baru
	e.POST("/projects", func(c echo.Context) error {
		project := new(Project)
		if err := c.Bind(project); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		result, err := db.Exec("INSERT INTO projects (name, description) VALUES (?, ?)", project.Name, project.Description)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		id, _ := result.LastInsertId()
		project.ID = int(id)

		return c.JSON(http.StatusOK, project)
	})

	// Route untuk mengedit proyek
	e.PUT("/projects/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		project := new(Project)
		if err := c.Bind(project); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		_, err := db.Exec("UPDATE projects SET name=?, description=? WHERE id=?", project.Name, project.Description, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, project)
	})

	// Route untuk menghapus proyek
}