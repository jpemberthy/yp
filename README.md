# yp
Go YAML parser sample

```
go get gopkg.in/yaml.v2
go get github.com/kr/pretty
go run main.go
```

The output looks like

```
map[string]map[string]main.Package{
    "1.0.1": {
        "jquery-1.0.1.pack.js": {Hash:"08341cd159e29f561ca0ec16c99bf4b85e43d30f", Normalized:"9e326989381601ca91f7a1c626367f0e594d80fd"},
        "jquery-1.0.1.js":      {Hash:"6c9dcc0685e56377016a0b50e008a42ef84c04a4", Normalized:"ac2fbe63d02798c62d447441c0dae6c5fd1b1043"},
    },
    "1.0.2": {
        "jquery-1.0.2.js":      {Hash:"d3b30b0fead39e4c40fb0c91408e74439020a279", Normalized:"1d32919fa41b4bff1f1ecb27516af5b709221b5e"},
        "jquery-1.0.2.pack.js": {Hash:"84b1514a01def3bc0b52f6fa03d0d9fa349bef72", Normalized:"bab5966fc200dff4ad4b02700821879d2a75be94"},
    },
    "1.0.3": {
        "jquery-1.0.3.js":      {Hash:"4c12e01d990bd2b1075812d9f28e3ffa50ca59df", Normalized:"3f8f77f3745c4711cb587b857001d95680a40566"},
        "jquery-1.0.3.pack.js": {Hash:"ddb4126bf4713cb4e0f2310401e58cb9e3f98997", Normalized:"3252e87185128ff7ba00cb88f8c2cbeda6452c56"},
    },
}
```
