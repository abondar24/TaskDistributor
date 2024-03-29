// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Alex",
            "email": "abondar24@yahoo.com"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/task": {
            "post": {
                "description": "Create a new task and send it to the queue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "description": "Task name",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TaskResponse"
                        }
                    },
                    "502": {
                        "description": "Failed to send command",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/{id}": {
            "put": {
                "description": "Change status of existing task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Complete task. Possible values: true/false",
                        "name": "complete",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Wrong id or completed param",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Failed to send command",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete existing task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Delete task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Wrong id param",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Failed to send command",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.TaskRequest": {
            "type": "object",
            "properties": {
                "taskName": {
                    "type": "string"
                }
            }
        },
        "model.TaskResponse": {
            "type": "object",
            "properties": {
                "taskId": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Task API",
	Description:      "Task API to send commands as tasks",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
