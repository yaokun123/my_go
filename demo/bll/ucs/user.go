package ucs

type User struct {
	UserId    int64 `xorm:"'user_id'"`
	user_no   string
	user_name string
	phone     string
}

func (this *User) TableName() string {
	return "tb_user"
}

func (this *User) GetInfoByUserId() *User {
	var user User
	UcsEngine.Alias("o").Where("o.user_id=?", this.UserId).Get(&user)
	return &user
}
