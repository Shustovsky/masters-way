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
        "/dev/reset-db": {
            "get": {
                "description": "resets db",
                "tags": [
                    "dev"
                ],
                "summary": "resets db",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/survey/user-intro": {
            "post": {
                "description": "Post survey user intro",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "survey"
                ],
                "summary": "Post survey user intro",
                "operationId": "survey-user-intro",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.PostSurveyUserIntroPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.PostSurveyUserIntroPayload": {
            "type": "object",
            "required": [
                "deviceId",
                "preferredInterfaceLanguage",
                "role",
                "source",
                "studentExperience",
                "studentGoals",
                "whyRegistered"
            ],
            "properties": {
                "deviceId": {
                    "type": "string"
                },
                "preferredInterfaceLanguage": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "studentExperience": {
                    "type": "string"
                },
                "studentGoals": {
                    "type": "string"
                },
                "whyRegistered": {
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
	BasePath:         "/survey",
	Schemes:          []string{},
	Title:            "Masters way survey API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
