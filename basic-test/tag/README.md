# testing with tag

go สามารถใส่ tag เพื่อให้รัน test เฉพาะที่ต้องการได้หรือใช้แยกประเภทของ test

```go
go test -v -tags integration ...

```

## การใช้ tag OR ในการรัน tag อื่นด้วยคือ

```go
// go:build integration || db
```

## การใช้ tag AND ในการรัน tag อื่นด้วยคือ

```go
// go:build integration && db
```

## การใช้ tag NOT ในการรัน tag อื่นด้วยคือ

```go
// go:build !integration
```
