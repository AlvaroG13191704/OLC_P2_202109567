// T-swift

// Print primitives
/*

print(10)
print(10.0)
print("Hello")
print('a')
print(true)
print(nil)
print(100.23232132132)
print(3.1413123)
*/

// compiled

printf("%d", (int) 10)

printf("%f", 10.0)

print("%c", 72) // H
printf("%c", 101) // e
printf("%c", 108) // l
printf("%c", 108) // l
printf("%c", 111) // o
printf("%c", 10) // \n

printf("%c", 97) // a

printf("%d", 1) // true
printf("%d", 0) // false

printf("%f",  9999999827968.00) // nil
printf("")


// Declaration of primitives
/*
var bol:Bool = true
var numInt: Int = 10
var numFloat: Float = 10.0
let varChar: Character = 'a'
var varString = "COECYS"

print(bol)
print(numInt)
print(numFloat)
print(varChar)
print(varString)

*/

#include <stdio.h>
float stack[100000];
float heap[100000];
float P;
float H;
float t1;

int main(){
t1 = 1;
// Declaración de la variable 'bol'
stack[(int)P] = t1;
P = P + 1;
t1 = 10;
// Declaración de la variable 'numInt'
stack[(int)P] = t1;
P = P + 1;
t1 = 10.0;
// Declaración de la variable 'numFloat'
stack[(int)P] = t1;
P = P + 1;
// Obtención de la posición inicial del heap
t1 = H;
// Almacenamiento de caracter 'C' en el heap
heap[(int)H] = 67;
H = H + 1;
// Almacenamiento de caracter 'O' en el heap
heap[(int)H] = 79;
H = H + 1;
// Almacenamiento de caracter 'E' en el heap
heap[(int)H] = 69;
H = H + 1;
// Almacenamiento de caracter 'C' en el heap
heap[(int)H] = 67;
H = H + 1;
// Almacenamiento de caracter 'Y' en el heap
heap[(int)H] = 89;
H = H + 1;
// Almacenamiento de caracter 'S' en el heap
heap[(int)H] = 83;
H = H + 1;
// Almacenamiento del caracter nulo en el heap
heap[(int)H] = -1;
H = H + 1;
// Declaración de la variable 'varString'
stack[(int)P] = t1;
P = P + 1;
t2 = 97;
// Declaración de la variable 'varChar'
stack[(int)P] = t2;
P = P + 1;

return 0;
}

// Assigment of primitives
/*
var bol:Bool = true

bol = false

var numInt: Int = 10

numInt = 20

var numFloat: Float = 10.0

numFloat = 20.0

var varChar = 'a'

varChar = 'b'

var varString = "COECYS"
varString = "COECYS2"

print(bol)
print(numInt)
print(numFloat)
print(varChar)

*/
#include <stdio.h>
float stack[100000];
float heap[100000];
float P;
float H;
float t1, t2;


int main(){
t1 = 1;
// Declaración de la variable 'bol'
stack[(int)P] = t1;
P = P + 1;
// Asignación de la variable 'bol'
t2 = 0
stack[(int)0] = t2

return 0;
}





#include <stdio.h>
float stack[1000000]; // STACK
float heap[1000000];  // HEAP
float P;              // Stack pointer 
float H;              // Heap pointer 
float t0, t1, t2, t3;
void printString() {
  L0: 
  t1 = heap[(int)t2]; 
  if (t1 != -1) goto L1; 
  printf("%c",(int) t2);
  t2 = t2 + 1; 
  goto L0; 
  L1: 
  return; 
} 
int main() {
// *** Declaration of a string *** 
t0 = H;
heap[(int)H] = 67; 
H = H + 1;
heap[(int)H] = 79; 
H = H + 1;
heap[(int)H] = 69; 
H = H + 1;
heap[(int)H] = 67; 
H = H + 1;
heap[(int)H] = 89; 
H = H + 1;
heap[(int)H] = 83; 
H = H + 1;
heap[(int)H] = -1; 
H = H + 1;
// Declaration of variable:  varString : String with value:  ' COECYS ' 
stack[(int)P] = t0; 
P = P + 1;

// access to the variable: varString
t3 = stack[(int)0];

// pass to the temp to print
t1 = t3;
// call
printString();
printf("%c", 32);

return 0;
}


 print(7 - (5 + 10 * (2 + 4 * (5 + 2 * 3)) - 8 * 3 * 3) + 50 * (6 * 2))
 print((2 * 2 * 2 * 2) - 9 - (8 - 6 + (3 * 3 - 6 * 5 - 7 - (9 + 7 * 7 * 7) + 10) - 5) + 8 - (6 - 5 * (2 * 3)))
