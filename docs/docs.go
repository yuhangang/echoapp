// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/mit-license.php"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login using username and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login using username and password.",
                "parameters": [
                    {
                        "description": "User name and Password for logged-in.",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success to the authentication.",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    },
                    "401": {
                        "description": "Failed to the authentication.",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/auth/loginAccount": {
            "get": {
                "description": "Get the account data of logged-in user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get the account data of logged-in user.",
                "responses": {
                    "200": {
                        "description": "Success to fetch the account data. If the security function is disable, it returns the dummy data.",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    },
                    "401": {
                        "description": "The current user haven't logged-in yet. Returns false.",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/auth/loginStatus": {
            "get": {
                "description": "Get the login status of current logged-in user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get the login status.",
                "responses": {
                    "200": {
                        "description": "The current user have already logged-in. Returns true.",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "401": {
                        "description": "The current user haven't logged-in yet. Returns false.",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Logout.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout.",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/books": {
            "get": {
                "description": "Get the list of matched books by searching",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Get a book list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Keyword",
                        "name": "query",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Item size per page",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success to fetch a book list.",
                        "schema": {
                            "$ref": "#/definitions/model.Page"
                        }
                    },
                    "400": {
                        "description": "Failed to fetch data.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Failed to the authentication. Returns false.",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Create a new book",
                "parameters": [
                    {
                        "description": "a new book data for creating",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BookDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success to create a new book.",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    },
                    "400": {
                        "description": "Failed to the registration.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Failed to the authentication. Returns false.",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/books/{book_id}": {
            "get": {
                "description": "Get a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Get a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success to fetch data.",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    },
                    "400": {
                        "description": "Failed to fetch data.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Failed to the authentication. Returns false.",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the existing book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Update the existing book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the book data for updating",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BookDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success to update the existing book.",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    },
                    "400": {
                        "description": "Failed to the update.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Failed to the authentication. Returns false.",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the existing book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Delete the existing book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success to delete the existing book.",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    },
                    "400": {
                        "description": "Failed to the delete.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Failed to the authentication. Returns false.",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/categories": {
            "get": {
                "description": "Get a category list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Get a category list",
                "responses": {
                    "200": {
                        "description": "Success to fetch a category list.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Category"
                            }
                        }
                    },
                    "401": {
                        "description": "Failed to the authentication.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/formats": {
            "get": {
                "description": "Get a format list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Formats"
                ],
                "summary": "Get a format list",
                "responses": {
                    "200": {
                        "description": "Success to fetch a format list.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Format"
                            }
                        }
                    },
                    "401": {
                        "description": "Failed to the authentication.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Get the status of this application",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Get the status of this application",
                "responses": {
                    "200": {
                        "description": "healthy: This application is started.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "None: This application is stopped.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BookDto": {
            "type": "object",
            "required": [
                "isbn",
                "title"
            ],
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "formatId": {
                    "type": "integer"
                },
                "isbn": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.LoginDto": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.Account": {
            "type": "object",
            "properties": {
                "authority": {
                    "$ref": "#/definitions/model.Authority"
                },
                "authority_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Authority": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Book": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/model.Category"
                },
                "categoryId": {
                    "type": "integer"
                },
                "format": {
                    "$ref": "#/definitions/model.Format"
                },
                "formatId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "isbn": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Category": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Format": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Page": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Book"
                    }
                },
                "last": {
                    "type": "boolean"
                },
                "numberOfElements": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "totalElements": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
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
	Version:     "1.5.1",
	Host:        "localhost:8080",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "go-webapp-sample API",
	Description: "This is API specification for go-webapp-sample project.",
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