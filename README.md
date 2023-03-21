# Parsley JSON
Parsley JSON is a JSON mapper that parses a JSON string into an object or array. Parsley JSON uses the definition of an struct to generate code for reading values efficiently. This library is inspired by [easyjson](https://github.com/mailru/easyjson).

## Installation 
```
go install github.com/Soreing/parsley/parsley@latest
go get github.com/Soreing/parsley
```

## Usage
Objects that implement the ParsleyJSONUnmarshaller interface can be used in the Unmarshal function with a byte array to parse. The interface can be implemented with the generator (or custom written for rare use cases).
```golang
dat := []byte (`{"name": "A Box", "weight": 10.5}`)
box := Box{}

err := parsley.Unmarshal(dat, &box)
if err != nil {
	panic(err)
}
```

## Code Generation
To read data into objects, first you need to generate the mapping functions for the structs you want to use. The explicit flag tells the generator to process the defined struct.
```golang
//parsley:explicit
type Box struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

//parsley:explicit
type BoxList []Box
```
Run the generator with the path to the file or folder to the module. Assuming the above file is in tests/models/box.go, use the following command:
```
parsley tests/models/box.go
```
You can alter the generator's behavior using command line argument flags. When `-all` is used, you can add `//parsley:skip` above definitions to force the generator to skip them.
| Flag | Description |
|------|-------------|
| -all | Processes all structs in a file or directory |
| -output_filename | Sets the filename of the generated code |
| -camel_case | Uses camel case naming for fields unless explicitly defined |
| -lower_case | Uses lower case naming for fields unless explicitly defined |
| -kebab_case | Uses kebab case naming for fields unless explicitly defined |
| -snake_case | Uses snake case naming for fields unless explicitly defined |
| -pascal_case | Uses pascal case naming for fields unless explicitly defined |

## Todo List
- Implement reading of scientific (exponential) notation for number fields
- Improve speed of floating point number creation
- Implement marshalling functions