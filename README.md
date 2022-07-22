`mdk` is a tool which helps you extract kubectl yaml manifests from your Markdown files.  

## Usage
```bash
go build -o mdk
```
Tested with `go1.17.7 linux/amd64` but should work with other versions (create an issue if it doesn't)

Add `+kubectl` annotation to your yaml codeblock in the markdown file. 
For example:
````md
## Hello world
Some text here.
```yaml
# +kubectl
# this is a comment
foo:
    bar: 2
```

Some other text here
```yaml
# this is another comment
baz:
    say: "hello world"
```
````
From the root directory, 
```bash
mdk <above-file-path.yaml> 
```
Output:
```
# +kubectl
# this is another comment
ooh:
    lala: 2
    la: "hello world"

---
```
You can pair it with `kubectl` like this:
```
mdk <md-file.yaml> | kubectl apply -
```
You can write the yaml codeblock to a file like this:
```
mdk <md-file.yaml> > code.yaml
```

## What is considered as a target for kubectl?
Anything with `+kubectl` annotation in code is a target.
### Examples that work
```yaml
# +kubectl
foo: bar
```
```yaml

# +kubectl
foo: bar
```

### Examples that don't work
```yaml
# +kubectlfoo
foo: bar
```

Example that *doesn't* work:
```yaml
foo: bar
# +kubectl
```

## Build
```bash
go build -o mdk
```
Tested with `go1.17.7 linux/amd64` but should work with other versions (create an issue if it doesn't)

