import json
import sys
import logging
import pymysql
import pymysql.cursors


#rds settings
rds_host  = "maindb.czpld8fke1ht.us-east-1.rds.amazonaws.com"
name = "username"
password = "password"
db_name = "dbnew"
newDatabaseName = "testdb0"

logger = logging.getLogger()
logger.setLevel(logging.INFO)

try:
    # conn = pymysql.connect(host=rds_host, user=name, passwd=password, db=db_name, connect_timeout=5)
    conn = pymysql.connect(host=rds_host, user=name, passwd=password, connect_timeout=5)

except pymysql.MySQLError as e:
    logger.error("ERROR: Unexpected error: Could not connect to MySQL instance.")
    logger.error(e)
    sys.exit()

logger.info("SUCCESS: Connection to RDS MySQL instance succeeded")

def lambda_handler(event, context):
    try:
        # Create a cursor object
        cursor = conn.cursor()

        # SQL Statement to create a database
        sqlStatement    = "CREATE DATABASE "+newDatabaseName  

        # Execute the create database SQL statment through the cursor instance
        cursor.execute(sqlStatement)

        # # SQL query string
        # sqlQuery    = "SHOW DATABASES"

        # # Execute the sqlQuery
        # cursor.execute(sqlQuery)

        # #Fetch all the rows
        # databaseList    = cursor.fetchall()

        # for datatbase in databaseList:
        #     print(datatbase)

    except Exception as e:
        print("Exeception occured:{}".format(e))
        
    finally:
        conn.close()

    message = 'Hello {} {}!'.format(event['key1'], event['key2'])  
    return { 
        'message' : message
    }