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
#### Single line Art text
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
### Multi line Art text
Example
```
go run . -m
```
Write your text in terminal and run with "Ctl + D"

```
[8  ]@|\[2 @]
[7  ]-[2  ][4 @]
[6  ]/7[3  ][4 @]
[5  ]/[4  ][6 @]
[5  ]\-' [8 @]`-[15 _]
[6  ]-[9 @][13  ]/[4  ]\
 [7 _]/[4  ]/_[7  ][6 _]/[6  ]|[10 _]-
/,[10 _]/  `-.[3 _]/,[13 _][10 -]_)
```
Run with "Ctl + D" (Linux/macOS)
Run with "Ctr + Z" then Enter (Windows)

Output
```
       @|\@@
       -  @@@@
      /7   @@@@
     /    @@@@@@
     \-' @@@@@@@@`-_______________
      -@@@@@@@@@             /    \
 _______/    /_       ______/      |__________-
/,__________/  `-.___/,_____________----------_)
```

### Help Flag
```bash
go run . -h
```


## Features and Bonus Functionality
