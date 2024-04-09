# HTML Link Parser

[Gophercises](https://gophercises.com/) Exercise Details:

In this exercise your goal is create a package that makes it easy to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each extracted link you should return a data structure that includes the `href`.

Links will be nested in different HTML elements, and it is very possible that you will have to deal with HTML similar to code below.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

In situations like these we want to get output that looks roughly like:

```go
Link{
  Href: "/dog",
}
```

Once you have a working program, try to write some tests for it to practice using the testing package in go.

<br>

## Technical Notes

- Use the `x/net/html` package. Package html implements an HTML5-compliant tokenizer and parser.
- Ignore nested links. Eg with following HTML:
    ```html
    <a href="#">
    Something here <a href="/dog">nested dog link</a>
    </a>
    ```
    It is okay if your code returns only the outside link - for the purposes of this exercise.
    <br>
    *Include the nested links as well in the output.*
- Test the code with example files included in the project repository. *Improve your tests and edge-case coverage.* Add Examples and Documentation for the code. Run the following in this order, using go tooling:
  - tests
    - go test
  - coverage
    - go test -cover
    - go test -coverprofile coverage.out
  - coverage shown in web browser
    - go tool cover -html=coverage.out
  - examples shown in documentation in a web browser
    - godoc -http=:8080