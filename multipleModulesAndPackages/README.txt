edsels-mbp:examples edsellandicho$ cd multipleModulesAndPackages/
edsels-mbp:multipleModulesAndPackages edsellandicho$ pwd
/Users/edsellandicho/goProjects/tutorials/examples/multipleModulesAndPackages
edsels-mbp:multipleModulesAndPackages edsellandicho$ ls
greetings
edsels-mbp:multipleModulesAndPackages edsellandicho$ cd greetings/
edsels-mbp:greetings edsellandicho$ go mod init examples.com/greetings
go: creating new go.mod: module examples.com/greetings
edsels-mbp:greetings edsellandicho$ pwd
/Users/edsellandicho/goProjects/tutorials/examples/multipleModulesAndPackages/greetings
edsels-mbp:greetings edsellandicho$ cd ../hello/
edsels-mbp:hello edsellandicho$ go mod init examples.com/hello
go: creating new go.mod: module examples.com/hello
edsels-mbp:hello edsellandicho$ pwd
/Users/edsellandicho/goProjects/tutorials/examples/multipleModulesAndPackages/hello
edsels-mbp:hello edsellandicho$ go mod init -replace examples.com/greetings=../greetings
flag provided but not defined: -replace
usage: go mod init [module-path]
Run 'go help mod init' for details.
edsels-mbp:hello edsellandicho$ go mod edit -replace examples.com/greetings=../greetings
edsels-mbp:hello edsellandicho$ more go.mod 
module examples.com/hello

go 1.20

replace examples.com/greetings => ../greetings
edsels-mbp:hello edsellandicho$ pwd
/Users/edsellandicho/goProjects/tutorials/examples/multipleModulesAndPackages/hello
edsels-mbp:hello edsellandicho$ go mod tidy
go: found examples.com/greetings in examples.com/greetings v0.0.0-00010101000000-000000000000
edsels-mbp:hello edsellandicho$ more go.mod 
module examples.com/hello

go 1.20

replace examples.com/greetings => ../greetings

require examples.com/greetings v0.0.0-00010101000000-000000000000
edsels-mbp:hello edsellandicho$ pwd
/Users/edsellandicho/goProjects/tutorials/examples/multipleModulesAndPackages/hello
edsels-mbp:hello edsellandicho$ go run .
Hi, Gladys. Welcome!
edsels-mbp:hello edsellandicho$ 