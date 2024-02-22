package initialize

func Init() {
	InitConfig()
	InitLogger()
	InitDB()

	InitRedis()
	InitValidator()

}
