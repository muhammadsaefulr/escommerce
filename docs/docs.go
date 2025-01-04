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
        "/user/auth/login": {
            "post": {
                "description": "Authenticates a user and returns user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserCustomer"
                ],
                "summary": "User Customer Auth",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.AuthLoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully authenticated user",
                        "schema": {
                            "$ref": "#/definitions/entity.UserDataReturnViews"
                        }
                    }
                }
            }
        },
        "/user/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "Tokens": []
                    }
                ],
                "description": "Delete user by id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserCustomer"
                ],
                "summary": "Delete user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully delete user",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/user/get/{id}": {
            "get": {
                "description": "Get user by id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserCustomer"
                ],
                "summary": "Get user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get user",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register new user customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserCustomer"
                ],
                "summary": "Create new user customer",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created new user",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/user/seller/auth/login": {
            "post": {
                "description": "Authenticates a user seller and returns user seller data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserSeller"
                ],
                "summary": "UserSeller Customer Auth Login",
                "parameters": [
                    {
                        "description": "User seller data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.AuthLoginUserSeller"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully authenticated user seller",
                        "schema": {
                            "$ref": "#/definitions/entity.UserSeller"
                        }
                    }
                }
            }
        },
        "/user/seller/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "Tokens": []
                    }
                ],
                "description": "Delete User Seller By Id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserSeller"
                ],
                "summary": "Delete User Seller By Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Seller id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully delete User Seller",
                        "schema": {
                            "$ref": "#/definitions/entity.UserSeller"
                        }
                    }
                }
            }
        },
        "/user/seller/get/{id}": {
            "get": {
                "description": "Get user seller by id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserSeller"
                ],
                "summary": "Get user seller by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User seller id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get user seller",
                        "schema": {
                            "$ref": "#/definitions/entity.UserSeller"
                        }
                    }
                }
            }
        },
        "/user/seller/register": {
            "post": {
                "description": "Register new user seller customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserSeller"
                ],
                "summary": "Create new user seller customer",
                "parameters": [
                    {
                        "description": "User seller data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserSeller"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created new user seller",
                        "schema": {
                            "$ref": "#/definitions/entity.UserSeller"
                        }
                    }
                }
            }
        },
        "/user/seller/update/{id}": {
            "put": {
                "security": [
                    {
                        "Tokens": []
                    }
                ],
                "description": "Update User Seller Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserSeller"
                ],
                "summary": "Update User Seller Data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Seller id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Seller data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserSeller"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully update User Seller",
                        "schema": {
                            "$ref": "#/definitions/entity.UserSeller"
                        }
                    }
                }
            }
        },
        "/user/update/{id}": {
            "put": {
                "security": [
                    {
                        "Tokens": []
                    }
                ],
                "description": "Update user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserCustomer"
                ],
                "summary": "Update user data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully update user",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.AuthLoginUser": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 8
                },
                "password": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "entity.AuthLoginUserSeller": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 8
                },
                "password": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "entity.ProductItems": {
            "type": "object",
            "required": [
                "category_id",
                "description",
                "name",
                "price",
                "seller_id"
            ],
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "seller": {
                    "$ref": "#/definitions/entity.UserSeller"
                },
                "seller_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 75,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "minLength": 3
                },
                "role_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.UserDataReturnViews": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                }
            }
        },
        "entity.UserSeller": {
            "type": "object",
            "required": [
                "nama_toko",
                "user_id"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nama_toko": {
                    "type": "string",
                    "maxLength": 75,
                    "minLength": 3
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.ProductItems"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "description": "Menghubungkan penjual dengan pengguna",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Tokens": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Escommerce API",
	Description:      "Test Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
