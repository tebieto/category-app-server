package database


func Available() *Config{

return FetchDB1()
	


}


func FetchDB1() *Config{

	db := new(Config)

	db.Host = "127.0.0.1"
	db.Port = "3306"
	db.User = "root"
	db.Password = "1223454allu4%"
	db.Database = "core"

	return db


}