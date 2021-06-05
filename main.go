package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
)

type Employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func addUser(c echo.Context) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	u := new(Employee)
	if err := c.Bind(u); err != nil {
		return err
	}

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `employees` (`name`) VALUES('%s')", u.Name))
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()

	return c.String(http.StatusOK, "ok")
}

func getAllUsers(c echo.Context) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `employees`")
	if err != nil {
		fmt.Println(err, res)
	}
	allUsers := []Employee{}

	for res.Next() {
		var user Employee
		err = res.Scan(&user.Id, &user.Name)
		if err != nil {
			panic(err)
		}

		allUsers = append(allUsers, user)
	}
	return c.JSON(http.StatusCreated, allUsers)

	//return c.String(http.StatusOK, "ok")
}

func deleteUser(c echo.Context) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := c.Param("id")

	delete, err := db.Query(fmt.Sprintf("DELETE FROM `employees` WHERE id = '%v'", id))
	if err != nil {
		fmt.Println(err)
	}
	defer delete.Close()

	return c.String(http.StatusOK, id+" Deleted")
}

func getUser(c echo.Context) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := c.Param("id")

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `employees` WHERE `id` = '%s'", id))
	if err != nil {
		panic(err)
	}
	allUser := Employee{}

	for res.Next() {
		var user Employee
		err = res.Scan(&user.Id, &user.Name)
		if err != nil {
			panic(err)
		}

		allUser = user

	}

	return c.String(http.StatusOK, allUser.Name)
}

func updUser(c echo.Context) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := c.Param("id")

	u := new(Employee)
	if err := c.Bind(u); err != nil {
		return err
	}

	upUser, err := db.Query(fmt.Sprintf("UPDATE `employees` SET `name`='%s' WHERE id='%v'", u.Name, id))
	if err != nil {
		fmt.Println(err)
	}

	defer upUser.Close()

	return c.String(http.StatusOK, id+"this id updated")
}

func main() {
	e := echo.New()
	e.POST("/employee", addUser)
	e.GET("/employee", getAllUsers)
	e.DELETE("/employee/:id", deleteUser)
	e.GET("/employee/:id", getUser)
	e.PUT("/employee/:id", updUser)
	e.Start(":8000")
}
