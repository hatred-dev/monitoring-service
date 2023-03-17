package validators

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"monitoring-service/database"
)

var ProjectExists validator.Func = func(fl validator.FieldLevel) bool {
	project := fl.Field().String()
	fmt.Println(project)
	exists, _ := database.Conn.ProjectExists(context.Background(), project)
	return exists
}
