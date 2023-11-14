package bd

import (
	"fmt"

	"github.com/villamar32/gambitUser/models"
	"github.com/villamar32/gambitUser/tools"
)

func SignUp(sig models.SignUp) error {

	fmt.Println("Comienza registro")
	err := DbConnect()

	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO user ( User_Email , User_UUID, User_DateAdd ) VALUES ( '" + sig.UserEmail + "' , '" + sig.UserUUID + "', '" + tools.FechaMySQL() + "' )"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Sign up > ejecucion exitosa")
	return nil
}
