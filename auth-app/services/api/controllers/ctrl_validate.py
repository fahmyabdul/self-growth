import json

from flask import jsonify, request
from flask_restful import Resource

from internal.bispro.bispro_validate import BisproValidate

class CtrlValidate(Resource):
    def post(self):
        """
        Auth Validate
        ---
        tags: ['validate']
        consumes:
            - application/x-www-form-urlencoded
        description: Validate JWT Token
        parameters:
            - name: jwt
              in: formData
              required: true
              type: string
        responses:
          200:
            description: Response for Valid JWT Token
            schema:
              id: validate.success
              properties:
                message:
                  type: string
                  example: "Valid JWT"
                data:
                  type: object
                  properties:
                    phone:
                      type: string
                    name:
                      type: string
                    role:
                      type: string
                    created_at:
                      type: string
                    logged_in:
                      type: string
                    exp:
                      type: number
          400:
            description: Response for Invalid Requests
            schema:
              id: validate.invalid
              properties:
                message:
                  type: string
                  example: Invalid request, check the request again!
                data:
                  type: string
                  example: ""
          401:
            description: Response for Invalid JWT Token
            schema:
              id: validate.invalid
              properties:
                message:
                  type: string
                  example: Invalid JWT
                data:
                  type: string
                  example: ""
          500:
            description: Response for Unexpected server error
            schema:
              id: validate.server-error
              properties:
                message:
                  type: string
                  example: Unexpected server error
                data:
                  type: string
                  example: ""
        """
        request_data = request.form

        current_bispro = BisproValidate(request_data)
        response_payload, status = current_bispro.validate()

        return response_payload, status