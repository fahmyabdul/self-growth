import json, string, requests
from secrets import choice

from app import logger
    
def remove_nulls(d:dict):
	return {k: v for k, v in d.items() if v is not None and v != []}
    
def create_random_password(pass_length):
	return ''.join([choice(string.ascii_uppercase + string.digits) for _ in range(pass_length)])

def send_api(url,param):
	headers = {'content-type': 'application/json'}
	logger.log.info("Target : [ %s ]" % url)

	jsonRequest = json.dumps(param)
	logger.log.info("JSON Request : %s" % jsonRequest)
	
	r = requests.post(url, data = jsonRequest, headers = headers)
	logger.log.info("Status code : [ %s ]" % str(r.status_code))
	logger.log.info("Response API : %s" % str(r.text))