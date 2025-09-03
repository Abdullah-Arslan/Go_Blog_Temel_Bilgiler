package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	//Burada var olan gorm.Model içerisinde bazı değerler otomatik olarak db e ekleniyor.
	gorm.Model

	//Struct ile oluşturulan Username ve Password bu kısımları biz veri tabanına ekliyoruz. Ekleme işlemini de
	//db.Create(&User{Username: "admin", Password: "123456"}) bu yapı ile yapıyoruz.
	Username, Password string
}

func main() {
	//Bu kısımda veri tabanının özellikleri ekleniyor. Aşağıdaki kısım gorm paketinde gelen yapıdır.
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	//Veri tabanı yönlendirmesini aşağıdaki yapı ile yapıyoruz.
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.AutoMigrate(&User{})//db.AutoMigrate(&User{}) otomatik olarak taplo oluşturmayı sağlıyor. Topla yukarıdaki type User içerisine oluşturulacak ve onun altındaki değerlere göre oluşturulması sağlanacak.

	//Database Yeni bir kayıt ekleme işlemi aşağıdaki gibi yapılıyor.
	//db.Create(&User{Username: "admin", Password: "123456"})

	//Buradaki kodlar db deki taployu alıyor var user User buradaki yapıya göre fieldleri dolduruyor.
	//conds değere göre bilgiyi getiriyor.
	//var user User
	//db.First(&user, "1")
	//db.First(&user, "password=?", "123456")//bu şekilde tek bir kaydı alabiliyoruz. Şifre yazılarak username getiriliyor.
	//fmt.Println(user.Username)

	/*
		//Birden fazla kayıt nasıl alınır? aşağıdaki kodlar ile bu seneryayu yapıyoruz.
		var users []User
		db.Find(&users, "password=?", "123456")
		fmt.Println(users)
	*/
	/*
		//Update işlemleri nasıl yapılır?
		var users []User
		db.First(&users, 1)
		db.Model(&users).Update("username", "gorm") //Burada db.First(&users, 1) bu kısımla conds:1 yani id si 1 olanın db.Model(&users).Update("username", "gorm") bu kısımla username mini gorm yap diyoruz. Yani silip yeni verilen değer ile güncellemsi yapılıyor

		var user User
		db.First(&user, 1)
		db.Model(&users).Updates(User{Username: "python", Password: "pip"}) //Birden faz update olacaksa "Updatades", tek bir update olacaksa "Update yi kullanıyoruz.
	*/

	//Delete işlemlerini nasıl yaparız? Aşağıdaki kod blokları ile bu işlemleri gerçekleştiriyoruz.
	//bunda da id vererek silme işlemlerini gerçekleştirebiliriz.
	db.Delete(&User{}, 2) //id 2 olanı sildi.

	//Database de var olan kaydı getirmek için aşağıdaki yapıyı kullanıyoruz.
	var user User
	db.First(&user, 1)
	fmt.Println(user.Username)

}
