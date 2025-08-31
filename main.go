package main

import (
	"html/template"
	"net/http"
)

func main() {
	//index.html adında oluşturulan bir dosya yada html dosyalarını çağırmak port üzeirnden çalıştırmak için handle içindeki FileServer komutunu kullanıyoruz. .Dir ile de dosyaları al diyoruz.
	//http.Handle("/", http.FileServer(http.Dir("")))//bu şekilde statik dosyaları çalıştırıyoryuz. Dinamik için farklı yöntemler kullanıyourz.

	//bir fonksiyon ile html dosyalarını çağırmak istersek http.HandleFucn komutu kullanıyoruz.
	http.HandleFunc("/", Anasayfa)   //burada anasayfa ile ilgili bir fonksiyon oluşturduk bunun func yazılarak çağırma işlemleri yaplabilir.
	http.HandleFunc("/detay", Detay) //html olarak arama çubuguna localhost:8080/detay yazarsak gelir. istenen kısmı index dışında bu şekilde veriyoruz.
	http.HandleFunc("/sidebar", Sidebar)

	//addr dediği 8080 portu bu nedenle tırnaklar içerisinde bu adresi veriyoruz handler şuanda yok bu nedenle nil olarak giriyoruz.
	//Öncelikle 8080 portunu açıyoruz bu işlemi yapmanın yolu http paketinden ListnAndServe rı çağırmaktır.
	http.ListenAndServe("localhost:8080", nil) //handler şuanda nil olarak veriliyor ileride data gibi atamalarda http yönlendirmesi yapılacaktır.

}

// http.Request gelen istekleri almamızı formdan gelen dataları post iseklerini almamızı sağlıyor.
// http.ResponseWriter ise http.ResponseWriter, HTTP yanıtının başlığını ve içeriğini yazmak için kullanılır. İstemciye gönderilecek olan HTTP cevabını bu arayüz üzerinden oluşturursun.
// Aşağıdaki tüm yapıların olabilmesi için http.HandleFunc na ihtiyacımız var yani HandleFunc oluşturduktan sonra fonksiyon oluşturarak aşağıdaki yapıyı oluşturuyoruz.
// HandleFunc iki tane fix aldığı kesin olarak alması gereken parametreler var onlar w http.ResponseWriter, r *http.Request bunlardır.
// response müşteriye dönüt vermek için kullanılır ve cevap demektir.
// r.GET ve r.POST alabilmek için r *http.Request parametresini kullanıyoruz.
func Anasayfa(w http.ResponseWriter, r *http.Request) { //burada yukarıdaki http.HandleFunc ile oluşturulan Anasayfa fonksiyonunu çağırıyoruz.

	//fmt.Println("Birileri Anasayfaya Bağlandı")

	//html dosyasını almak için aşağıdaki yöntemi kullanıyoruz.
	view, _ := template.ParseFiles("index.html", "navbar.html", "sidebar.html") //html dosyalarını bu şekilde çağırıyoruz. hangi dosyaların parse edilmesi hmtl olarak isteniyorsa onlar bu şekilde yazılıyor.

	//----
	//Go dan Bir veri göndereceğimiz zaman yapılacak işlemler sırası ile şu şekildedir. Burada navbar içerisine bir veri göndermesi yaptık
	//bu veri göndermesi genel tanımda şu şekilde yapılıyor html içerisinde {{ template "navbar"  .}} bu index.html de navbar kısmına yapılan tanımlama
	//bu kısım ise navbar.html içerisine yapılan tanımlama <h1>{{.}}</h1> nokta ile çağırma işlemlerini yapıyoruz.
	data := "Go'dan Gelen Veri"
	//---

	//----
	//şimdi biz navbar.html i parse ettik fakat go hangisini önce çalıştıracak bunu anlayamayz onun için
	//view.Execute(w, nil) yerine farklı bir fonksiyona ihtiyacımız var oda view.ExecyteTamplate dir.
	view.ExecuteTemplate(w, "anasayfa", data) //name istenen yere benim hangi template önce çalıştırmak istediğimi soruyor bunu tanımlamamız lazım ama önce html de bunun tanımlaması yapılmalı.
	//"anasayfa" burada ilk önce çalıştırılacak html sayfası verilmiş oluyor oda {{define "anasayfa"}} hangi html de tanıtılmış ise odur.
	//---

	//****
	//V1
	//index.html tarafında yada html tarafından her hangibir dosyasında  atama yapılırken süslü parantezler içerisine {{}} atama yapılarak go dan veril gönderilir.
	//Örneğin index.html tarafına datayı göndermek için yapılması gereken {{.}} nokta koyarak veriyi çekmektir.
	//html tarafına veril göndermek istenseydi {{.veril}} olarak yada başka isimde nokta ile atama yapılarak gönderme yapılacaktı.
	//tüm datayı almak için {{.}} kullanırız html tarafında
	//**data := "Go'dan Gelen Veri..."
	//***

	//----
	//V2
	//html tarafına veri göndermek istenseydi {{.veril}} olarak yada başka isimde nokta ile atama yapılarak gönderme yapılacaktır.
	//html tarafına veri gönderme da {{}} parantezler içerisine data["veri"] verilen isim ne ise {{.veril}} yada başka isimde yazıyoruz.
	//Bir değer değilde bir çok değeri göndermek veri tabanından çekilekn bir çok veriyi göndermek için make(map) leri kullanıyoruz.
	//**data := make(map[string]interface{})
	//**data["veri"] = "Godan gelen bir başka veri"
	//-----

	/*
		{{range $index,$value:=.Sayilar}}
		<h1>{{$value}}</h1>
		{{end}}
	*/
	//Burada html dosyasında döngü ile sayıları alt alta yazdırma ve bunları çağırmak için yukarıdaki range yöntemini ve veri çağırma koşullarını kullanıyoruz.
	//data := make(map[string]interface{})      //bir kere bu kısımda data bu şekilde oluşturulduktan sonra
	//data["Sayilar"] = []int{1, 2, 3, 4, 5, 6} //diğer kısımlarda datayı bu şekilde çağırıyoruz. Hepsini çünkü bir kere make oluşturulduğu için

	/*
		Sayının eşit olup olmadığını aşağıdaki yöntem ile yapıyoruz.
		{{ if eq .sayi 10 }}
		        <h1>Sayi 10 na eşittir.</h1>
		        {{else}}
		        <h1>Sayı 10 na eşit değildir.</h1>
		        {{end}}

	*/

	//data["sayi"] = 12
	/*
		//html tarafında if, else ile döngü şeklinde bir yapı oluşturmak için
		{{ if .is_admin }}
		<h1>Bu kullancı bir admindir.</h1>
		{{else}}
		<h1>Bu kullanıcı admin değildir...</h1>
		{{end}}
		bu yapıyı kullanıyoruz.

		data["is_admin"] = false

	*/

	/*
		//V3
		//type struct oluşturarak da data ya veri göndermesi oradan da html tarafına gönderme yapabiliriz. Buda version 3 olarak başka bir yöntemdir.
		data := Data{
			Veri:    "Godan Gelen Daha Başka Veri",
			Sayilar: []int{1, 2, 3, 4, 5, 6},
		}

	*/

	//Execute bir şablonu (template) çalıştırır ve çıktısını belirtilen hedefe (örneğin bir HTTP cevabı, dosya ya da konsol) yazar.
	//view.Execute(w, nil) → Şablonu işler, çıktı w'ye yazılır, fakat veri gönderilmediği için değişken placeholder’lar boş kalır.
	//view.Execute(w, data) //"w" burda ResponseWriter olarak veriyoruz verilerin istemci tarayıcıya gidecek verileri bu şekilde yazdırıyoruz.

}

/*
// V3
// type structlar kullanılarak da veri göndermesi yapılabilir. Burada oluşturulan structları yukarıda data olarak temkar tanımlama yapmak gerekiyor.
// Yapıları kullanarak html tarafına veri göndermesi yapabiliriz.
type Data struct {
	Veri    string
	Sayilar []int
}

*/

func Detay(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("detay.html", "navbar.html", "sidebar.html") //burada navbar.html parse ediliyor ki gelsin.
	view.ExecuteTemplate(w, "detay", nil)                                       //burası da aynı şekilde detay.html sayfası çalıştırılması için template olarak tanımlaması yapılyor.

	//view.Execute(w, nil)

}

func Sidebar(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("sidebar.html", "navbar.html", "detay.html")
	view.ExecuteTemplate(w, "sidebar", nil)

}
