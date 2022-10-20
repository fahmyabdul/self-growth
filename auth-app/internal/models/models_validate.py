from pydantic import BaseModel, StrictStr

class RequestValidate(BaseModel):
    jwt: StrictStr