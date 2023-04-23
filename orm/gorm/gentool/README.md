# gen

gorm example with [gorm.io/gen/tools/gentool](https://github.com/go-gorm/gen/tree/master/tools/gentool)

## prepare

```console
go install gorm.io/gen/tools/gentool@latest
```

## Problems

- gentool does not generate relations
  - it require [manual configurations](https://github.com/go-gorm/gen/issues/607#issuecomment-1427671010)
