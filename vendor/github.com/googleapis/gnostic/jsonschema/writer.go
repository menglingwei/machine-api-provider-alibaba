<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Copyright 2017 Google LLC. All Rights Reserved.
=======
// Copyright 2017 Google Inc. All Rights Reserved.
>>>>>>> 79bfea2d (update vendor)
=======
// Copyright 2017 Google LLC. All Rights Reserved.
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
// Copyright 2017 Google LLC. All Rights Reserved.
>>>>>>> 03397665 (update api)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jsonschema

import (
	"fmt"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD

	"gopkg.in/yaml.v3"
=======
	"gopkg.in/yaml.v2"
>>>>>>> 79bfea2d (update vendor)
=======

	"gopkg.in/yaml.v3"
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======

	"gopkg.in/yaml.v3"
>>>>>>> 03397665 (update api)
)

const indentation = "  "

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 03397665 (update api)
func renderMappingNode(node *yaml.Node, indent string) (result string) {
	result = "{\n"
	innerIndent := indent + indentation
	for i := 0; i < len(node.Content); i += 2 {
		// first print the key
		key := node.Content[i].Value
		result += fmt.Sprintf("%s\"%+v\": ", innerIndent, key)
		// then the value
		value := node.Content[i+1]
		switch value.Kind {
		case yaml.ScalarNode:
			result += "\"" + value.Value + "\""
		case yaml.MappingNode:
			result += renderMappingNode(value, innerIndent)
		case yaml.SequenceNode:
			result += renderSequenceNode(value, innerIndent)
		default:
			result += fmt.Sprintf("???MapItem(Key:%+v, Value:%T)", value, value)
		}
		if i < len(node.Content)-2 {
			result += ","
		}
		result += "\n"
<<<<<<< HEAD
=======
func renderMap(info interface{}, indent string) (result string) {
=======
func renderMappingNode(node *yaml.Node, indent string) (result string) {
>>>>>>> e879a141 (alibabacloud machine-api provider)
	result = "{\n"
	innerIndent := indent + indentation
	for i := 0; i < len(node.Content); i += 2 {
		// first print the key
		key := node.Content[i].Value
		result += fmt.Sprintf("%s\"%+v\": ", innerIndent, key)
		// then the value
		value := node.Content[i+1]
		switch value.Kind {
		case yaml.ScalarNode:
			result += "\"" + value.Value + "\""
		case yaml.MappingNode:
			result += renderMappingNode(value, innerIndent)
		case yaml.SequenceNode:
			result += renderSequenceNode(value, innerIndent)
		default:
			result += fmt.Sprintf("???MapItem(Key:%+v, Value:%T)", value, value)
		}
<<<<<<< HEAD
	default:
		// t is some other type that we didn't name.
>>>>>>> 79bfea2d (update vendor)
=======
		if i < len(node.Content)-2 {
			result += ","
		}
		result += "\n"
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
	}

	result += indent + "}"
	return result
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 03397665 (update api)
func renderSequenceNode(node *yaml.Node, indent string) (result string) {
	result = "[\n"
	innerIndent := indent + indentation
	for i := 0; i < len(node.Content); i++ {
		item := node.Content[i]
		switch item.Kind {
		case yaml.ScalarNode:
			result += innerIndent + "\"" + item.Value + "\""
		case yaml.MappingNode:
			result += innerIndent + renderMappingNode(item, innerIndent) + ""
		default:
			result += innerIndent + fmt.Sprintf("???ArrayItem(%+v)", item)
		}
		if i < len(node.Content)-1 {
<<<<<<< HEAD
=======
func renderArray(array []interface{}, indent string) (result string) {
=======
func renderSequenceNode(node *yaml.Node, indent string) (result string) {
>>>>>>> e879a141 (alibabacloud machine-api provider)
	result = "[\n"
	innerIndent := indent + indentation
	for i := 0; i < len(node.Content); i++ {
		item := node.Content[i]
		switch item.Kind {
		case yaml.ScalarNode:
			result += innerIndent + "\"" + item.Value + "\""
		case yaml.MappingNode:
			result += innerIndent + renderMappingNode(item, innerIndent) + ""
		default:
			result += innerIndent + fmt.Sprintf("???ArrayItem(%+v)", item)
		}
<<<<<<< HEAD
		if i < len(array)-1 {
>>>>>>> 79bfea2d (update vendor)
=======
		if i < len(node.Content)-1 {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
			result += ","
		}
		result += "\n"
	}
	result += indent + "]"
	return result
}

func renderStringArray(array []string, indent string) (result string) {
	result = "[\n"
	innerIndent := indent + indentation
	for i, item := range array {
		result += innerIndent + "\"" + item + "\""
		if i < len(array)-1 {
			result += ","
		}
		result += "\n"
	}
	result += indent + "]"
	return result
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
// Render renders a yaml.Node as JSON
func Render(node *yaml.Node) string {
	if node.Kind == yaml.DocumentNode {
		if len(node.Content) == 1 {
			return Render(node.Content[0])
		}
	} else if node.Kind == yaml.MappingNode {
		return renderMappingNode(node, "") + "\n"
	} else if node.Kind == yaml.SequenceNode {
		return renderSequenceNode(node, "") + "\n"
	}
	return ""
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 03397665 (update api)
}

func (object *SchemaNumber) nodeValue() *yaml.Node {
	if object.Integer != nil {
		return nodeForInt64(*object.Integer)
	} else if object.Float != nil {
		return nodeForFloat64(*object.Float)
<<<<<<< HEAD
=======
func Render(info yaml.MapSlice) string {
	return renderMap(info, "") + "\n"
=======
>>>>>>> e879a141 (alibabacloud machine-api provider)
}

func (object *SchemaNumber) nodeValue() *yaml.Node {
	if object.Integer != nil {
		return nodeForInt64(*object.Integer)
	} else if object.Float != nil {
<<<<<<< HEAD
		return object.Float
>>>>>>> 79bfea2d (update vendor)
=======
		return nodeForFloat64(*object.Float)
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
	} else {
		return nil
	}
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (object *SchemaOrBoolean) nodeValue() *yaml.Node {
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.Boolean != nil {
		return nodeForBoolean(*object.Boolean)
=======
func (object *SchemaOrBoolean) jsonValue() interface{} {
=======
func (object *SchemaOrBoolean) nodeValue() *yaml.Node {
>>>>>>> e879a141 (alibabacloud machine-api provider)
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.Boolean != nil {
<<<<<<< HEAD
		return *object.Boolean
>>>>>>> 79bfea2d (update vendor)
=======
		return nodeForBoolean(*object.Boolean)
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (object *SchemaOrBoolean) nodeValue() *yaml.Node {
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.Boolean != nil {
		return nodeForBoolean(*object.Boolean)
>>>>>>> 03397665 (update api)
	} else {
		return nil
	}
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
func nodeForStringArray(array []string) *yaml.Node {
	content := make([]*yaml.Node, 0)
	for _, item := range array {
		content = append(content, nodeForString(item))
	}
	return nodeForSequence(content)
}

func nodeForSchemaArray(array []*Schema) *yaml.Node {
	content := make([]*yaml.Node, 0)
	for _, item := range array {
		content = append(content, item.nodeValue())
	}
	return nodeForSequence(content)
}

func (object *StringOrStringArray) nodeValue() *yaml.Node {
<<<<<<< HEAD
<<<<<<< HEAD
	if object.String != nil {
		return nodeForString(*object.String)
	} else if object.StringArray != nil {
		return nodeForStringArray(*(object.StringArray))
=======
func (object *StringOrStringArray) jsonValue() interface{} {
=======
>>>>>>> e879a141 (alibabacloud machine-api provider)
	if object.String != nil {
		return nodeForString(*object.String)
	} else if object.StringArray != nil {
<<<<<<< HEAD
		array := make([]interface{}, 0)
		for _, item := range *(object.StringArray) {
			array = append(array, item)
		}
		return array
>>>>>>> 79bfea2d (update vendor)
=======
		return nodeForStringArray(*(object.StringArray))
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
	if object.String != nil {
		return nodeForString(*object.String)
	} else if object.StringArray != nil {
		return nodeForStringArray(*(object.StringArray))
>>>>>>> 03397665 (update api)
	} else {
		return nil
	}
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (object *SchemaOrStringArray) nodeValue() *yaml.Node {
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.StringArray != nil {
		return nodeForStringArray(*(object.StringArray))
=======
func (object *SchemaOrStringArray) jsonValue() interface{} {
=======
func (object *SchemaOrStringArray) nodeValue() *yaml.Node {
>>>>>>> e879a141 (alibabacloud machine-api provider)
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.StringArray != nil {
<<<<<<< HEAD
		array := make([]interface{}, 0)
		for _, item := range *(object.StringArray) {
			array = append(array, item)
		}
		return array
>>>>>>> 79bfea2d (update vendor)
=======
		return nodeForStringArray(*(object.StringArray))
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (object *SchemaOrStringArray) nodeValue() *yaml.Node {
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.StringArray != nil {
		return nodeForStringArray(*(object.StringArray))
>>>>>>> 03397665 (update api)
	} else {
		return nil
	}
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (object *SchemaOrSchemaArray) nodeValue() *yaml.Node {
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.SchemaArray != nil {
		return nodeForSchemaArray(*(object.SchemaArray))
=======
func (object *SchemaOrSchemaArray) jsonValue() interface{} {
=======
func (object *SchemaOrSchemaArray) nodeValue() *yaml.Node {
>>>>>>> e879a141 (alibabacloud machine-api provider)
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.SchemaArray != nil {
<<<<<<< HEAD
		array := make([]interface{}, 0)
		for _, item := range *(object.SchemaArray) {
			array = append(array, item.jsonValue())
		}
		return array
>>>>>>> 79bfea2d (update vendor)
=======
		return nodeForSchemaArray(*(object.SchemaArray))
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (object *SchemaOrSchemaArray) nodeValue() *yaml.Node {
	if object.Schema != nil {
		return object.Schema.nodeValue()
	} else if object.SchemaArray != nil {
		return nodeForSchemaArray(*(object.SchemaArray))
>>>>>>> 03397665 (update api)
	} else {
		return nil
	}
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 03397665 (update api)
func (object *SchemaEnumValue) nodeValue() *yaml.Node {
	if object.String != nil {
		return nodeForString(*object.String)
	} else if object.Bool != nil {
		return nodeForBoolean(*object.Bool)
<<<<<<< HEAD
=======
func (object *SchemaEnumValue) jsonValue() interface{} {
=======
func (object *SchemaEnumValue) nodeValue() *yaml.Node {
>>>>>>> e879a141 (alibabacloud machine-api provider)
	if object.String != nil {
		return nodeForString(*object.String)
	} else if object.Bool != nil {
<<<<<<< HEAD
		return *object.Bool
>>>>>>> 79bfea2d (update vendor)
=======
		return nodeForBoolean(*object.Bool)
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
	} else {
		return nil
	}
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func nodeForNamedSchemaArray(array *[]*NamedSchema) *yaml.Node {
	content := make([]*yaml.Node, 0)
	for _, pair := range *(array) {
		content = appendPair(content, pair.Name, pair.Value.nodeValue())
	}
	return nodeForMapping(content)
}

func nodeForNamedSchemaOrStringArray(array *[]*NamedSchemaOrStringArray) *yaml.Node {
	content := make([]*yaml.Node, 0)
	for _, pair := range *(array) {
		content = appendPair(content, pair.Name, pair.Value.nodeValue())
	}
	return nodeForMapping(content)
}

func nodeForSchemaEnumArray(array *[]SchemaEnumValue) *yaml.Node {
	content := make([]*yaml.Node, 0)
	for _, item := range *array {
		content = append(content, item.nodeValue())
	}
	return nodeForSequence(content)
}

func nodeForMapping(content []*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: content,
	}
}

func nodeForSequence(content []*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Content: content,
	}
}

func nodeForString(value string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!str",
		Value: value,
	}
}

func nodeForBoolean(value bool) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!bool",
		Value: fmt.Sprintf("%t", value),
	}
}

func nodeForInt64(value int64) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!int",
		Value: fmt.Sprintf("%d", value),
	}
}

func nodeForFloat64(value float64) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!float",
		Value: fmt.Sprintf("%f", value),
	}
}

func appendPair(nodes []*yaml.Node, name string, value *yaml.Node) []*yaml.Node {
	nodes = append(nodes, nodeForString(name))
	nodes = append(nodes, value)
	return nodes
}

func (schema *Schema) nodeValue() *yaml.Node {
	n := &yaml.Node{Kind: yaml.MappingNode}
	content := make([]*yaml.Node, 0)
	if schema.Title != nil {
		content = appendPair(content, "title", nodeForString(*schema.Title))
	}
	if schema.ID != nil {
		content = appendPair(content, "id", nodeForString(*schema.ID))
	}
	if schema.Schema != nil {
		content = appendPair(content, "$schema", nodeForString(*schema.Schema))
	}
	if schema.Type != nil {
		content = appendPair(content, "type", schema.Type.nodeValue())
	}
	if schema.Items != nil {
		content = appendPair(content, "items", schema.Items.nodeValue())
	}
	if schema.Description != nil {
		content = appendPair(content, "description", nodeForString(*schema.Description))
	}
	if schema.Required != nil {
		content = appendPair(content, "required", nodeForStringArray(*schema.Required))
	}
	if schema.AdditionalProperties != nil {
		content = appendPair(content, "additionalProperties", schema.AdditionalProperties.nodeValue())
	}
	if schema.PatternProperties != nil {
		content = appendPair(content, "patternProperties", nodeForNamedSchemaArray(schema.PatternProperties))
	}
	if schema.Properties != nil {
		content = appendPair(content, "properties", nodeForNamedSchemaArray(schema.Properties))
	}
	if schema.Dependencies != nil {
		content = appendPair(content, "dependencies", nodeForNamedSchemaOrStringArray(schema.Dependencies))
	}
	if schema.Ref != nil {
		content = appendPair(content, "$ref", nodeForString(*schema.Ref))
	}
	if schema.MultipleOf != nil {
		content = appendPair(content, "multipleOf", schema.MultipleOf.nodeValue())
	}
	if schema.Maximum != nil {
		content = appendPair(content, "maximum", schema.Maximum.nodeValue())
	}
	if schema.ExclusiveMaximum != nil {
		content = appendPair(content, "exclusiveMaximum", nodeForBoolean(*schema.ExclusiveMaximum))
	}
	if schema.Minimum != nil {
		content = appendPair(content, "minimum", schema.Minimum.nodeValue())
	}
	if schema.ExclusiveMinimum != nil {
		content = appendPair(content, "exclusiveMinimum", nodeForBoolean(*schema.ExclusiveMinimum))
	}
	if schema.MaxLength != nil {
		content = appendPair(content, "maxLength", nodeForInt64(*schema.MaxLength))
	}
	if schema.MinLength != nil {
		content = appendPair(content, "minLength", nodeForInt64(*schema.MinLength))
	}
	if schema.Pattern != nil {
		content = appendPair(content, "pattern", nodeForString(*schema.Pattern))
	}
	if schema.AdditionalItems != nil {
		content = appendPair(content, "additionalItems", schema.AdditionalItems.nodeValue())
	}
	if schema.MaxItems != nil {
		content = appendPair(content, "maxItems", nodeForInt64(*schema.MaxItems))
	}
	if schema.MinItems != nil {
		content = appendPair(content, "minItems", nodeForInt64(*schema.MinItems))
	}
	if schema.UniqueItems != nil {
		content = appendPair(content, "uniqueItems", nodeForBoolean(*schema.UniqueItems))
	}
	if schema.MaxProperties != nil {
		content = appendPair(content, "maxProperties", nodeForInt64(*schema.MaxProperties))
	}
	if schema.MinProperties != nil {
		content = appendPair(content, "minProperties", nodeForInt64(*schema.MinProperties))
	}
	if schema.Enumeration != nil {
		content = appendPair(content, "enum", nodeForSchemaEnumArray(schema.Enumeration))
	}
	if schema.AllOf != nil {
		content = appendPair(content, "allOf", nodeForSchemaArray(*schema.AllOf))
	}
	if schema.AnyOf != nil {
		content = appendPair(content, "anyOf", nodeForSchemaArray(*schema.AnyOf))
	}
	if schema.OneOf != nil {
		content = appendPair(content, "oneOf", nodeForSchemaArray(*schema.OneOf))
	}
	if schema.Not != nil {
		content = appendPair(content, "not", schema.Not.nodeValue())
	}
	if schema.Definitions != nil {
		content = appendPair(content, "definitions", nodeForNamedSchemaArray(schema.Definitions))
	}
	if schema.Default != nil {
		// m = append(m, yaml.MapItem{Key: "default", Value: *schema.Default})
	}
	if schema.Format != nil {
		content = appendPair(content, "format", nodeForString(*schema.Format))
	}
	n.Content = content
	return n
=======
func namedSchemaArrayValue(array *[]*NamedSchema) interface{} {
	m2 := yaml.MapSlice{}
=======
func nodeForNamedSchemaArray(array *[]*NamedSchema) *yaml.Node {
	content := make([]*yaml.Node, 0)
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func nodeForNamedSchemaArray(array *[]*NamedSchema) *yaml.Node {
	content := make([]*yaml.Node, 0)
>>>>>>> 03397665 (update api)
	for _, pair := range *(array) {
		content = appendPair(content, pair.Name, pair.Value.nodeValue())
	}
	return nodeForMapping(content)
}

func nodeForNamedSchemaOrStringArray(array *[]*NamedSchemaOrStringArray) *yaml.Node {
	content := make([]*yaml.Node, 0)
	for _, pair := range *(array) {
		content = appendPair(content, pair.Name, pair.Value.nodeValue())
	}
	return nodeForMapping(content)
}

func nodeForSchemaEnumArray(array *[]SchemaEnumValue) *yaml.Node {
	content := make([]*yaml.Node, 0)
	for _, item := range *array {
		content = append(content, item.nodeValue())
	}
	return nodeForSequence(content)
}

func nodeForMapping(content []*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: content,
	}
}

func nodeForSequence(content []*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Content: content,
	}
}

func nodeForString(value string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!str",
		Value: value,
	}
}

func nodeForBoolean(value bool) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!bool",
		Value: fmt.Sprintf("%t", value),
	}
}

func nodeForInt64(value int64) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!int",
		Value: fmt.Sprintf("%d", value),
	}
}

func nodeForFloat64(value float64) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!float",
		Value: fmt.Sprintf("%f", value),
	}
}

func appendPair(nodes []*yaml.Node, name string, value *yaml.Node) []*yaml.Node {
	nodes = append(nodes, nodeForString(name))
	nodes = append(nodes, value)
	return nodes
}

func (schema *Schema) nodeValue() *yaml.Node {
	n := &yaml.Node{Kind: yaml.MappingNode}
	content := make([]*yaml.Node, 0)
	if schema.Title != nil {
		content = appendPair(content, "title", nodeForString(*schema.Title))
	}
	if schema.ID != nil {
		content = appendPair(content, "id", nodeForString(*schema.ID))
	}
	if schema.Schema != nil {
		content = appendPair(content, "$schema", nodeForString(*schema.Schema))
	}
	if schema.Type != nil {
		content = appendPair(content, "type", schema.Type.nodeValue())
	}
	if schema.Items != nil {
		content = appendPair(content, "items", schema.Items.nodeValue())
	}
	if schema.Description != nil {
		content = appendPair(content, "description", nodeForString(*schema.Description))
	}
	if schema.Required != nil {
		content = appendPair(content, "required", nodeForStringArray(*schema.Required))
	}
	if schema.AdditionalProperties != nil {
		content = appendPair(content, "additionalProperties", schema.AdditionalProperties.nodeValue())
	}
	if schema.PatternProperties != nil {
		content = appendPair(content, "patternProperties", nodeForNamedSchemaArray(schema.PatternProperties))
	}
	if schema.Properties != nil {
		content = appendPair(content, "properties", nodeForNamedSchemaArray(schema.Properties))
	}
	if schema.Dependencies != nil {
		content = appendPair(content, "dependencies", nodeForNamedSchemaOrStringArray(schema.Dependencies))
	}
	if schema.Ref != nil {
		content = appendPair(content, "$ref", nodeForString(*schema.Ref))
	}
	if schema.MultipleOf != nil {
		content = appendPair(content, "multipleOf", schema.MultipleOf.nodeValue())
	}
	if schema.Maximum != nil {
		content = appendPair(content, "maximum", schema.Maximum.nodeValue())
	}
	if schema.ExclusiveMaximum != nil {
		content = appendPair(content, "exclusiveMaximum", nodeForBoolean(*schema.ExclusiveMaximum))
	}
	if schema.Minimum != nil {
		content = appendPair(content, "minimum", schema.Minimum.nodeValue())
	}
	if schema.ExclusiveMinimum != nil {
		content = appendPair(content, "exclusiveMinimum", nodeForBoolean(*schema.ExclusiveMinimum))
	}
	if schema.MaxLength != nil {
		content = appendPair(content, "maxLength", nodeForInt64(*schema.MaxLength))
	}
	if schema.MinLength != nil {
		content = appendPair(content, "minLength", nodeForInt64(*schema.MinLength))
	}
	if schema.Pattern != nil {
		content = appendPair(content, "pattern", nodeForString(*schema.Pattern))
	}
	if schema.AdditionalItems != nil {
		content = appendPair(content, "additionalItems", schema.AdditionalItems.nodeValue())
	}
	if schema.MaxItems != nil {
		content = appendPair(content, "maxItems", nodeForInt64(*schema.MaxItems))
	}
	if schema.MinItems != nil {
		content = appendPair(content, "minItems", nodeForInt64(*schema.MinItems))
	}
	if schema.UniqueItems != nil {
		content = appendPair(content, "uniqueItems", nodeForBoolean(*schema.UniqueItems))
	}
	if schema.MaxProperties != nil {
		content = appendPair(content, "maxProperties", nodeForInt64(*schema.MaxProperties))
	}
	if schema.MinProperties != nil {
		content = appendPair(content, "minProperties", nodeForInt64(*schema.MinProperties))
	}
	if schema.Enumeration != nil {
		content = appendPair(content, "enum", nodeForSchemaEnumArray(schema.Enumeration))
	}
	if schema.AllOf != nil {
		content = appendPair(content, "allOf", nodeForSchemaArray(*schema.AllOf))
	}
	if schema.AnyOf != nil {
		content = appendPair(content, "anyOf", nodeForSchemaArray(*schema.AnyOf))
	}
	if schema.OneOf != nil {
		content = appendPair(content, "oneOf", nodeForSchemaArray(*schema.OneOf))
	}
	if schema.Not != nil {
		content = appendPair(content, "not", schema.Not.nodeValue())
	}
	if schema.Definitions != nil {
		content = appendPair(content, "definitions", nodeForNamedSchemaArray(schema.Definitions))
	}
	if schema.Default != nil {
		// m = append(m, yaml.MapItem{Key: "default", Value: *schema.Default})
	}
	if schema.Format != nil {
		content = appendPair(content, "format", nodeForString(*schema.Format))
	}
<<<<<<< HEAD
<<<<<<< HEAD
	return m
>>>>>>> 79bfea2d (update vendor)
=======
	n.Content = content
	return n
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
	n.Content = content
	return n
>>>>>>> 03397665 (update api)
}

// JSONString returns a json representation of a schema.
func (schema *Schema) JSONString() string {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	node := schema.nodeValue()
	return Render(node)
=======
	info := schema.jsonValue()
	return Render(info)
>>>>>>> 79bfea2d (update vendor)
=======
	node := schema.nodeValue()
	return Render(node)
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
	node := schema.nodeValue()
	return Render(node)
>>>>>>> 03397665 (update api)
}
