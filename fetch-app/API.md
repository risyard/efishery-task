openapi: 3.0.0
info:
  version: 1.0.0
  title: Fetch-App
  description: A simple collections of API for fetch-app 

servers:
  - url: http://localhost:3000/

components:
  securitySchemes:
    bearerAuth:            
      type: http
      scheme: bearer
      bearerFormat: JWT 

security:
  - bearerAuth: []  

paths:
  /hello:
    get:
      description: Returns string "Hello World!"
      responses:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: string
                description: "Hello World!" string  

        
        default:
          description: Unregistered error/response

  /komoditas:
    get:
      description: Get list of clean commodities data with additional USD currency for its price 
      response:
        '200':
          description: Created!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  data:
                    type: array
                    description: List of all commodities data with additional USD currency for its price
                    properties:
                      uuid:
                        type: string
                        description: Commodity data ID
                      komoditas:
                        type: string
                        description: Commodity name/type
                      area_provinsi:
                        type: string
                        description: Commodity sell area, province level
                      area_kota:
                        type: string
                        description: Commodity sell area, city level
                      size:
                        type: string
                        description: Commodity's size/amount
                      price:
                        type: string
                        description: Commodity's each item's price in rupiah
                      tgl_parsed: 
                        type: string
                        description: Commodity's parsed date time
                      timestamp:
                        type: string
                        description: Timestamp
                      price_usd:
                        type: string
                        description: Commodity's each item's price in USD      
        
        '500'
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message

        default:
          description: Unregistered error/response

  /token:
    get:
      description: Generate JWT Token for stored user with correct phone number and password
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
                properties:
                  phone:
                    type: string
                    description: User's phone number
                  password:
                    type: string
                    description: User's password
      response:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  data:
                    type: string
                    description: Related user JWT Token
        
        '500'
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message

  /claims:
    get:
      description: Generate JWT Token for stored user with correct phone number and password
      security:
        - bearerAuth: []
      response:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  data:
                    type: object
                    properties:
                      phone:
                        type: string
                        description: User's phone number
                      name:
                        type: string
                        description: User's username
                      role:
                        type: string
                        description: User's role
                      timestamp:
                        type: string
                        description: String formatted timestamp 
        
        '500'
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message
