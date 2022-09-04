package config

type config struct {
	Database struct {
		User                string
		Password            string
		Net                 string
		Addr                string
		DBName              string
		AllowNativePassword bool
		Params              struct {
			ParaseTime string
		}
	}
	Server struct {
		Address string
	}
}
