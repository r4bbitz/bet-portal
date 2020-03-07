

### Install

```
go get -u github.com/r4bbitz/SayHello
```

### build
```
go build
```
### Run SayHello

```
Run SayHello.exe
```

### Install Newman  Automation test Cli
```
npm i newman
```


### Automation Test
```
newman run   automationtest/SayHelloAutomationTest.postman_collection.json -d automationtest/input.json
```
###URL
```
GET http://localhost:1991/api/v1/sayHello
```
###Expect Result test
```
automationtest/input.json
```
