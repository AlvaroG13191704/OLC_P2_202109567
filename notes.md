## Notas sobre cambios dichas por los auxiliares sobre el enunciado || proyecto || ejecución 
## Canal del ing navarro https://www.youtube.com/@martinfdev/videos
### Lectura del proyecto
- Las funciones en structs estan pendientes
- Las funciones solo pueden retornar valores primitivos como int, float, char, string.

## C3D

- Los metodos son funciones sencillas que no van a tener parametros.
- Los parametros se manejan con el heap y stack.

## Optimización

- Generar algoritmos de optimización de código intermedio.
- Crear una estructura independiente para la opimización 




### Manejo de variables, structs, funciones, metodos , vectores en el heap y stack


### HEAP
En el heap se guardan las variables de larga duración (especificamente las que son globales)

1. Almacenamiento de variables globales
   1. cadenas (strings)
   2. vectores
   3. matrices
   4. structs
   5. nil
2. Funciones como datos (guardar el código de las funciones en el heap)


### STACK
En el stack se guardan las variables de corta duración (especificamente las que son locales)

1. Gestión de ámbitos (Scope Management)
2. Almacenamiento de variables locales
3. Manejo de llamadas a funciones


## Expressions

### Primitive types
- [x] Int
- [x] Float
- [x] Char
- [x] String
- [x] Bool
- [x] Nil
- [ ] Structs
- [ ] Vectors
- [ ] Matrix


### Arithmetic
Start working with arithmetic expressions, play with temps to generate te code, understand how to generate the code for each operation, and how to handle the result of each operation.

- [ ] Sum
- [ ] Substraction
- [ ] Multiplication
- [ ] Division
- [ ] Parenthesis

## Instrucctions

### Working with print 
- Primitives can be printend for example: String, Int, Float, Char
- [x] Print String using %c
- [x] Print Characters using %c
- [x] Print Integers using %d (convert it to int)
- [x] Print Float using %f (convert it to float)
- [x] Print scientific notation using %e (convert it to float)
- [x] Print reduce scientific notation using %g (convert it to float)

### Declaration of variables
- [x] Int
- [x] Float
- [x] Character
- [x] Bool
- [x] String 