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
        "/files": {
            "post": {
                "description": "Uploads a file to the server and stores it in the designated storage path",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Upload file to storage",
                "operationId": "upload-file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Room id",
                        "name": "roomId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UploadFileResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete multiple files from the server storage using their IDs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Delete files by IDs",
                "operationId": "delete-files",
                "parameters": [
                    {
                        "description": "List of file IDs to delete",
                        "name": "fileIDs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.UploadFileResponse": {
            "type": "object",
            "required": [
                "id",
                "name",
                "ownerId",
                "previewUrl",
                "srcUrl"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ownerId": {
                    "type": "string"
                },
                "previewUrl": {
                    "type": "string"
                },
                "srcUrl": {
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
	BasePath:         "/notification",
	Schemes:          []string{},
	Title:            "Masters way notification-bff API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
