package routes

import (
	"alta/model"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// // handlerRegister to handle user's registration
// func handlerRegister(c echo.Context) error {
// 	u := new(user.User)

// 	// binding data
// 	err := c.Bind(u)
// 	if err != nil {
// 		return err
// 	}

// 	// call service register

// 	return c.JSON(http.StatusCreated, u)
// }

func (rest *Rest) controllerGetAllUsers(c echo.Context) error {
	// 1. Get data dari database
	users, err := model.GetAll(c, rest.DB)
	// 2. Kalo semisal ada error, maka ...
	if err != nil {
		// 2.1 Kemungkinan yang pertama adalah aplikasi error karena data tidak ditemukan
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Data Not Found"})
		}

		// 2.2 Kemungkinan yang kedua adalah aplikasi error karena fakto eksternal (ex : database mati, atau ada code yang masih error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// 3. kalo tidak ada error, tampilkan user nya
	return c.JSON(http.StatusOK, users)
}

func (rest *Rest) controllerGetUserById(c echo.Context) error {
	// 1. Get Id
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// 2. Get data dari database
	users, err := model.Get(c, rest.DB, id)
	// 2. Kalo semisal ada error, maka ...
	if err != nil {
		// 2.1 Kemungkinan yang pertama adalah aplikasi error karena data tidak ditemukan
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Data Not Found"})
		}

		// 2.2 Kemungkinan yang kedua adalah aplikasi error karena fakto eksternal (ex : database mati, atau ada code yang masih error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// 3. kalo tidak ada error, tampilkan user nya
	return c.JSON(http.StatusOK, users)
}

func (rest *Rest) controllerRegister(c echo.Context) error {
	user := new(model.User)

	// 1. Bind body
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Bind " + err.Error()})
	}

	// 2. create data into database
	err = model.Create(c, rest.DB, *user)
	// 2. Kalo semisal ada error, maka ...
	if err != nil {
		// 2.1 Kemungkinan yang kedua adalah aplikasi error karena fakto eksternal (ex : database mati, atau ada code yang masih error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Register user " + err.Error()})
	}

	// 3. kalo tidak ada error, tampilkan user nya
	return c.JSON(http.StatusCreated, user)
}

func (rest *Rest) controllerUpdate(c echo.Context) error {
	user := new(model.User)

	// 1. Bind Id user
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// 2. Bind body
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Bind " + err.Error()})
	}

	// update data users
	err = model.Update(c, rest.DB, *user, id)
	// 3. Kalo semisal ada error, maka ...
	if err != nil {
		// 3.1 Kemungkinan yang kedua adalah aplikasi error karena fakto eksternal (ex : database mati, atau ada code yang masih error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Update user data " + err.Error()})
	}

	// 4. kalo tidak ada error, tampilkan user nya
	return c.JSON(http.StatusCreated, user)
}

func (rest *Rest) controllerDelete(c echo.Context) error {
	user := new(model.User)

	// 1. Bind Id user
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// delete data users from database
	err := model.Delete(c, rest.DB, id)
	// 3. Kalo semisal ada error, maka ...
	if err != nil {
		// 3.1 Kemungkinan yang kedua adalah aplikasi error karena fakto eksternal (ex : database mati, atau ada code yang masih error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Delete data " + err.Error()})
	}

	// 4. kalo tidak ada error, tampilkan user nya
	return c.JSON(http.StatusCreated, user)
}
