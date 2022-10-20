import json

from flask import jsonify, request
from flask_restful import Resource

from internal.bispro.bispro_claims import BisproClaims

class CtrlClaims(Resource):
    def post(self):
        """
        Auth Claims
        ---
        tags: ['claims']
        consumes:
            - application/x-www-form-urlencoded
        description: Get Jwt Token Claims
        parameters:
            - name: phone
              in: formData
              required: true
              type: string
            - name: password
              in: formData
              required: true
              type: string
        responses:
          200:
            description: Response for Authorized claims
            schema:
              id: claims.authorized
              properties:
                message:
                  type: string
                  example: Authorized
                data:
                  type: object
                  properties:
                    jwt:
                      type: string
          400:
            description: Response for Unathorized claims
            schema:
              id: claims.unauthorized
              properties:
                message:
                  type: string
                  example: Unauthorized
                data:
                  type: string
                  example: ""
          500:
            description: Response for Unexpected server error
            schema:
              id: claims.server-error
              properties:
                message:
                  type: string
                  example: Unexpected server error
                data:
                  type: string
                  example: ""
        """
        request_data = request.form

        current_bispro = BisproClaims(request_data)
        response_payload, status = current_bispro.claims()

        return response_payload, status