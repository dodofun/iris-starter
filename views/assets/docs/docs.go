// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/user": {
            "get": {
                "description": "根据用户ID获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.User": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "createAt": {
                    "type": "integer"
                },
                "enable": {
                    "type": "boolean"
                },
                "id": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "integer"
                },
                "mobile": {
                    "description": "@gotags: validate:\"gte=10000000000,lte=100000000000\"",
                    "type": "integer",
                    "maximum": 100000000000,
                    "minimum": 10000000000
                },
                "name": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "string"
                },
                "updateAt": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "API DOC",
	Description:      "API文档.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
