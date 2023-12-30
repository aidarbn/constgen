# constgen

Generates golang constants out of the file.

Install:

```
go install github.com/alexeykhan/constgen@latest
```

Command:

```
constgen -input path/to/your/file.txt -output path/to/output/file.go -package YourPackageName
```

Example command:

```
constgen -input example/example.txt -output example/output/output.gen.go -package output
```