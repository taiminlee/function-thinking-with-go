## What is functional programming?
It is a general-purpose paradigm, characterized by the use of **_mathematical functions_** and the avoidance of **_side effect_**.

A programming style that used only **_pure function_** without **_side effects_**

### What is side effects?
> Are anything a function does other than returning a value. Eg: sending an email, modify global state.

### What is pure functions?
> Are functions that depend only on their arguments and don't have any side effects. Give the same arguments, they will always produce the same return value.

✅ _Function programmers ~~completely avoid side effects and only use pure function~~ will use side effects and impure functions._

## How can we use Function Programming in software?
+ Actions _depend on when it called_ 🔴
+ Calculations _computations from inputs to outputs_ 🟢
+ Data _recorded facts about event_ 🟢
> In general, Function programmer prefer data to calculations and calculations to actions. _Data is easiest to work with_.


| Action                                      | Calculation                                     | Data                                                     | 
|---------------------------------------------|-------------------------------------------------|----------------------------------------------------------|
| Tools for safely changing state over time   | Static analysis to aid correctness              | Ways to organize data for efficient access               |   
| Ways to guarantee ordering                  | Mathematical tools that works well for software | Disciplines to keep records long term                    |  
| Tools to ensure actions happen exactly once | Testing strategies                              | Principles for capturing what is important to using data |  


## What is functional thinking?
### 1. Distinguishing actions, calculations and data.
### 2. Using first-class abstractions
