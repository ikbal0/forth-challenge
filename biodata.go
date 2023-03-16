package main

import (
	"fmt"
	"os"
	"strconv"
)

type person struct {
	Name                 string
	Address              string
	Job                  string
	ReasonToChooseGolang string
}

var students = []person{
	{Name: "MUHAMAD SAGGY", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "Viki Wahyudi", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "ANDREAN DWI ANDARU", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "Saluri Alam", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "FAKHRUL KURNIAWAN", Address: "BANTEN", Job: "", ReasonToChooseGolang: ""},
	{Name: "JIHAN CAMILLIA RITONGA", Address: "SUMATERA UTARA", Job: "", ReasonToChooseGolang: ""},
	{Name: "Ikbal yaduar taupik", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "AMIR MAHMUD", Address: "JAWA TENGAH", Job: "", ReasonToChooseGolang: ""},
	{Name: "FEBRIAN DWI PUTRA", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "Rizky Saputra", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "ANITA LASMARIA SIAGIAN", Address: "SUMATERA UTARA", Job: "", ReasonToChooseGolang: ""},
	{Name: "Ricky Surya Adiputra", Address: "KALIMANTAN BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "YUDA KURNIA NURUL FIKRI", Address: "BANTEN", Job: "", ReasonToChooseGolang: ""},
	{Name: "ADITYA SURYADI", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "RAFLI BIMA PRATANDRA", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "Alvin Martin Djong", Address: "DKI JAKARTA", Job: "", ReasonToChooseGolang: ""},
	{Name: "M FIRMANSYAH ARROZI", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "ADAM ARISTAMA", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "Muhammad Anwar", Address: "SULAWESI SELATAN", Job: "", ReasonToChooseGolang: ""},
	{Name: "Mohamad Zaelani", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "I PUTU WIRA PRATAMA PUTRA", Address: "BALI", Job: "", ReasonToChooseGolang: ""},
	{Name: "IBRAM MUHARAM BACHRI", Address: "DKI JAKARTA", Job: "", ReasonToChooseGolang: ""},
	{Name: "NANDA PRABU ANGGANATA", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "Yazid Ridwan", Address: "KALIMANTAN BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "Ahmad Anshari", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "RIAN FEBRIANSYAH", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "TRISNA BAYU PERMADI", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "TEZA ALFIAN", Address: "DKI JAKARTA", Job: "", ReasonToChooseGolang: ""},
	{Name: "HENDRA ANTONIUS", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "DITHA LOZERA DEVI", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "NURHALISA MADUKUBAH", Address: "SULAWESI TENGGARA", Job: "", ReasonToChooseGolang: ""},
	{Name: "ARGUMELAR PAMUNGKAS", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "RIZKI CAHYO SASONGKO", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "Hamim Yusuf", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "BENI SAFANGAT", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "Ayu Putry Regita Sari", Address: "JAWA TENGAH", Job: "", ReasonToChooseGolang: ""},
	{Name: "MUHAMMAD FAISAL PANGESTU", Address: "BANTEN", Job: "", ReasonToChooseGolang: ""},
	{Name: "FAIZ ROFI HENCYA", Address: "SUMATERA SELATAN", Job: "", ReasonToChooseGolang: ""},
	{Name: "Findryan Kurnia Pradana", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "APRILIA KHOIRUNNISA", Address: "JAWA TENGAH", Job: "", ReasonToChooseGolang: ""},
	{Name: "RAMADINA AINIRIZQI GARNIZAR", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "Akram Firmansyah", Address: "SULAWESI SELATAN", Job: "", ReasonToChooseGolang: ""},
	{Name: "ONAI NADAPDAP", Address: "SUMATERA UTARA", Job: "", ReasonToChooseGolang: ""},
	{Name: "FX. Mario Jevta Prasetio", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "Muhammad Fakhryan Zulfahmi", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "MOHAMAD IQBAL AMRULLAH", Address: "JAWA TIMUR", Job: "", ReasonToChooseGolang: ""},
	{Name: "MUHAMMAD JAUHARI", Address: "SUMATERA SELATAN", Job: "", ReasonToChooseGolang: ""},
	{Name: "Heldi Tio Pratama", Address: "JAWA BARAT", Job: "", ReasonToChooseGolang: ""},
	{Name: "WAHYU HAUZAN RAFI", Address: "JAWA TENGAH", Job: "", ReasonToChooseGolang: ""},
	{Name: "M FARID ALGHOZALI}", Address: "LAMPUNG", Job: "", ReasonToChooseGolang: ""},
}

func main() {
	args := os.Args[1]
	intVar, _ := strconv.Atoi(args)
	findStudent(intVar)
}

func findStudent(args int) {

	var i int = args - 1

	fmt.Println("Nama:", students[i].Name)
	fmt.Println("Alamat:", students[i].Address)
}
