## Monkey Language

This is a programming language which is designed as a learning project to learn how Interpreters and Compilers work.

## Features

- C-like syntax
- variable building
- integers and booleans
- arithmetic expressions
- built-in functions
- first class and hight order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

## Example

```cpp
let age = 33;
let name = "Dimitri";
let result = 10 * (20 / 2);
```

Arrays:

```cpp
let myArray = [1, 2, 3, 4, 5];
```

Hashes:

```cpp
let user = {"name": "Dimitri", "age": 33};
```

Accessing elements in arrays:

```cpp
myArray[0]; // => 1
```

Accessing elements in hashes:

```cpp
user["name"]; // => "Dimitri"
```

Binding function to statements:

```cpp
let add = fn(a, b) { return a + b };
```

Implicit return:

```cpp
let add = fn(a, b) { a + b };
```

Calling a function:

```cpp
add(1, 2); // => 3
```

A more complex example:

```cpp
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      1
    } else {
      fibonacci(x - 1) + fibonacci(x - 2)
    }
  }
}
```

High order function:

```cpp
let twice = fn(f, x){
  return f(f(x))
}

let addTwo = fn(x){
  return x + 2
}

twice(addTwo, 2) // => 6 : addTwo(addTwo(2))
```

```cpp
let map = fn(arr,f){
  let iter = fn(arr, acc) {
    if(len(arr) == 0) {
      acc;
      } else {
        iter(rest(arr), push(acc, f(first(arr))))
      }
    };
  iter(arr, []);
  };

let timesTwo = fn(x) { x * 2 };
let a = [1, 2, 3, 4];
map(a, timesTwo) // [2, 4, 6, 8]

let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
      if(len(arr) == 0) {
        result;
      } else {
        iter(rest(arr), f(result, first(arr)));
      }
    };

    iter(arr, initial);
  };

let sum = fn(arr) {
    reduce(arr, 0, fn(initial, el) { initial + el } );
  };
let b = [1, 2, 3, 4, 5];

sum(b) // 15
```

### Book

This project was built following the book [Writing an Interpreter in Go](https://interpreterbook.com/) by Thorsten Ball & [Writing a Compiler in Go](https://compilerbook.com/) by Thorsten Ball.
