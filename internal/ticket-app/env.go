package ticketapp

type env struct{}

func InitEnvs() {}

func GetEnv() env {
	return env{}
}
