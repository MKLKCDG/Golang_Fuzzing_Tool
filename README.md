# Golang_Fuzzing_Tool


Golang ile herhangi bir internet sitesinin alt sayfalarını(gizli veya açık) arama yapmamızı kolaylaştıran bir tool 
####
!(hala geliştirme aşamasında)
####
- Help Komutu
    * -help yazarak tool nasıl kullanılacağı hakkında bilgi alabilirsiniz
    <img src="https://github.com/MKLKCDG/Golang_Fuzzing_Tool/blob/main/resimler/-help.png" alt="-help " width="500">
    
- U Komutu
    * -u yazıp tarama yapacağımız internet sitesini belirtiyoruz (örn: go run main.go -u https://www.google.com)
      - eğer istek türünü belirtmezsek varsayılan olarak GET isteği atar.
      - arama yapacağımız sitede kullanacağımız listeyi vermek için -w parametresini kullanıyoruz
      - fc parametresini kullanarak dönen cavapları filitreleyebiliriz
     <img src="https://github.com/MKLKCDG/Golang_Fuzzing_Tool/blob/main/resimler/get.png" alt="get" width="500">

- Post Komutu
    * post isteği atabilmek için -post yazmamız yeterli olacaktır
      - eğer istek türünü belirtmezsek varsayılan olarak GET isteği atar.
      - arama yapacağımız sitede kullanacağımız listeyi vermek için -w parametresini kullanıyoruz
      - fc parametresini kullanarak dönen cavapları filitreleyebiliriz
     <img src="https://github.com/MKLKCDG/Golang_Fuzzing_Tool/blob/main/resimler/post.png" alt="post" width="500">


