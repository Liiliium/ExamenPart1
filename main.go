package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//Modelo Persona
type Persona struct {
	Nombre string
	edad   int
	NSS    string
	Sexo   string
	Altura float32
	Peso   float32
}

//Init (Constructor) para el modelo Persona
func (p *Persona) Init(nombre string, edad int, NSS string, sexo string, altura float32, peso float32) {
	p.Nombre = nombre
	p.edad = edad
	p.NSS = NSS
	p.Sexo = sexo
	p.Altura = altura
	p.Peso = peso
}

//interface Personer para el modelo Persona
type Personer interface {
	calcularIMC() int
	esMayorDeEdad() bool
	comprobarSexo() bool
	generarNSS()
	ToString() string
}

//CalcularIMC calcula el IMC de cada persona dependiendo si es Hombre o Mujer y retorna un valor int
func (p Persona) calcularIMC(peso float32, altura float32, sexo string) int {
	imc := peso / (altura * altura)
	if imc < 20 && sexo == "H" {
		return -1
	} else if imc >= 20 && imc <= 25 && sexo == "H" {
		return 0
	} else if imc > 25 && sexo == "H" {
		return 1
	} else if imc < 19 && sexo == "M" {
		return -1
	} else if imc >= 19 && imc <= 24 && sexo == "M" {
		return 0
	} else if imc > 24 && sexo == "M" {
		return 1
	}
	return 2
}

//Init crea la semilla para la runa
func init() {
	rand.Seed(time.Now().UnixNano())
}

//charRuna contiene los caracteres validos para la creacion del NSS
var charRuna = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")

//RandString genera un NSS aleatorio
func (p Persona) generarNSS(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charRuna[rand.Intn(len(charRuna))]
	}
	return string(b)
}

//Imc muestra un mensaje de acuerdo al valor que se reciba como parametro de entrada
func Imc(imc int) {
	fmt.Print("De acuerdo a su IMC usted: ")
	if imc == -1 {
		fmt.Println("Esta por debajo de su peso")
	} else if imc == 0 {
		fmt.Println("Esta en su peso ideal")
	} else if imc == 1 {
		fmt.Println("Tiene sobrepeso")
	}
}

//EsMayorDeEdad indica si la persona es mayor de edad
func (p Persona) EsMayorDeEdad(edad int) bool {
	if edad >= 18 {
		return true
	}
	return false
}

func (p Persona) comprobarSexo(sexo string) bool {
	if sexo == "H" || sexo == "M" {
		return true
	}
	return false
}

//ToString convierte en una sola cadena los valores ingresado al sistema y los retorna
func (p Persona) ToString(per *Persona) string {
	altura := fmt.Sprintf("%f", per.Altura)
	peso := fmt.Sprintf("%f", per.Peso)
	edad := strconv.Itoa(per.edad)
	return "Nombre: " + per.Nombre +
		"\n Edad: " + edad +
		"\n NSS: " + per.NSS +
		"\n Sexo: " + per.Sexo +
		"\n Altura: " + altura +
		"\n Peso: " + peso
}

//variables
var nombre string
var edad int
var NSS string

const sexo = "H"

var peso float32
var altura float32

//Funcion principal main
func main() {
	p := new(Persona)
	fmt.Println("Examen parte 1")
	fmt.Println("Favor de llenar el siguiente cuestionario")
	fmt.Println("Ingresar nombre")
	fmt.Scanf("%s\n", &nombre)
	fmt.Println("Ingresar edad")
	fmt.Scanf("%d\n", &edad)
	NSS = p.generarNSS(8)
	fmt.Println("Ingresar sexo (H/M)")
	fmt.Scanf("%s\n", sexo)
	fmt.Println("Ingresar peso")
	fmt.Scanf("%g\n", &peso)
	fmt.Println("Ingresar altura")
	fmt.Scanf("%g\n", &altura)
	p.Init(nombre, edad, NSS, sexo, peso, altura)
	fmt.Println("El Sexo introducido es correcto? ", p.comprobarSexo(p.Sexo))
	imc := p.calcularIMC(p.Peso, p.Altura, p.Sexo)
	Imc(imc)
	fmt.Println("Usted es mayor de edad: ", p.EsMayorDeEdad(p.edad))
	fmt.Println("Datos ingresados al sistema: \n" + p.ToString(p))
}
