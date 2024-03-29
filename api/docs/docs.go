// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-08-26 23:36:18.676969 -0400 EDT m=+0.058398009

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Swagger API for Yet Another Issue Tracking System.",
        "title": "YAITS Swagger API",
        "contact": {
            "name": "API Support",
            "email": "anhkhoi.vunguyen@gmai.com"
        },
        "license": {},
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/api",
    "paths": {
        "/issue": {
            "post": {
                "description": "Create a new issue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Creation"
                ],
                "summary": "Create an issue",
                "parameters": [
                    {
                        "description": "YAITS creation request",
                        "name": "issueRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.NewIssueRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.IssueIDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    }
                }
            }
        },
        "/issue/{id}": {
            "get": {
                "description": "Retrieves an issue given issue id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrieval"
                ],
                "summary": "Retrieves an issue given issue id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the issue",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.IssueResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an issue given an issue id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deletion"
                ],
                "summary": "Delete an issue",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the issue",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates an issue given an issue id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update"
                ],
                "summary": "Update an issue",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the issue",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "YAITS update request",
                        "name": "updateIssueRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.UpdateIssueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.IssueResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    }
                }
            }
        },
        "/issues": {
            "get": {
                "description": "Retrieves all issues",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrieval"
                ],
                "summary": "Retrieves all existing issues",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.IssueResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    }
                }
            }
        },
        "/issues/priority": {
            "get": {
                "description": "Retrieves an issue given priority",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrieval"
                ],
                "summary": "Retrieves an issue given priority",
                "parameters": [
                    {
                        "type": "models.PriorityQueryParam",
                        "description": "priority start bound",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "models.PriorityQueryParam",
                        "description": "priority end bound",
                        "name": "end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.IssueResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    }
                }
            }
        },
        "/issues/status": {
            "get": {
                "description": "Retrieves an issue given status (open, closed, in progress)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrieval"
                ],
                "summary": "Retrieves an issue given status",
                "parameters": [
                    {
                        "type": "models.StatusQueryParam",
                        "description": "issue priority request",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.IssueResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ErrorWrapper"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                }
            }
        },
        "models.ErrorWrapper": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.StandardError"
                    }
                }
            }
        },
        "models.IssueIDResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.IssueResponse": {
            "type": "object",
            "properties": {
                "assignee": {
                    "type": "string"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "createDate": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "priority": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                }
            }
        },
        "models.NewIssueRequest": {
            "type": "object",
            "required": [
                "description",
                "priority",
                "summary"
            ],
            "properties": {
                "assignee": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                }
            }
        },
        "models.StandardError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.UpdateIssueRequest": {
            "type": "object",
            "properties": {
                "assignee": {
                    "type": "string"
                },
                "comment": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "summary": {
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
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
