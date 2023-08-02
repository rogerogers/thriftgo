// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reflection_tpl

// UtilFile .
var UtilFile = `// Code generated by thriftgo ({{Version}}). DO NOT EDIT.
{{InsertionPoint "bof"}}

package {{.FilePackage}}

import (
	{{InsertionPoint "imports"}}
	{{define "Imports"}}
	{{end}}
	"github.com/cloudwego/thriftgo/thrift_reflection"
)

{{$IDLName := .IDLName}}
{{$IDLPath := .AST.Filename}}
{{$FilePackage := .FilePackage}}

{{- range .Structs}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetStructDescriptor("{{.Name}}")
}
{{- end}}
{{- range .Enums}}
func (p {{.GoName}}) GetDescriptor() *thrift_reflection.EnumDescriptor{
	return file_{{$IDLName}}_thrift.GetEnumDescriptor("{{.Name}}")
}
{{- end}}
{{- range .Unions}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetUnionDescriptor("{{.Name}}")
}
{{- end}}
{{- range .Exceptions}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetExceptionDescriptor("{{.Name}}")
}
{{- end}}
{{- InsertionPoint "eof"}}
`
