package env
import(
  "github.com/joho/godotenv"
)
type Env struct{
	App_Port int `env:"APP_PORT"`
}

func New() (*Env,error) {
	 err := godotenv.Load()
  if err != nil {
    return nil, err
  }
	_env:= new(Env)
  err = env.Parse(_env)
  if err != nil{
    return nil,err
  }
	return _env,nil
}
