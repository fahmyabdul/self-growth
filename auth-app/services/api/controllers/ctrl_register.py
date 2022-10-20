import json

from flask import jsonify, request
from flask_restful import Resource

from internal.bispro.bispro_register import BisproRegister

class CtrlRegister(Resource):
    def post(self):
        """
        Auth Register
        ---
        tags: ['register']
        consumes:
            - application/x-www-form-urlencoded
        description: Register New JWT Account
        parameters:
            - name: phone
              in: formData
              required: true
              type: string
            - name: name
              in: formData
              required: true
              type: string
            - name: role
              in: formData
              required: true
              type: string
        responses:
          200:
            description: Response for Successful Registration
            schema:
              id: register.success
              properties:
                message:
                  type: string
                  example: "Registration success"
                data:
                  type: object
                  properties:
                    password:
                      type: string
          400:
            description: Response for Invalid Requests
            schema:
              id: register.invalid
              properties:
                message:
                  type: string
                  example: Invalid request, check the request again!
                data:
                  type: string
                  example: ""
          500:
            description: Response for Unexpected server error
            schema:
              id: register.server-error
              properties:
                message:
                  type: string
                  example: Unexpected server error
                data:
                  type: string
                  example: ""
        """
        request_data = request.form

        current_bispro = BisproRegister(request_data)
        response_payload, status = current_bispro.create_user()

        return response_payload, status