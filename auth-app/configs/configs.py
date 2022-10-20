import yaml

properties = None

def init_config(filename: str):
    global properties
    with open(filename, 'r') as config_file:
        try:
            properties = yaml.load(config_file, Loader=yaml.FullLoader)
            if properties == None:
                print("Failed to load config file, error: file is empty")
                return None
            
            properties = properties
        except yaml.YAMLError as e:
            print ("Failed to load config file, error: {}".format(e))
            return None
        
        return properties