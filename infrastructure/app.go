package boostrap

import "fmt"

type Application struct {
	Env *Env
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	return *app
}

func (App *Application) closeDBConnection() {
	fmt.Println("DB connection close called.")
}
