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
        "/rpc": {
            "post": {
                "description": "fetch task by id,status or history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get taks",
                "parameters": [
                    {
                        "description": "RPC Request",
                        "name": "rpcRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rpc.TaskRPCRequest"
                        }
                    },
                    {
                        "description": "RPC Request for status",
                        "name": "rpcRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rpc.TaskRPCRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/data.TaskHistory"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "data.Task": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/data.TaskStatus"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "data.TaskHistory": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "statusHistory": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/data.TaskStatusEntry"
                    }
                }
            }
        },
        "data.TaskStatus": {
            "type": "string",
            "enum": [
                "created",
                "updated",
                "deleted",
                "completed"
            ],
            "x-enum-varnames": [
                "TASK_CREATED",
                "TASK_UPDATED",
                "TASK_DELETED",
                "TASK_COMPLETED"
            ]
        },
        "data.TaskStatusEntry": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/data.TaskStatus"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "rpc.TaskRPCRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "request id",
                    "type": "integer"
                },
                "jsonrpc": {
                    "description": "json-rpc version 2",
                    "type": "string"
                },
                "method": {
                    "description": "json-rpc method: TaskRPC.GetTask, TaskRPC.GetTaskHistory",
                    "type": "string"
                },
                "params": {
                    "description": "params",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "",
	Description:      "Task store - accepts commands and exposes JSON-RPC API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
