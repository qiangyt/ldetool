# ldetool means line data extraction tool
`ldetool` is a command line utility to generate Go code for parsing log files. 

```bash
go get -u github.com/sirkon/ldetool
```

### Preamble

There's a traditional solution for this kind of tasks: regular expression with capture groups. But it has numerous generic and Go-specific disadvantages:

1. Regexes are hard to read and debug.
2. Speed. While simple non-capturing regular expressions can be speedy, they quickly becomes slow as the complexity of the regular expression grows
3. They are overpowered for simple log parsing. In our experience with log processing we are not looking for patterns within the line. Usually our data is well structured and it is easier to think (and compute!) in terms of bounds and separators. And if the data is not well structured then it is a good idea to make it so, just for the sake of readability.
4. Go regular expressions are slow. Go regular expressions with group capture are even slower. Simple comparison on performance [here](PERFORMANCE.md)
5. There are no named captured groups in Go regexes, we must to use arrays instead which is hard for reading and comprehension.
6. Named capture groups are not supported in Go regexes.

There is another traditional approach: manual data extraction. We manually command to find a symbol or substring and pass
it or take everything before it and put into variable, it also has his share of generic disadvantages:

1. It is annoying as hell to write it
2. It can be hard to read

Still, the major advantage is:
1. It can be fast

We had severe shortage of resources at my last job, we couldn't just buy some more power, so we had no choice. We had to write it manually.
It turned out most of things to retrieve data are repetitive and we are writing nearly the same things again and again.

##### Typical operations:
1. Check if the rest starts with the certain string or character and pass it
1. Look for the char or substring in the rest and pass it
    ```
    0x23a719bdf5589bc.Receive() -> ID[alissa] Country[RU] … 
    ```
    We obviously only need data after `-> ` in this case, so we just need to find `-> ` and pass it. We also may have
    some apriori knowledge of the length of chunk where the `-> ` might be found, thus the next command.
2. Look for the char or substring in the first N characters of the rest and pass it
3. Take all data from the rest up to the certain string or character and save it under some name.
4. etc

So, we wrote a code generator for this purpose. The code turned to be even faster than one we used to write, since we actually
were trying to reduce amount of code we are writing introducing helper abstractions what might have some cost, the generator just put raw code.

### How it works.
1. Write extraction script.
2. Generate go code using extraction script.
3. Use it via the generated extraction method

#### Example

Take a look at these two lines

```
[2017-09-02T22:48:13] FETCH first[1] format[JSON] hidden[0] userAgent[Android App v1.0] rnd[21341975] country[MA]
[2017-09-02T22:48:14] FETCH first[0] format[JSON] userAgent[Android App v1.0] rnd[10000000] country[LC]
```

We likely need a time, value of parameter `first`, `format`, `hidden`, `userAgent` and `country`. We obviously don't need `rnd` 

##### Extraction script syntax
See [more details](TOOL_RULES.md) on extraction rules

```perl
# filename: line.lde
Line =                                   # Name of the extraction object' type
  ^'['                                   # The line must start with [
  Time(string) ']'                       # Take everything as a string for Time right to ']' character
  ^" FETCH "                             # Current rest must starts with " FETCH " string
  ^"first[" First(uint8) ']'             # The rest must starts with "first[" characters, then take the rest until ']' as uint8
                                         # under the name of First
  ^" format[" Format(string) ']'         # Take format id
  ?Hidden (^" hidden[" Value(uint8) ']') # Optionally look for " hidden[\d+]"
  ^" user_agent[" UserAgent(string) ']'  # User agent data
  _ "country[" Country(string)  ']'      # Look for the piece starting with country[
;
```

##### Code generation
The easiest way is to put `//go:generate ldetool generate --package main Line.lde` somewhere in the Go file and then generate a code with 
```bash
go generate <project path>
```
It will be written into `line_lde.go` file in the same directory. It will look like [this](SAMPLE.md)

Now, we have
1. Data extractor type
    ```go
    // Line autogenerated parser
    type Line struct {
        rest   []byte
        Time   []byte
        First  uint8
        Format []byte
        Hidden struct {
            Valid bool
            Value uint8
        }
        UserAgent []byte
        Country   []byte
    }
    ```
2. Parse method
    ```go
    // Extract autogenerated method of Line
    func (p *Line) Extract(line []byte) (bool, error) {
       …
    }
    ```
    Take a look at return data. First bool signals if the data was successfully matched and error signals if there were
    any error. String to numeric failures are always treated as errors, you can put `!` into extraction script and all
    mismatches after the sign will be treated as errors
3. Helper to access optional `Hidden` area returning default Go value if the the area was not matched
    ```go
    // GetHiddenValue retrieves optional value for HiddenValue.Name
    func (p *Line) GetHiddenValue() (res uint8) {
        if !p.Hidden.Valid {
            return
        }
        return p.Hidden.Value
        …
    }    
    ```
    
##### Generated code usage
It is easy: put
```go
l := &Line{}
```
before and then feed `Parse` method with lines:
```go
scanner := bufio.NewScanner(reader)
for scanner.Scan() {
    ok, err := l.Extract(scanner.Bytes())
    if !ok {
        if err != nil {
            return err
        }
        continue
    }
    …
    l.Format
    l.Time
    l.GetHiddenValue()
    …
}
```

