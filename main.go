package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//println("-----TXT------")
	//leerArchivo(("./archivo.txt"))
	//println("-----CSV------")
	//readCsv("./prueba.csv")
	// ejecutar funcion cada 2 segundos
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//go func() {
	//	for {
	//		println("-----CSV------")
	//		readCsv("./prueba.csv")
	//		time.Sleep(time.Second * 2)
	//	}
	//}()
	//time.Sleep(time.Second * 10)
	sumar(array)
}

func leerArchivo(nombre string) {
  archivo, err := os.Open(nombre)
  if err != nil {
	  fmt.Println(err)
	  return
  }
  defer archivo.Close()
  scanner := bufio.NewScanner(archivo)
  for scanner.Scan() {
	  fmt.Println(scanner.Text())
  }
}
// leer csv
func readCsv(nombre string)  {
    archivo, err := os.Open(nombre)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer archivo.Close()
    scanner := bufio.NewScanner(archivo)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

}
// sumar datos de un array
func sumar(numeros []int) int {
    total := 0
    for _, numero := range numeros {
        total += numero
    }
	println(total)
    return total
}

