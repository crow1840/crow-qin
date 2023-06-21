// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://127.0.0.1",
        "contact": {
            "name": "刘炜",
            "url": "https://blog.csdn.net/xingzuo_1840",
            "email": "40010355@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/admin/user": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "填写用户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/admin/user-password": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "admin重置用户密码",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "admin重置用户密码",
                "parameters": [
                    {
                        "description": "填写用户Id和xin密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ResetUserPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/admin/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取用户信息",
                "tags": [
                    "admin"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户角色",
                        "name": "role",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户名（可模糊查询）",
                        "name": "user_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page_num",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页行数",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/service.GetUsersResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "修改用户信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "description": "user_id必须，其他可选",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UpdateUsersInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/admin/users/{uuid}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取指定用户信息",
                "tags": [
                    "admin"
                ],
                "summary": "获取指定用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SysUser"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "删除指定用户",
                "tags": [
                    "admin"
                ],
                "summary": "删除指定用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "用户登录，并获取token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名，用户密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/login/refresh": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "刷新用户token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "刷新token",
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/logout": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "用户登出",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "用户登出",
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user/info": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "用户获取本人信息",
                "tags": [
                    "user"
                ],
                "summary": "用户获取本人信息",
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SysUser"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "用户修改自己信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户修改自己信息",
                "parameters": [
                    {
                        "description": "选项可选",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.UpdateUserInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user/password": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "用户修改自己的密码",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户修改密码",
                "parameters": [
                    {
                        "description": "用户的新密码和旧密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.UpdateUserPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"requestId\": \"string\",\"code\": 200,\"msg\": \"ok\",\"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "{\"requestId\": \"string\",\"code\": 500,\"msg\": \"string\",\"status\": \"error\",\"data\": null}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.SysUser": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "requestId": {
                    "description": "数据集",
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "router.UpdateUserInfoRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "router.UpdateUserPasswordRequest": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "password": {
                    "description": "UserId      int64  ` + "`" + `json:\"user_id\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "service.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "service.GetUsersResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "page_num": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "sysUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SysUser"
                    }
                }
            }
        },
        "service.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "service.ResetUserPasswordRequest": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "service.UpdateUsersInfoRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "crow-qin",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
