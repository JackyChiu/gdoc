## Install
```bash
go get github.com/JackyChiu/gdoc
```

## Use
```bash
# example project
foo
├── bar
│   └── bar.go
└── foo.go

# view current package
gdoc

# view subpackage 
gdoc bar

# view std packages
gdoc os/exec

# view any package path
gdoc ../../anotherUser/anotherPkg
```
