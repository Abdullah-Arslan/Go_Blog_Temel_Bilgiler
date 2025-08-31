package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func main() {

	//Router işlemlerilerini aşağıdaki gibi yapıyoruz.
	r := httprouter.New() //routing işlemlerini aaşağıdaki gibi yapıyoruz.

	//---
	//"/yazilar/go-web-programlar" gibi seo uyumlu bir url yapısı oluşturup bunu aşağıdaki gibi bir yöntem ile router tanımlaması yapıyoruz.
	r.GET("/yazilar/:slug", Anasayfa) //httprouter ı burada "github.com/julienschmidt/httprouter" bu paket ile kullanıyoruz. http.Handle ile aynı yapıya sahip
	//"http://localhost:8080/yazilar/bu-kısma-istene-yazılınca-tarayıcı-gösterir" bu gösterme işlemleri burada tanımlanan "r.GET("/yazilar/:slug", Anasayfa)" parametresi iel oluyor
	//---

	http.ListenAndServe(":8080", r) //buradaki nil yazan kısma r vermeliyiz ki bu da httprouter dan r oluyor site çalışssın.
}

// Url üzerinden bir parametre gelirse " params httprouter.Params" üzerinden gönderecek.
func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) { //httprouter bize gelen parametreleri params üzerinden gönderiyor.
	view, _ := template.ParseFiles("index_2.html")

	//---
	//httprouter bize gelen parametreleri params üzerinden gönderiyor.
	//Yukarıda "r.GET("/yazilar/:slug", Anasayfa)" oluşturduğumuz bu parametreyi yani "/yazilar/:slug" kısmını almak için aşağıdaki yöntemi kullanıyoruz.
	data := params.ByName("slug")
	//---

	view.Execute(w, data)
}
