Wanneer er een nieuwere versie van `alpinods` gebruikt wordt
moeten waarschijnlijk deze regels in `v2/alpino.go` aangepast worden:

```go
    if v1 < 1 || (v1 == 1 && v2 < 17) {
        alpino.Version = "1.17"
    }
```
