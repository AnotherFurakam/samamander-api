// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/post": {
            "get": {
                "description": "Get Posts",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "List Posts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pageNumber",
                        "name": "pageNumber",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "Create post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Create post",
                "parameters": [
                    {
                        "description": "postDto",
                        "name": "postDto",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.PostDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/post/{postId}": {
            "put": {
                "description": "Update post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Update post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "postId",
                        "name": "postId",
                        "in": "path"
                    },
                    {
                        "description": "postDto",
                        "name": "postDto",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.PostDto"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "delete post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Delete post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "postId",
                        "name": "postId",
                        "in": "path"
                    }
                ],
                "responses": {}
            }
        },
        "/product": {
            "get": {
                "description": "get products",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "List products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pageNumber",
                        "name": "pageNumber",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "create product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create product",
                "parameters": [
                    {
                        "description": "productDto",
                        "name": "productDto",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.ProductDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/product/{productId}": {
            "get": {
                "description": "get product by id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Find products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productId",
                        "name": "productId",
                        "in": "path"
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "update product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productId",
                        "name": "productId",
                        "in": "path"
                    },
                    {
                        "description": "productDto",
                        "name": "productDto",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.ProductDto"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "delete product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productId",
                        "name": "productId",
                        "in": "path"
                    }
                ],
                "responses": {}
            }
        },
        "/user": {
            "get": {
                "description": "get users",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "List users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pageNumber",
                        "name": "pageNumber",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create users",
                "parameters": [
                    {
                        "description": "userDto",
                        "name": "userDto",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.CreateUserDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/{idUser}": {
            "get": {
                "description": "get user by id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Find users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "idUser",
                        "name": "idUser",
                        "in": "path"
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "idUser",
                        "name": "idUser",
                        "in": "path"
                    },
                    {
                        "description": "userDto",
                        "name": "userDto",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserDto"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete user by id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "idUser",
                        "name": "idUser",
                        "in": "path"
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.CreateUserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.PostDto": {
            "type": "object",
            "required": [
                "body",
                "title",
                "urlImage"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "maxLength": 700,
                    "minLength": 5
                },
                "title": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 5
                },
                "urlImage": {
                    "type": "string",
                    "maxLength": 700,
                    "minLength": 5
                }
            }
        },
        "model.ProductDto": {
            "type": "object",
            "required": [
                "description",
                "name",
                "urlImage"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 500,
                    "minLength": 10
                },
                "name": {
                    "type": "string",
                    "maxLength": 300,
                    "minLength": 5
                },
                "urlImage": {
                    "type": "string",
                    "maxLength": 500,
                    "minLength": 10
                }
            }
        },
        "model.UpdateUserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Samamander API",
	Description:      "This is a Samamander API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
