# Problem 10

Create a function that will return:

1. map of brands mapped to their amount 
2. function to create a brand, which will return incrementer of numbers of brands item.

```go
brands, makeBrand := factory()
toyotaIncrementer := makeBrand("Toyota")
toyotaIncrementer(1)
toyotaIncrementer(2)
hyundaiIncrementer := makeBrand("Hyundai")
hyundaiIncrementer(5)
_ := makeBrand("Kia")
fmt.Println(brands) // map[Hyundai:5 Kia:0 Toyota:3]
```
