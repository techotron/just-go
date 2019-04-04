# Compilation

- In the directory where the code is:
```
go build
```
This will output an exe.

- If you have dependancies, you can run this within the directory:
```
go install
``` 
This will compile the application along with any external dependancies that have been referenced into a binary in ../../bin/

# Running the app

- You can run the code without compiling it with:
```
go run ./file.go
```