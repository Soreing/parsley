# Parsley JSON
Parsley JSON is a JSON decoder that parses a JSON string into objects. Parsley JSON uses the definition of a struct to generate code that is efficient and fast. This library is inspired by [easyjson](https://github.com/mailru/easyjson).

## Installation 
```
go install github.com/Soreing/parsley/parsley@latest
go get github.com/Soreing/parsley
```

## Usage
Objects that implement the ParsleyJSON interfaces can be used in the Unmarshal functions to encode or decode data. The interface can be implemented with the generator (or custom written for rare use cases).
```golang
dat := []byte (`{"name": "A Box", "weight": 10.5}`)
box := Box{}

err := parsley.Unmarshal(dat, &box)
if err != nil {
	panic(err)
}
```

## Code Generation
To encode/decode data, first you need to generate the mapping functions for the structs you want to use. The json flag tells the generator to process the defined struct. You can add more flags separated by commas.
```golang
//parsley:json
type Box struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

//parsley:json
type BoxList []Box
```
Run the generator with the path to the file or folder to the module. Assuming the above file is in tests/models/box.go, use the following command:
```
parsley tests/models/box.go
```
### Struct Flags
| Flag | Description |
|------|-------------|
| json   | Process the struct or type define in the generator |
| skip   | Skip the struct or type define during generation |
| public | Include private fields in encoding/decoding |

### Command Line Arguments
| Arg | Description |
|------|-------------|
| -all | Processes all structs in a file or directory |
| -public | Include private fields in encoding/decoding for all types |
| -output_filename | Sets the filename of the generated code |
| -camel_case | Uses camel case naming for fields unless explicitly defined |
| -lower_case | Uses lower case naming for fields unless explicitly defined |
| -kebab_case | Uses kebab case naming for fields unless explicitly defined |
| -snake_case | Uses snake case naming for fields unless explicitly defined |
| -pascal_case | Uses pascal case naming for fields unless explicitly defined |

## Todo List
- Implement reading of scientific (exponential) notation for number fields
- Improve speed of floating point number creation
