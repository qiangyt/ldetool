
package main

var staticTemplatesData = map[string]string{
"at_end": "\n// Check if the rest is empty\nif len(p.rest) > 0 {\nreturn false, {{ if .Serious }}fmt.Errorf(\"Rest must be empty, not `\\033[1m%s\\033[0m`\", string(p.rest)){{else}}nil{{end}}\n}\n",
"close_option_scope": ";p.{{.Name}} = true\n{{.PRest}} = {{.Rest}}\n{{ if .WasAbandoned }}{{.ScopeLabel}}:\n{{end}}\n",
"close_scope": "};\n",
"decode_float": "if tmpFloat, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&{{.Source}})), {{ .Bits }}); err != nil {\n    return false, fmt.Errorf(\"Error parsing \\033[1m%s\\033[0m value as {{ .Type }} for field \\033[1m{{ .Dest }}\\033[0m: %s\", string({{.Source}}), err)\n}\np.{{.Dest}} = float{{.Bits}}(tmpFloat);\n",
"decode_int": "if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&{{ .Source }})), 10, {{ .Bits }}); err != nil {\n    return false, fmt.Errorf(\"Error parsing \\033[1m%s\\033[0m value as {{ .Type }} for field \\033[1m{{ .Dest }}\\033[0m: %s\", string({{ .Source }}), err)\n}\np.{{.Dest}} = int{{.Bits}}(tmpInt);\n",
"decode_string": "p.{{.Dest}} = {{.Source}};\n",
"decode_uint": "if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&{{.Source}})), 10, {{ .Bits }}); err != nil {\n    return false, fmt.Errorf(\"Error parsing \\033[1m%s\\033[0m value as {{ .Type }} for field \\033[1m{{ .Dest }}\\033[0m: %s\", string({{.Source}}), err)\n}\np.{{.Dest}} = uint{{.Bits}}(tmpUint);\n",
"getter": "\n// Get{{.LongName}} retrieves optional value for {{.LongName}}.Name\nfunc (p *{{.ParserName}}) Get{{.LongName}}() (res {{.Type}}) {\n\t if {{ range $index, $data := .Accesses }}{{ if $index }}||{{end}}!p.{{$data}}.Valid{{end}} {\n  \t return\n\t };\n\t return p.{{.Access}}.{{.Name}}\n}\n",
"head_char": "\n// Checks if the rest starts with {{.Char}} symbol and pass it\nif len({{.Rest}}) > 0 && {{.Rest}}[0] == {{.Char}} {\n   {{.Rest}} = {{.Rest}}[1:]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"The rest (\\033[1m%s\\033[0m) doesn't start with `\\033[1m%s\\033[0m`\", string({{.Rest}}), {{.Char}}){{else}}nil{{end}};{{end}}\n}\n",
"head_char_maybe": "\n// Checks if the rest starts with {{.Char}} symbol and pass it if matched\nif len({{.Rest}}) > 0 && {{.Rest}}[0] == {{.Char}} {\n   {{.Rest}} = {{.Rest}}[1:]\n}\n",
"head_string": "\n// Checks if the rest starts with {{.ConstValue}} and pass it\nif bytes.HasPrefix({{.Rest}}, {{.ConstName}}) {\n   {{.Rest}} = {{.Rest}}[len({{.ConstName}}):]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"The rest (\\033[1m%s\\033[0m) doesn't start with `\\033[1m%s\\033[0m`\", string({{.Rest}}), {{.ConstName}}){{else}}nil{{end}};{{end}}\n}\n",
"head_string_maybe": "\n// Checks if the rest starts with {{.ConstValue}} and pass it if matched\nif bytes.HasPrefix({{.Rest}}, {{.ConstName}}) {\n   {{.Rest}} = {{.Rest}}[len({{.ConstName}}):]\n}\n",
"lookup_bounded_char": "\n// Looking for {{ .Char }} symbol from {{.Lower}} until {{.Upper}} symbols of the rest and then pass it\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in {{.Lower}}..%d-th symbols of the rest while only %s left\", {{.Upper}}-1, len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\npos = bytes.IndexByte({{.Rest}}[{{.Lower}}:{{.Upper}}], {{ .Char }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + {{.Lower}} + 1:]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%c\\033[0m` in `\\033[1m%s\\033[0m`\", {{.Char}}, string({{.Rest}}[{{.Lower}}:{{.Upper}}])){{else}}nil{{end}};{{end}}\n}\n",
"lookup_bounded_char_noerror": "\n// Looking for {{ .Char }} symbol from {{.Lower}} until {{.Upper}} symbols of the rest and then pass it\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in {{.Lower}}..%d-th symbols of the rest while only %s left\", {{.Upper}}-1, len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\npos = bytes.IndexByte({{.Rest}}[{{.Lower}}:{{.Upper}}], {{ .Char }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + {{.Lower}} + 1:]\n}\n",
"lookup_bounded_string": "\n// Looking for {{ .ConstValue }} from {{.Lower}} until {{.Upper}} symbols of the rest and then pass it\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in {{.Lower}}..%d-th symbols of the rest while only %d left\", {{.Upper}}, len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\npos = bytes.Index({{.Rest}}[{{.Lower}}:{{.Upper}}], {{ .ConstName }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + {{.Lower}} + len({{.ConstName}}):]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%s\\033[0m` in `\\033[1m%s\\033[0m`\", {{.ConstName}}, string({{.Rest}}[{{.Lower}}:{{.Upper}}])){{else}}nil{{end}};{{end}}\n}\n",
"lookup_bounded_string_noerror": "\n// Looking for {{ .ConstValue }} from {{.Lower}} until {{.Upper}} symbols of the rest and then pass it if found\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in {{.Lower}}..%d-th symbols of the rest while only %d left\", {{.Upper}}-1, len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\npos = bytes.Index({{.Rest}}[{{.Lower}}:{{.Upper}}], {{ .ConstName }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + {{.Lower}} + len({{.ConstName}}):]\n}\n",
"lookup_char": "\n// Looking for {{ .Char }} symbol and then pass it\npos = bytes.IndexByte({{.Rest}}, {{ .Char }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + 1:]\n} else {\n{{ if .Namespace }}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%c\\033[0m` in `\\033[1m%s\\033[0m`\", {{.Char}}, string({{.Rest}})){{else}}nil{{end}};{{end}}\n}\n",
"lookup_char_noerror": "\n// Looking for {{ .Char }} symbol and then pass it if found\npos = bytes.IndexByte({{.Rest}}, {{ .Char }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + 1:]\n}\n",
"lookup_limited_char": "\n// Looking for {{ .Char }} symbol in first {{.Upper}} symbols of the rest and then pass it\nif len({{.Rest}}) < {{.Upper}} {\n{{ if .Namespace }}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\npos = bytes.IndexByte({{.Rest}}[:{{.Upper}}], {{ .Char }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + 1:]\n} else {\n{{ if .Namespace }}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%c\\033[0m` in `\\033[1m%s\\033[0m`\", {{.Char}}, string({{.Rest}}[:{{.Upper}}])){{else}}nil{{end}};{{end}}\n}\n",
"lookup_limited_char_noerror": "\n// Looking for {{ .Char }} symbol in first {{.Upper}} symbols of the rest and then pass it if found\nif len({{.Rest}}) < {{.Upper}} {\n{{ if .Namespace }}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols of the rest while only %s left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\npos = bytes.IndexByte({{.Rest}}[:{{.Upper}}], {{ .Char }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + 1:]\n}\n",
"lookup_limited_string": "\n// Looking for {{ .ConstValue }} in first {{.Upper}} symbols of the rest and then pass it\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\npos = bytes.Index({{.Rest}}[:{{.Upper}}], {{ .ConstName }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + len({{.ConstName}}):]\n} else {\n{{if .Namespace }}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%s\\033[0m` in `\\033[1m%s\\033[0m`\", {{.ConstName}}, string({{.Rest}}[:{{.Upper}}])){{else}}nil{{end}};{{end}}\n}\n",
"lookup_limited_string_noerror": "\n// Looking for {{ .ConstValue }} in first {{.Upper}} symbols of the rest and then pass it if found\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols of the rest while only %s left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\npos = bytes.Index({{.Rest}}[:{{.Upper}}], {{ .ConstName }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + len({{.ConstName}}):]\n}\n",
"lookup_string": "\n// Looking for {{ .ConstValue }} and then pass it\npos = bytes.Index({{.Rest}}, {{ .ConstName }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + len({{.ConstName}}):]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%s\\033[0m` in `\\033[1m%s\\033[0m`\", {{.ConstName}}, string({{.Rest}})){{else}}nil{{end}};{{end}}\n}\n",
"lookup_string_noerror": "\n// Looking for {{ .ConstValue }} and then pass it if found\npos = bytes.Index({{.Rest}}, {{ .ConstName }});\nif pos >= 0 {\n    {{.Rest}} = {{.Rest}}[pos + len({{.ConstName}}):]\n}\n",
"open_option": "{{.Name}} struct {\n   Valid bool;\n",
"open_option_scope": ";{{.Rest}} = {{.PRest}};\n",
"parser_body": "// Parse autogenerated method of {{.ParserName}}\nfunc (p *{{.ParserName}}) Parse(line []byte) (bool, error) {\n   {{ range .Vars }} var {{.Name}} {{.Type}};{{end}}\n   p.rest = line\n   {{.Parser}}\n   return true, nil\n}\n",
"parser_code": "package {{ .PkgName }}\n\n{{ if .Imports }}import (\n{{ range .Imports }}{{.Name}} \"{{.Path}}\"\n{{end}}\n){{end}}\n\n{{range $name, $value := .Consts }}var {{$name}} = []byte({{$value}});{{end}}\n\n\n{{.Struct}}\n\n{{.Parser}}\n\n{{.Getters}}\n",
"pass_n_items": "\n// Cut out first N items from the rest\nif len({{.Rest}}) >= {{.Upper}} {\n {{.Rest}} = {{.Rest}}[{{.Upper}}:]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot cut out first {{.Upper}} symbols from the rest: it is shorter (%d)\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\n",
"soft_exit": "return false, nil\n",
"struct_body": "// {{.ParserName}} autogenerated parser\ntype {{.ParserName}} struct {\nrest []byte\n{{.Struct}}\n}\n",
"struct_field": "{{.Name}} {{.Type}};\n",
"take_before_bounded_char": "\n// Put data before {{ .Char }} into {{ .Dest }} with limited to {{.Lower}}..{{.Upper}} symbol range boundary lookup\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols to bound value for {{.Dest}} of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\nif pos = bytes.IndexByte({{.Rest}}[{{.Lower}}:{{.Upper}}], {{ .Char }}); pos >= 0 {\n   {{ if .UseTmp }}tmp = {{.Rest}}[:{{.Lower}}+pos]\n     {{ call .Decoder \"tmp\"  .Dest }}{{else}} {{ call .Decoder ( printf \"%s[:%d+pos]\" .Rest .Lower )  .Dest }}{{end}}\n  {{.Rest}} = {{.Rest}}[pos+{{.Lower}}+1:]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%c\\033[0m` in `\\033[1m%s\\033[0m` to bound data for field {{.Dest}}\", {{.Char}}, string({{.Rest}}[{{.Lower}}:{{.Upper}}])){{else}}nil{{end}};{{end}}\n}\n",
"take_before_bounded_char_or_rest": "\n// Put data before {{ .Char }} into {{ .Dest }} with limited to {{.Lower}}..{{.Upper}} symbols boundary lookup or everything to the very rest if not found\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols to bound value for {{.Dest}} of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\nif pos = bytes.IndexByte({{.Rest}}[{{.Lower}}:{{.Upper}}], {{ .Char }}); pos >= 0 {\n  tmp = {{.Rest}}[:pos+{{.Lower}}]\n  {{.Rest}} = {{.Rest}}[pos+{{.Lower}}+1:]\n} else {\n  tmp = {{.Rest}}\n  {{.Rest}} = {{.Rest}}[len({{.Rest}}):]\n}\n{{ call .Decoder \"tmp\"  .Dest }}\n",
"take_before_bounded_string": "\n// Put data before {{ .ConstValue }} into {{ .Dest }} with limited to {{.Lower}}..{{.Upper}} symbol range boundary lookup\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols to bound value for {{.Dest}} of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\nif pos = bytes.Index({{.Rest}}[{{.Lower}}:{{.Upper}}], {{ .ConstName }}); pos >= 0 {\n   {{ if .UseTmp }}tmp = {{.Rest}}[:{{.Lower}}+pos]\n     {{ call .Decoder \"tmp\"  .Dest }}{{else}} {{ call .Decoder ( printf \"%s[:%d+pos]\" .Rest .Lower )  .Dest }}{{end}}\n  {{.Rest}} = {{.Rest}}[pos+{{.Lower}}+len({{.ConstName}}):]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%s\\033[0m` in `\\033[1m%s\\033[0m` to bound data for field {{.Dest}}\", {{.ConstName}}, string({{.Rest}}[{{.Lower}}:{{.Upper}}])){{else}}nil{{end}};{{end}}\n}\n",
"take_before_bounded_string_or_rest": "\n// Put data before {{ .ConstValue }} into {{ .Dest }} with limited to {{.Lower}}..{{.Upper}} symbols boundary lookup or everything to the very rest if not found\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols to bound value for {{.Dest}} of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\nif pos = bytes.Index({{.Rest}}[{{.Lower}}:{{.Upper}}], {{ .ConstName }}); pos >= 0 {\n  tmp = {{.Rest}}[:pos+{{.Lower}}]\n  {{.Rest}} = {{.Rest}}[pos+{{.Lower}}+len({{.ConstName}}):]\n} else {\n  tmp = {{.Rest}}\n  {{.Rest}} = {{.Rest}}[len({{.Rest}}):]\n}\n{{ call .Decoder \"tmp\"  .Dest }}\n",
"take_before_char": "\n// Put data before {{ .Char }} into {{ .Dest }}\nif pos = bytes.IndexByte({{.Rest}}, {{ .Char }}); pos >= 0 {\n  {{ if .UseTmp }}tmp = {{.Rest}}[:pos]\n  {{ call .Decoder \"tmp\"  .Dest }}{{else}}{{ call .Decoder ( printf \"%s[:pos]\" .Rest )  .Dest }}{{end}}\n  {{.Rest}} = {{.Rest}}[pos+1:]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%c\\033[0m` in `\\033[1m%s\\033[0m` to bound data for field {{.Dest}}\", {{.Char}}, string({{.Rest}})){{else}}nil{{end}};{{end}}\n}\n",
"take_before_char_or_rest": "\n// Put data before {{ .Char }} into {{ .Dest }} if found otherwise take to the end\nif pos = bytes.IndexByte({{.Rest}}, {{ .Char }}); pos >= 0 {\n  tmp = {{.Rest}}[:pos]\n  {{.Rest}} = {{.Rest}}[pos+1:]\n} else {\n  tmp = {{.Rest}}\n  {{.Rest}} = {{.Rest}}[len({{.Rest}}):]\n}\n{{ call .Decoder \"tmp\"  .Dest }}\n",
"take_before_limited_char": "\n// Put data before {{ .Char }} into {{ .Dest }} with limited to {{.Upper}} symbols boundary lookup\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols to bound value for {{.Dest}} of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\nif pos = bytes.IndexByte({{.Rest}}[:{{.Upper}}], {{ .Char }}); pos >= 0 {\n  {{ if .UseTmp }}tmp = {{.Rest}}[:pos]\n  {{ call .Decoder \"tmp\"  .Dest }}{{else}}{{ call .Decoder ( printf \"%s[:pos]\" .Rest )  .Dest }}{{end}}\n  {{.Rest}} = {{.Rest}}[pos+1:]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%c\\033[0m` in `\\033[1m%s\\033[0m` to bound data for field {{.Dest}}\", {{.Char}}, string({{.Rest}}[:{{.Upper}}])){{else}}nil{{end}};{{end}}\n}\n",
"take_before_limited_char_or_rest": "\n// Put data before {{ .Char }} into {{ .Dest }} with limited to {{.Upper}} symbols boundary lookup or everything to the very rest if not found\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols to bound value for {{.Dest}} of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\nif pos = bytes.IndexByte({{.Rest}}[:{{.Upper}}], {{ .Char }}); pos >= 0 {\n  tmp = {{.Rest}}[:pos]\n  {{.Rest}} = {{.Rest}}[pos+1:]\n} else {\n  tmp = {{.Rest}}\n  {{.Rest}} = {{.Rest}}[len({{.Rest}}):]\n}\n{{ call .Decoder \"tmp\"  .Dest }}\n",
"take_before_limited_string": "\n// Put data before {{ .ConstValue }} into {{ .Dest }} with limited to {{.Upper}} symbols boundary lookup\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols to bound value for {{.Dest}} of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\nif pos = bytes.Index({{.Rest}}[:{{.Upper}}], {{ .ConstName }}); pos >= 0 {\n  {{ if .UseTmp }}tmp = {{.Rest}}[:pos]\n  {{ call .Decoder \"tmp\"  .Dest }}{{else}}{{ call .Decoder ( printf \"%s[:pos]\" .Rest )  .Dest }}{{end}}\n  {{.Rest}} = {{.Rest}}[pos+len({{.ConstName}}):]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%s\\033[0m` in `\\033[1m%s\\033[0m` to bound data for field {{.Dest}}\", {{.ConstName}}, string({{.Rest}}[:{{.Upper}}])){{else}}nil{{end}};{{end}}\n}\n",
"take_before_limited_string_or_rest": "\n// Put data before {{ .ConstValue }} into {{ .Dest }} with limited to {{.Upper}} symbols boundary lookup or everything to the very rest if not found\nif len({{.Rest}}) < {{.Upper}} {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if.Serious }}fmt.Errorf(\"Requested lookup in first {{.Upper}} symbols to bound value for {{.Dest}} of the rest while only %d left\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\nif pos = bytes.Index({{.Rest}}[:{{.Upper}}], {{ .ConstName }}); pos >= 0 {\n  tmp = {{.Rest}}[:pos]\n  {{.Rest}} = {{.Rest}}[pos+len({{.ConstName}}):]\n} else {\n  tmp = {{.Rest}}\n  {{.Rest}} = {{.Rest}}[len({{.Rest}}):]\n}\n{{ call .Decoder \"tmp\"  .Dest }}\n",
"take_before_string": "\n// Put data before {{ .ConstValue }} into {{ .Dest }}\nif pos = bytes.Index({{.Rest}}, {{ .ConstName }}); pos >= 0 {\n  {{ if .UseTmp }}tmp = {{.Rest}}[:pos]\n  {{ call .Decoder \"tmp\"  .Dest }}{{else}}{{ call .Decoder ( printf \"%s[:pos]\" .Rest )  .Dest }}{{end}}\n  {{.Rest}} = {{.Rest}}[pos+len({{.ConstName}}):]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot find `\\033[1m%s\\033[0m` in `\\033[1m%s\\033[0m` to bound data for field {{.Dest}}\", {{.ConstName}}, string({{.Rest}})){{else}}nil{{end}};{{end}}\n}\n",
"take_before_string_or_rest": "\n// Put data before {{ .ConstValue }} into {{ .Dest }} if found otherwise take to the end\nif pos = bytes.Index({{.Rest}}, {{ .ConstName }}); pos >= 0 {\n  tmp = {{.Rest}}[:pos]\n  {{.Rest}} = {{.Rest}}[pos+len({{.ConstName}}):]\n} else {\n  tmp = {{.Rest}}\n  {{.Rest}} = {{.Rest}}[len({{.Rest}}):]\n}\n{{ call .Decoder \"tmp\"  .Dest }}\n",
"take_rest": "\n// Put the rest of the data into {{.Dest}}\n{{ call .Decoder (printf \"%s\" .Rest) .Dest }}\n{{.Rest}} = {{.Rest}}[len({{.Rest}}):]\n",

}

