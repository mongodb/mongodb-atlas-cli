# This is a work in progress

# Code style

Make sure your code pases the gofmt and golint rules

```bash
make fmt lint
```

# Testing

Just do it

```bash
make test
```

## To update mocks

```bash
mockgen -source=internal/store/projects.go -destination=mocks/mock_projects.go -package=mocks
```
