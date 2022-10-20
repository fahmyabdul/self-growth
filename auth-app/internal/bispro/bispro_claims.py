import json, jwt, datetime, time
from pydantic import ValidationError
from typing import Optional, List

from app import logger
from configs import configs
from internal.common import common
from internal.models.models_users import Users
from internal.models.models_claims import RequestClaims

class BisproClaims:
    def __init__(self, request_data):
        self.request_data = request_data
        self.jwt_conf = configs.properties["etc"]["jwt"]

    def claims(self):
        if self.request_data == None:
            return {'message': "Unauthorized", 'data': "" }, 400
        
        logger.log.info("API | Claims | Request: {}".format(self.request_data['phone']))

        # Load & Validate Form Data request
        try:
            self.models_claims = RequestClaims.parse_obj(self.request_data)
        except ValidationError as err:
            logger.log.warn("API | Claims | Request: {} | Error: {}".format(self.request_data['phone'], err))
            return {'message': "Unauthorized", 'data': "" }, 400
        
        # If Form Data Request Valid

        # Initialize User Models
        self.models_user = Users(
            phone="",
            name="",
            role="",
            created_at="",
            password=""
        )
        # Check whether phone number exists
        exist_user, err = self.models_user.get_by_filter([
            "phone = '%s'" % (self.models_claims.phone),
            "password = '%s'" % (self.models_claims.password)
        ])
        if err != None:
            logger.log.warn("API | Claims | Request: {} | Error: {}".format(self.request_data['phone'], err))
            return {'message': "Unexpected server error", 'data': None }, 500

        if len(exist_user) == 0:
            logger.log.warn("API | Claims | Request: {} | Error: {}".format(self.request_data['phone'], err))
            return {'message': "Unauthorized", 'data': None }, 401

        # Setting JWT
        self.now = datetime.datetime.now()
        date_exp = self.now + datetime.timedelta(seconds=self.jwt_conf["expiration"])

        # Setting JWT Payload
        self.models_user = Users(
            phone=exist_user[0][1],
            name=exist_user[0][2],
            role=exist_user[0][3],
            created_at=exist_user[0][5],
            logged_in=datetime.datetime.strftime(self.now, '%Y-%m-%d %H:%M:%S'),
        )
        jwt_payload = json.loads(self.models_user.json(), object_hook=common.remove_nulls)
        jwt_payload["exp"] = (time.mktime(date_exp.timetuple()))

        # Encode JWT
        encoded_jwt = jwt.encode(jwt_payload, self.jwt_conf["secret"], algorithm="HS256")

        return {'message': "Authorized", 'data': {"jwt": encoded_jwt} }, 200