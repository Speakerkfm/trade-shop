// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Trade Shop",
    "version": "0.1.0"
  },
  "paths": {
    "/inventory": {
      "get": {
        "description": "User inventory",
        "tags": [
          "inventory"
        ],
        "operationId": "inventory",
        "responses": {
          "200": {
            "description": "Inventory",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "401": {
            "description": "Пользователь не авторизован"
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "Аутентификация пользователя по почте и паролю.\n",
        "tags": [
          "login"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "description": "Информация пользователя для входа в систему",
              "type": "object",
              "required": [
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "type": "string",
                  "x-insullable": false
                },
                "password": {
                  "type": "string",
                  "x-insullable": false
                }
              },
              "example": {
                "email": "speaker@mail.com",
                "password": "123456"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Wrong username or password"
          }
        }
      }
    },
    "/sale": {
      "post": {
        "description": "Продажа предмета пользователем",
        "tags": [
          "login"
        ],
        "operationId": "sale",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "description": "Информация пользователя для входа в систему",
              "type": "object",
              "required": [
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "type": "string",
                  "x-insullable": false
                },
                "password": {
                  "type": "string",
                  "x-insullable": false
                }
              },
              "example": {
                "email": "speaker@mail.com",
                "password": "123456"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Пользователь не авторизован"
          }
        }
      }
    },
    "/sales": {
      "get": {
        "description": "Список лотов",
        "tags": [
          "sales"
        ],
        "operationId": "sales_list",
        "responses": {
          "200": {
            "description": "Список лотов",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/sale"
              }
            }
          },
          "401": {
            "description": "Пользователь не авторизован"
          }
        }
      }
    },
    "/sales/{sale_id}/buy": {
      "get": {
        "description": "Купить лот",
        "tags": [
          "buy"
        ],
        "operationId": "buy",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "идентификатор продажи",
            "name": "sale_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Пользователь не авторизован"
          }
        }
      }
    }
  },
  "definitions": {
    "item": {
      "description": "Item",
      "type": "object",
      "properties": {
        "count": {
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
    "sale": {
      "description": "Продажа",
      "properties": {
        "id": {
          "description": "id продажи",
          "type": "integer"
        },
        "items": {
          "type": "array",
          "items": {
            "description": "Предметы на продажу",
            "type": "object",
            "properties": {
              "count": {
                "type": "integer"
              },
              "id": {
                "type": "integer"
              },
              "name": {
                "type": "string"
              },
              "price": {
                "type": "number"
              }
            }
          }
        },
        "total_count": {
          "description": "Общая стоимость лота",
          "type": "number"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Trade Shop",
    "version": "0.1.0"
  },
  "paths": {
    "/inventory": {
      "get": {
        "description": "User inventory",
        "tags": [
          "inventory"
        ],
        "operationId": "inventory",
        "responses": {
          "200": {
            "description": "Inventory",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "401": {
            "description": "Пользователь не авторизован"
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "Аутентификация пользователя по почте и паролю.\n",
        "tags": [
          "login"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "description": "Информация пользователя для входа в систему",
              "type": "object",
              "required": [
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "type": "string",
                  "x-insullable": false
                },
                "password": {
                  "type": "string",
                  "x-insullable": false
                }
              },
              "example": {
                "email": "speaker@mail.com",
                "password": "123456"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Wrong username or password"
          }
        }
      }
    },
    "/sale": {
      "post": {
        "description": "Продажа предмета пользователем",
        "tags": [
          "login"
        ],
        "operationId": "sale",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "description": "Информация пользователя для входа в систему",
              "type": "object",
              "required": [
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "type": "string",
                  "x-insullable": false
                },
                "password": {
                  "type": "string",
                  "x-insullable": false
                }
              },
              "example": {
                "email": "speaker@mail.com",
                "password": "123456"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Пользователь не авторизован"
          }
        }
      }
    },
    "/sales": {
      "get": {
        "description": "Список лотов",
        "tags": [
          "sales"
        ],
        "operationId": "sales_list",
        "responses": {
          "200": {
            "description": "Список лотов",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/sale"
              }
            }
          },
          "401": {
            "description": "Пользователь не авторизован"
          }
        }
      }
    },
    "/sales/{sale_id}/buy": {
      "get": {
        "description": "Купить лот",
        "tags": [
          "buy"
        ],
        "operationId": "buy",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "идентификатор продажи",
            "name": "sale_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Пользователь не авторизован"
          }
        }
      }
    }
  },
  "definitions": {
    "item": {
      "description": "Item",
      "type": "object",
      "properties": {
        "count": {
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
    "sale": {
      "description": "Продажа",
      "properties": {
        "id": {
          "description": "id продажи",
          "type": "integer"
        },
        "items": {
          "type": "array",
          "items": {
            "description": "Предметы на продажу",
            "type": "object",
            "properties": {
              "count": {
                "type": "integer"
              },
              "id": {
                "type": "integer"
              },
              "name": {
                "type": "string"
              },
              "price": {
                "type": "number"
              }
            }
          }
        },
        "total_count": {
          "description": "Общая стоимость лота",
          "type": "number"
        }
      }
    }
  }
}`))
}
