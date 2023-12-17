![Logo](https://github.com/AliiAhmadi/APL/assets/107758775/46a9faa6-7bed-4915-879e-56c20bcc3b1e)

# APL

Ahmadi programming language is a language under development. This language is like high-level programming languages used today and it is a interpreted language. 
APL is also dynamically typed language. At this point only support structured and functional paradigm.

> [!WARNING]  
> Do not use this language in your projects. This language is for fun now :)

### Setup and use REPL

```zsh
# Clone APL source code.
git clone https://github.com/AliiAhmadi/APL.git
```

Navigate to APL directory and run tests.
```zsh
go test ./... -v
```

Now compile source and use executable file.
```zsh
go build -o APL
```

### Features

Now is the time to use it and doing some evaluation:

```APL
Ahmadi programming language - Copyright (c) 2023 Ali Ahmadi

APL>> 
```

In APL we can define a new variable of any type with `def` keyword(even functions and closures - in following):

```APL
APL>> def age = 20;
null
APL>> def name = "Ali";
null
APL>> 
```

As you know to see value of a variable just type its identifier:

```APL
APL>> age
20
APL>> name
Ali
APL>> 
```

APL also support numerical calculations and most important operations will work with values or with identifiers (- / * +):

```APL
APL>> 2 * 12
24
APL>> 56 / 7
8
APL>> def x = 100;
null
APL>> def y = 20;
null
APL>> x + y
120
APL>> x * y
2000
APL>> x - y
80
APL>> y - x
-80
APL>> 
```

You can define your `array` and `map` with `def` keyword in APL and use indexing like other programming language to access items:

```APL
APL>> def arr = [100, 300, 200, 500];
null
APL>> arr[0];
100
APL>> 
```

You can define multidimensional arrays:

```APL
APL>> def first = [[1, 2]];
null
APL>> def second = [[4, 16], [5, 25]];
null
APL>> def result = (first[0][1] * second[1][0]) + second[0][0];
null
APL>> result
14
APL>> 
```
As you can see should use parentheses to specify precedence of expressions (by default parser will have precedence like below):
| operations      |
|-----------------|
| Index           |
| Function call   |
| Prefix          |
| Product         |
| Sum             |
| Less \| Greater |
| Equality        |

Now lets work with `map` data type in APL and also combine mutiple string(concatenating):

```APL
APL>> def mp = {"name": "Ali", "family": "Ahmadi", "age": 20, "country": "IR."};
null
APL>> mp["name"]
Ali
APL>> mp["age"]
20
APL>> mp["name"] + " " + mp["family"] + " from " + mp["country"]
Ali Ahmadi from IR.
APL>> 
```
