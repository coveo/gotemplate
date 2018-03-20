{% include navigation.html %}
{% raw %}
# Comments in gotemplate

## Pseudo comment

If you insert gotemplate code into file that contains another kind of code such as hcl, json, yaml or xml, you code editor or linter may complains
because it will detect invalid characters.

To solve that problem, it is possible to inject pseudo comment into you code to hide the gotemplate code to your editor. The gotemplate is still interpretated, but obfuscated to the editor.

**It is important to render code that is valid for the host language.**

| Razor expression        | Go Template     | Note
| ----------------        | -----------     | ----
| `# @(2+2)`              | `{{ add 2 2 }}`   | Pseudo comment starting with #
| `// @(2+2)`             | `{{ add 2 2 }}`   | Pseudo comment starting with //
| `/*@ @(2+2) @*/`        | `{{ add 2 2 }}`   | Pseudo comment within /* */

### Example with JSON code

```go
/*@ $value := 2 + 8 * 15 @*/
{
    "Str": "string",
    "Int": 123,
    "Float": 1.23,
    "PiAsString": "@Math.Pi",
    "ComputedAsString": "@$value",

    /* You can use the special << syntax to extract the value from the string delimiter */
    "Pi": "<<@Math.Pi",
    "Computed": "<<@$value",
}
```

will give :

```go
{
    "Str": "string",
    "Int": 123,
    "Float": 1.23,
    "PiAsString": "3.141592653589793",
    "ComputedAsString": "122",

    /* You can use the special << syntax to extract the value from the string delimiter */
    "Pi": 3.141592653589793,
    "Computed": 122,
}
```

### Example with HCL code

```go
# @value := 2 + 8 * 15

Str              = "string"
Int              = 123
Float            = 1.23
PiAsString       = "@Math.Pi"
ComputedAsString = "@value"

// You can use the special << syntax to extract the value from the string delimiter
Pi       = "<<@Math.Pi"
Computed = "<<@value"
```

will give:
```go
Str              = "string"
Int              = 123
Float            = 1.23
PiAsString       = "3.141592653589793"
ComputedAsString = "122"

// You can use the special << syntax to extract the value from the string delimiter
Pi       = 3.141592653589793
Computed = 122
```
{% endraw %}