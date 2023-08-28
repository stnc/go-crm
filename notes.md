Hayvan Kesim Yerleri için Hisseli Kurban Takip Prgramı 
```
learnGoWithTests
    |
    |-> helloworld
    |    |- hello.go
    |    |- hello_test.go    
    |
    |-> integers
    |    |- adder_test.go
    |
    |- go.mod
    |- README.md

```


```
$ go mod init example
go: creating new go.mod: module example
$ mkdir -p external
$ touch external/{external.go,external_test.go}
$ tree .
.
├── external
│   ├── external.go
│   └── external_test.go
└── go.mod

1 directory, 3 files
```
