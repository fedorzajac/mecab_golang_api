https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry


# steps to reproduce

```bash
go build .
docker build -t mecab_golang_api .
docker images
docker tag c01a8ac6bf8a ghcr.io/fedorzajac/mecab_golang_api:latest
docker push ghcr.io/fedorzajac/mecab_golang_api:latest
```

then, on vm

```bash
docker pull ghcr.io/fedorzajac/mecab_golang_api:latest
docker run -d -p 8080:8080 ghcr.io/...
```

# Reading file in golang

reading file in golang and return lines

```golang
func rf(filepath string) []string {
	out := []string{}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return out
}
```

and json -> returns a map with keys and values

```go
// json:
// { "a" : {"k" : "", ...}, "b": {...}}


type entry struct {
	K string `json:"k"`
	R string `json:"r"`
	E string `json:"e"`
}

func read_categories_from_json(str string) map[string]entry {
	jsonFile, err := os.Open(str)
	// if we os.Open returns an error then handle it
	if err != nil {
			fmt.Println(err)
	}
	// fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Users array
	var temp map[string]entry
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &temp)
	return temp
}
```

check if key exists in map

```go
if val, ok := N1[c]; ok {
			print(val.K, " is N1 \n")
			continue
		} else {
      // not exists
    }
```
