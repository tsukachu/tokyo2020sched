// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-06-13 19:54:43.192945 +0900 JST m=+0.028600778

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": [
        "https"
    ],
    "swagger": "2.0",
    "info": {
        "description": "TOKYO2020 schedule API",
        "title": "TOKYO2020 schedule API",
        "contact": {},
        "license": {},
        "version": "0.1"
    },
    "host": "tokyo2020sched.herokuapp.com",
    "basePath": "/",
    "paths": {
        "/classifications": {
            "get": {
                "description": "種別等の一覧を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Classifications"
                ],
                "summary": "種別等一覧を取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ClassificationWithCompetition"
                            }
                        }
                    }
                }
            }
        },
        "/classifications/{id}": {
            "get": {
                "description": "指定した種別等の詳細情報を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Classifications"
                ],
                "summary": "種別等の詳細情報を取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "種別等ID",
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
                            "$ref": "#/definitions/models.ClassificationWithCompetition"
                        }
                    }
                }
            }
        },
        "/competitions": {
            "get": {
                "description": "競技の一覧を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Competitions"
                ],
                "summary": "競技一覧を取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Competition"
                            }
                        }
                    }
                }
            }
        },
        "/competitions/{id}": {
            "get": {
                "description": "指定した競技の詳細情報を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Competitions"
                ],
                "summary": "競技の詳細情報を取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "競技ID",
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
                            "$ref": "#/definitions/models.Competition"
                        }
                    }
                }
            }
        },
        "/places": {
            "get": {
                "description": "場所の一覧を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Places"
                ],
                "summary": "場所一覧を取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Place"
                            }
                        }
                    }
                }
            }
        },
        "/places/{id}": {
            "get": {
                "description": "指定した場所の詳細情報を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Places"
                ],
                "summary": "場所の詳細情報を取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "場所ID",
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
                            "$ref": "#/definitions/models.Place"
                        }
                    }
                }
            }
        },
        "/schedules/olympic": {
            "get": {
                "description": "スケジュールの一覧を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OlympicSchedules"
                ],
                "summary": "スケジュール一覧を取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.OlympicSchedule"
                            }
                        }
                    }
                }
            }
        },
        "/schedules/olympic/{id}": {
            "get": {
                "description": "指定したスケジュールの詳細情報を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OlympicSchedules"
                ],
                "summary": "スケジュールの詳細情報を取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "スケジュールID",
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
                            "$ref": "#/definitions/models.OlympicSchedule"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Classification": {
            "type": "object",
            "properties": {
                "competitionId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.ClassificationWithCompetition": {
            "type": "object",
            "properties": {
                "competition": {
                    "type": "object",
                    "$ref": "#/definitions/models.Competition"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Competition": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.OlympicSchedule": {
            "type": "object",
            "properties": {
                "begin": {
                    "type": "string"
                },
                "classification": {
                    "type": "object",
                    "$ref": "#/definitions/models.Classification"
                },
                "competition": {
                    "type": "object",
                    "$ref": "#/definitions/models.Competition"
                },
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "end": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "place": {
                    "type": "object",
                    "$ref": "#/definitions/models.Place"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Place": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
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
