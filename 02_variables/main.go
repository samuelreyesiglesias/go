package main;
//use pointers to store the memory address of a variable

import "fmt"

func main() {
  //primer manera de definir variables
  var nombre, apellido string
  nombre= "Belen";
  apellido = "Evangeline";


  fmt.Println(nombre, apellido)

  //segunda manera de definir variables
  apellido2 := "Reyes"

  fmt.Println(apellido2)

  //tercera manera de definir variables
  usuario1 , usuario2 := "Santiago", "Alexander"

  fmt.Println(usuario1," ", usuario2)

  //declaracion y asignacion 
  language, year := "Go", 2009

  
  
  fmt.Println(language, " ", year);
  
  var role string = "Admin";
  var age int =25;
  var salary float32 = 35000;
  var isConfirmed bool = true; 
  var description string = "This is Admin user";
  fmt.Println(role, " ", age, " ", salary, " ", isConfirmed, " ", description);

  var ceo, company string = "Steve Jobs", "Apple";
  fmt.Println(ceo, company );


  var poet, book string = "William Blake", "Songs of Innocence and of Experience";
  fmt.Println(poet, book);


  var song, artist , album , country string = "Viva La Vida", "Coldplay", "Viva La Vida or Death and All His Friends", "United Kingdom";
  fmt.Println(song, artist, album, country);

  //one single declaration but different types
  var (device string = "iPhone"; year2 int = 2007; price float32 = 499.99; isAvailable bool = true;);
  fmt.Println(device, year2, price, isAvailable);


  var (hisName string ="Belen"; herAge int = 25; herSalary float32 = 35000; status bool = true;);
  fmt.Println(hisName, herAge, herSalary, status);

  //types of data variables  : string, int , float32, bool , byte, rune, complex64, complex128, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, uintptr , array, slice, struct, pointer, function, interface, map, channel, etc.
  var pointerName *string   = &hisName;
  fmt.Println(pointerName);

  var pointerAge *int = &herAge;
  fmt.Println(pointerAge);

  //use Person struct
  var person1 Person = Person{name: "Belen", age: 25, salary: 35000};
  fmt.Println(person1);

   

}

type Person struct{
  name string;
  age int;
  salary float32;
}

//create interface IProgrammer , and a class Programmer that implements the interface IProgrammer
type IProgrammer interface{
  writeCode() string;
}

 