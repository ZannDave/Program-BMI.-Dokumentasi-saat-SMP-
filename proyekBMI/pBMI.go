package main
import (
    
    "context"
    "fmt"
    "log"
    "proyekBMI/config"
    
)

func categoriesBMI (bmi float32) {
    kategori := "UnderWeight"
    var teksKategori, saran string
    
    if bmi < 18.5 {
        kategori = "UnderWeight"
    } else if bmi >= 18.5 && bmi <= 24.9 {
        kategori = "NormalWeight"
    } else if bmi >= 25.0 && bmi <= 29.9 {
        kategori = "OverWeight"
    } else {
        kategori = "Obesitas"
    }
    
    err := config.DB.QueryRow(
        context.Background(),
        "SELECT TeksKategori, Saran FROM cattbmi WHERE kategori=$1" ,kategori). Scan(&teksKategori, &saran)
        
    if err != nil {
        log.Println("Gagal mengambil data kategori: ",err)
        return
    }
    
    fmt.Println("Kategori : ",kategori)
    fmt.Println("Penjelasan : ", teksKategori)
    fmt.Println("Saran : ", saran)
    
}

func bmiCalculate (tb float32, bb float32) {
    
    tinggiBadanMeter := tb / 100
    hasilBmi := bb / (tinggiBadanMeter * tinggiBadanMeter)
    fmt.Println("Hasil BMI = ", hasilBmi)
    
    categoriesBMI(hasilBmi)
    
}

func bmi () {
    
    fmt.Println("Selamat datang, di program BMI!!!")
    
    var tb float32
    for {
        
        fmt.Println("Silahkan masukkan tinggi badan anda! : ")
        fmt.Scanln(&tb)
        if tb == 0 {
            fmt.Println("•> NOTED : ")
            fmt.Println("TB tidak boleh kosong!!")
        } else if tb  >= 272 {
            fmt.Println("•> NOTED : ")
            fmt.Println("TB Tidak masuk akal!")
        } else {
            break
        }
        
    }
    
    fmt.Println("#----------------#")
    
    var bb float32
    for {
        
        fmt.Println("Silahkan masukkan berat badan anda! : ")
        fmt.Scanln(&bb)
        if bb == 0 {
            fmt.Println("•> NOTED : ")
            fmt.Println("BB tidak boleh kosong!!")
        } else if bb >= 635 {
            fmt.Println("•> NOTED : ")
            fmt.Println("BB Tidak masuk akal!")
        } else {
            break
        }
    
    }
    
    fmt.Println("#----------------#")
    
    if tb != 0 && bb != 0 {
        bmiCalculate(tb, bb)
    } 
    
}

func sayHello (name string) {
    
    fmt.Println("#----------------#")
    fmt.Println("Hello, ", name, "!")
    fmt.Println("#----------------#")
    fmt.Println("Selamat datang di...")
    fmt.Println("Program BMI !!!")
    fmt.Println("#----------------#")
    fmt.Println("Made by : Zann Dave ( 15 th ), with Golang ( Tech )")
    fmt.Println("#----------------#")
    
    for {
        
        var validasi string
        fmt.Println("Apakah anda ingin menggunakannya ( y / n ) ??")
        fmt.Scanln(&validasi)
        fmt.Println("#----------------#")
        
        if validasi == "n" {
            fmt.Println("Baiklah.., ", name, "!")
            fmt.Println("Terimakasih telah menggunakan!!")
            break
        } else if validasi == "y" {
            bmi()
            break
        } else if validasi == "" {
            fmt.Println("Tidak boleh kosong, harus memilih!!")
        } else {
            fmt.Println("Pilih opsi yang tersedia!!")
        }
        
    }
    
}

func getName () {
    
    var nama string
    for {
        fmt.Println("#----------------#")
        fmt.Println("Masukkan nama anda terlebih dahulu ! : ")
        fmt.Scanln(&nama)
        
        if nama != "" {
            break
        }
        
        fmt.Println("•> NOTED :")
        fmt.Println("Nama tidak boleh Kosong!!")
        
    }
    
    sayHello(nama)
    
}

func main () {
    
    config.ConnectCBMI()
    defer config.CloseCBMI()
    
    getName()
    
}
