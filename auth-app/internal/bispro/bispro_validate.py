import json, jwt, datetime, time
from pydantic import ValidationError
from typing import Optional, List

from app import logger
from configs import configs
from internal.common import common
from internal.models.models_users import Users
from internal.models.models_validate import RequestValidate

class BisproValidate:
    def __init__(self, request_data):
        self.request_data = request_data
        self.jwt_conf = configs.properties["etc"]["jwt"]
    
    def validate(self):
        if self.request_data == None:
            return {'message': "Invalid request, request must not empty!", 'data': "" }, 400
        
        logger.log.info("API | Validate | Request: {}".format(self.request_data['jwt']))

        # Load & Validate Form request
        try:
            self.models_validate = RequestValidate.parse_obj(self.request_data)
        except ValidationError as err:
            logger.log.warn("API | Validate | Request: {} | Error: {}".format(self.request_data['jwt'], err))
            return {'message': "Invalid request, check the request again!", 'data': "" }, 400
        
        # If Form Request Valid
        # Decode JWT and check for invalid & expired signature
        try:
            decoded_jwt = jwt.decode(self.models_validate.jwt, self.jwt_conf["secret"], algorithms=["HS256"])
        except jwt.InvalidSignatureError:
            logger.log.warn("API | Validate | Request: {} | Error: Invalid JWT".format(self.request_data['jwt']))
            return {'message': "Invalid JWT", 'data': "" }, 401
        except jwt.ExpiredSignatureError:
            logger.log.warn("API | Validate | Request: {} | Error: Expired JWT".format(self.request_data['jwt']))
            return {'message': "Expired JWT", 'data': "" }, 401

        return {'message': "Valid JWT", 'data': decoded_jwt }, 200