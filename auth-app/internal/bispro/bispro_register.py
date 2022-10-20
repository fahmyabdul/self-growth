import time, json, datetime

from pydantic import ValidationError
from typing import Optional, List

from app import logger
from configs import configs
from internal.common import common
from internal.models.models_users import RequestUser, Users

class BisproRegister:
    def __init__(self, request_data):
        self.password_length = 4
        self.request_data = request_data
    
    def create_user(self):
        if self.request_data == None:
            return {'message': "Invalid request, request body must not empty!", 'data': None }, 400
        
        logger.log.info("API | Register | Request: {}".format(self.request_data))

        # Load & Validate Form request
        try:
            self.models_user = Users.parse_obj(self.request_data)
        except ValidationError as err:
            logger.log.warn("API | Register | Request: {} | Error: {}".format(self.request_data, err))
            return {'message': "Invalid request, check the request again!", 'data': None }, 400
        
        # If Form Request Valid

        # Check whether phone number exists
        exist_user, err = self.models_user.get_by_filter([
            "phone = '%s'" % (self.models_user.phone),
        ])
        if err != None:
            logger.log.warn("API | Register | Request: {} | Error: {}".format(self.request_data, err))
            return {'message': "Unexpected server error", 'data': "" }, 500

        if len(exist_user) > 0:
            logger.log.warn("API | Register | Request: {} | Error: Phone number already exists".format(self.request_data))
            return {'message': "Phone number already exists", 'data': "" }, 400

        # Check password already exists
        password_check_status = True
        while password_check_status == True: 
            # Generate random password
            gen_password = common.create_random_password(self.password_length)
            exist_password, err = self.models_user.check_password(gen_password)
            if err != None:
                logger.log.warn("API | Register | Request: {} | Error: {}".format(self.request_data, err))
                return {'message': "Unexpected server error", 'data': "" }, 500

            if len(exist_password) == 0:
                password_check_status = False
            else:
                time.sleep(1)

        self.models_user.password = gen_password
        self.now = datetime.datetime.now()
        self.models_user.created_at = datetime.datetime.strftime(self.now, '%Y-%m-%d %H:%M:%S')

        err = self.models_user.register()
        if err != None:
            logger.log.warn("API | Register | Request: {} | Error: {}".format(self.request_data, err))
            return {'message': "Unexpected server error", 'data': "" }, 500

        return {'message': "Registration success", 'data': {"password": self.models_user.password} }, 200