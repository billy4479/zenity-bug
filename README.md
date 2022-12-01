# Demonstrate a bug in [zenity](https://github.com/ncruces/zenity)

**How to reproduce**:

```sh
# 1. Run the program
go run main.go

# 2. Press on the cancel button

# 3. Observe the panic
```

**Expected behavior**: it should not panic on cancel

This also works using the [example on the docs](https://pkg.go.dev/github.com/ncruces/zenity#example-Progress).