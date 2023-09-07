# BlackHatGo
### Disclaimer - this repository is only for educational purposes

https://nostarch.com/blackhatgo

All the nifty little things from the book BlackHatGo


    Packages: A collection of Go source files in the same directory that are compiled together. They can be used to create libraries or executables.

    Functions: A block of code that performs a specific task. Functions can take zero or more arguments and can return zero or more values.

    Structs: A composite data type that groups together zero or more values with different types under a single name. Structs are used to represent objects with multiple properties.

    Methods: Functions that are associated with a specific type, and can access and modify the properties of that type. Methods are defined using the dot notation.

    Interfaces: A collection of method signatures that define a set of behaviors. Types can implement interfaces to indicate that they support those behaviors.

    Goroutines: A lightweight thread of execution that can be started with the "go" keyword. Goroutines allow concurrent execution of code without the overhead of traditional threads.

    Channels: A typed conduit through which Goroutines communicate. Channels can be used to synchronize and share data between Goroutines.

    Pointers: A variable that stores the memory address of another variable. Pointers can be used to pass values by reference, to modify values indirectly, and to reduce memory usage.

    Defer: A statement that schedules a function call to be executed after the surrounding function returns. Defer is commonly used to clean up resources or to ensure that certain actions always occur.

GOROOT is for compiler/tools that comes from go installation.
GOPATH is for your own go projects / 3rd party libraries (downloaded with "go get").
