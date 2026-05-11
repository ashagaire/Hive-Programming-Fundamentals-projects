# Art decoder
A command-line tool that which converts art data into text-based art.

<br/>
<hr style="border: 2px solid blue;">


## Project Overview 

A command line tool, wakes a string as an input and converts it into a piece of text-based art.


## Setup and Installation
Installation 
1.Clone the repository:
```bash  
git clone https://gitea.kood.tech/ashagaire/decoder
cd decoder
```
2.Initialize the Go module:
```bash  
go mod init decoder
```

3.Verify Go is installed:
```bash
go version
```

### Usage Guide
```bash
go run . "Your code for Art"
```

Example
```
go run . "[5 #][5 -_]-[5 #]"
```

Output
```
#####-_-_-_-_-_-#####
```

### Help Flag
```bash
go run . -h
```


## Features and Bonus Functionality
