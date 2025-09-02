package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {

	//Router işlemlerilerini aşağıdaki gibi yapıyoruz.
	//r := httprouter.New() tanımlaması ile yeni oluşturulan bütün routing işlemlerini r ile tanımlayıp httprouter a gönderileceğini başta belirtiyoruz.
	//sonra r.GET ve r.POST işlemlerinin başlarındaki r değişkenini r := httprouter.New() routerın r si olarak buradan alıyoruz.
	r := httprouter.New() //routing işlemlerini aaşağıdaki gibi yapıyoruz.

	//---
	//"/yazilar/go-web-programlar" gibi seo uyumlu bir url yapısı oluşturup bunu aşağıdaki gibi bir yöntem ile router tanımlaması yapıyoruz.
	//r.GET("/yazilar/:slug", Anasayfa) //httprouter ı burada "github.com/julienschmidt/httprouter" bu paket ile kullanıyoruz. http.Handle ile aynı yapıya sahip
	//"http://localhost:8080/yazilar/bu-kısma-istene-yazılınca-tarayıcı-gösterir" bu gösterme işlemleri burada tanımlanan "r.GET("/yazilar/:slug", Anasayfa)" parametresi iel oluyor
	//---

	r.GET("/", Anasayfa)

	//-----
	/*
		<form action="/deneme" method="get">
		<input type="text" name="username">
		<button type="submit"> Kaydet </button>
		</form>
		Burada yapılan index_2.html içerisinde var olan input işlemleri için yazılan html kodları bu html kısmını go da
		main_routing.go kısmında tanımlıyoruz ve kaydet butonuna basarak gitmesi gereken yere yönlendirmesi sağlanacak
		aşağıdaki yönlendirme işlemi ile bunu gerçekleştiriyoruz.
	*/
	//r.GET("/deneme", Deneme)
	//r.POST("/deneme", Deneme)

	//-----

	//---
	//Burada index_upload.html kısmının tanımlamasını yapıyoruz.
	//r.POST formda yüklenen bir resim yada dosyayı göndermek için post etmek yada kaydetme de yönlendirme routing işlemini gerçekleştiriyor.
	r.POST("/upload", Upload)

	http.ListenAndServe(":8080", r) //buradaki nil yazan kısma r vermeliyiz ki bu da httprouter dan r oluyor site çalışssın.
}

// Url üzerinden bir parametre gelirse " params httprouter.Params" üzerinden gönderecek.
func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) { //httprouter bize gelen parametreleri params üzerinden gönderiyor.
	view, _ := template.ParseFiles("index_upload.html")

	//---
	//httprouter bize gelen parametreleri params üzerinden gönderiyor.
	//Yukarıda "r.GET("/yazilar/:slug", Anasayfa)" oluşturduğumuz bu parametreyi yani "/yazilar/:slug" kısmını almak için aşağıdaki yöntemi kullanıyoruz.
	//data := params.ByName("slug")
	//---

	//view.Execute(w, data)
	view.Execute(w, nil)
}

//Html tarafındaki formları almak için aşağıdaki html kodlarını ve go kodlarını yazıyoruz. Get tanımlaması yukarıda yapıldı
//-----
/*
	<form action="/deneme" method="get">
	<input type="text" name="username">
	<button type="submit"> Kaydet </button>
	</form>
	Burada yapılan index_2.html içerisinde var olan input işlemleri için yazılan html kodları bu html kısmını go da
	main_routing.go kısmında tanımlıyoruz ve kaydet butonuna basarak gitmesi gereken yere yönlendirmesi sağlanacak
	aşağıdaki yönlendirme işlemi ile bunu gerçekleştiriyoruz.
*/

func Deneme(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//Gelen Get isteğini nasıl alıyoruz? aşağıdaki kod ile bunları alıyoruz gelen request leri r (r *http.Request) tutuyoruz.
	//Formdan gelen verinin değerini r.FormValue ile alıyoruz.
	check := r.FormValue("check") //html de name adı altında oluşturdugumuz username mi veriyoruz. r.FormValue bizden string değer istediği kısma
	//burada checkbox olup yada submit olmasını htmlde yapılan tanımlama ile belirliyoruz.
	select_input := r.FormValue("select")

	fmt.Println(check, select_input)
}

/*
Burada oluşturudugumuz index_upload.html dosyasının içerisindeki verileri çekme için yazılan kod blogları var
Yukarıda r.GET ile tanımlaması yapıldı ve aşağıda fonksiyonu yazıldı.

<form action="/upload" method="post" enctype="multipart/form-data">

	<input type="file" name="file">
	<button type="submit">Kaydet</button>

</form>
*/
func Upload(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseMultipartForm(10 << 20)

	//FormFile üzerine gelindiğinde tuttuğu değerleri veriyoruz.
	file, header, _ := r.FormFile("file") //index_upload.html içersinde file name ne ise burada da onu veriyoruz.

	//os.OpenFile ile yeni bir dosya olşuturuyoruz.
	f, _ := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666) //header.Filename header adında dosya oluştur demek.
	io.Copy(f, file)                                                    //burada tam olarak file içeriğini f ye kopyalamış olduk. bunuda io.Copy ile yapıyoruz.

}
