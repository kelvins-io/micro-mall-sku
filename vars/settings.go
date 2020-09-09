package vars

type EmailConfigSettingS struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type MongoDBSettingS struct {
	Uri         string `json:"host"`
	Username    string `json:"user"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	AuthSource  string `json:"auth_source"`
	MaxPoolSize int    `json:"max_pool_size"`
	MinPoolSize int    `json:"min_pool_size"`
}
