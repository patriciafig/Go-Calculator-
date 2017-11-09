**CSCI 3415 – Principles of Programming Languages**

Programming Assignment 4 – Go

Due Nov. 27, 2017 

-----------------

>Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language. In this program you will implement a simple calculator in Go that uses reflection to support both.

-----------------
# Part I - Download Go 

From the home page of the Go Programming Language web site, select the Download Go link on the right side of the page. This takes you to a web page with installation instructions for the various platforms – Windows, Mac OS X, or Linux. Then select the Download Go link on that page and select the appropriate download for your platform. Follow the install instructions for your platform.

-----------------

# Part II - Calculator

For this program you will be implementing a simple calculator. By simple I mean it only need to support the binary operators with their standard definitions (i.e., addition, subtraction, multiplication, and division) and parenthesis. Multiplication and division have higher precedence than addition and subtraction, all of them have left to right associativity, and parentheses can be used to override the precedence.The program will accept a string containing the expression to be evaluated and print the result of evaluating the expression. Some simple example are:>2*36>2*3.5-25.0>2*(3.5-2)3.0 You should evaluate the expression in a single left to right scan of the input string, using an operator and an operand stack to hold intermediate results during the evaluation. You may handle parenthesis either by a recursive call or by using the stack.

You can decide how to handle whitespace in an expression. For example, you could completely ignore them, ignore them in limited cases (like around operators and operands), or always return an error. [My preference would be to allow spaces around operators and operands, but not in an operand. So, 20 + 30 is valid, but 2 0 + 3 0 is not.] Specify how you process them.


----------------- 

# Specifics 

- Start with just integers , that is, no parentheses, just integer parsing, and limited error checking. This will get the basic parsing and evaluation algorithm in place.
- Add parenthesis handling.This can either use recursion to evaluate a parenthesized sub-expression, or push the open parenthesis on the operator stack and handle it in the code.
- Add error checking.[This is required for a grade of C.]
- Add floats. Initially, treat everything as floating point – that is, no mixed arithmetic.
[This is required for a grade of B.]
- Add reflection to support mixed integer and floating point calculations. [This is required for a grade of A.]


