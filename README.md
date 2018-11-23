# GoLang


Reads a configuration file in JSON format, gets all required configuration
- Copies all directories mentioned in the configuration to a temporary folder.
- Dumps and backs up all the MongoDB databases mentioned in the configuration to the same temporary folder.
- The resulting compressed file should be uploaded to any object storage of your choice, the details for which should be taken from the configuration file.
- The configuration file for the application should be passed as a command line argument
