// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-21 22:03:56.192706 +0800 CST m=+0.036159630

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://axiaoxin.com",
            "email": "254606826@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/demo/label": {
            "get": {
                "description": "query label by id or name, or paging query",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "label"
                ],
                "summary": "Query label by query params",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "query by ID, other conditions do not enter into force.",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "query by name, other conditions do not enter into force.",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "fuzzy query by remark",
                        "name": "remark",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "page number",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "page size",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "id asc",
                        "description": "order field and order type",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "give name and remark to add new label, return label ID, if label exist, update then remark field",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "label"
                ],
                "summary": "Add new label",
                "parameters": [
                    {
                        "description": "label info",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/demo.AddLabelBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/demo/labeling": {
            "put": {
                "description": "replace association for a given object ID list and tag ID list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labeling"
                ],
                "summary": "Update labeling by replace way",
                "parameters": [
                    {
                        "description": "labeling association info",
                        "name": "labeling",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/demo.AddLabelingBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "create association for a given object ID list and tag ID list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labeling"
                ],
                "summary": "Labeling for object",
                "parameters": [
                    {
                        "description": "labeling association info",
                        "name": "labeling",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/demo.AddLabelingBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete association for a given object ID list and tag ID list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labeling"
                ],
                "summary": "Delete labeling",
                "parameters": [
                    {
                        "description": "labeling association info",
                        "name": "labeling",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/demo.AddLabelingBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/demo/labeling/label/{id}": {
            "get": {
                "description": "return labeling object list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labeling"
                ],
                "summary": "Query labeling object list by label ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "label id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/demo/labeling/object/{id}": {
            "get": {
                "description": "return labeling label list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labeling"
                ],
                "summary": "Query labeling label list by object ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "object id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/demo/object": {
            "get": {
                "description": "query object by params",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "object"
                ],
                "summary": "Query object by params",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "query by object ID, other conditions do not enter into force.",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter result list by appid",
                        "name": "appID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter result list by system",
                        "name": "system",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter result list by entity",
                        "name": "entity",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter result list by identity",
                        "name": "identity",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "page number",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "page size",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "id asc",
                        "description": "order field and way",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "add new object return object ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "object"
                ],
                "summary": "Add new object",
                "parameters": [
                    {
                        "description": "object info",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/demo.AddObjectBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/x/ping": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "x"
                ],
                "summary": "Ping for server is living will respond API version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "demo.AddLabelBody": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "description": "label name",
                    "type": "string",
                    "example": "label name"
                },
                "remark": {
                    "description": "label remark",
                    "type": "string",
                    "example": "label remark"
                }
            }
        },
        "demo.AddLabelingBody": {
            "type": "object",
            "required": [
                "labelIDs",
                "objectIDs"
            ],
            "properties": {
                "labelIDs": {
                    "description": "which label ids need to be labling with the object ids",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "objectIDs": {
                    "description": "which object ids need to be labling with the label ids",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "demo.AddObjectBody": {
            "type": "object",
            "required": [
                "appID",
                "entity",
                "identity",
                "system"
            ],
            "properties": {
                "appID": {
                    "description": "APP ID",
                    "type": "string",
                    "example": "1"
                },
                "entity": {
                    "description": "entity name",
                    "type": "string",
                    "example": "server"
                },
                "identity": {
                    "description": "identity",
                    "type": "string",
                    "example": "1"
                },
                "system": {
                    "description": "system name",
                    "type": "string",
                    "example": "cmdb"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "127.0.0.1:4869",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "pink-lady Web API",
	Description: "pink-lady web API list.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
