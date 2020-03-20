#!/usr/bin/python3.7

import sys
import RPi.GPIO as GPIO
import mysql.connector
import configparser
from mfrc522 import SimpleMFRC522

# TODO: create powerloop as operational supervisor
# TODO: uid muss in 'ification' definiert sein um TAG zuzuweisen

# basic initialisation of used fVariables
reader = SimpleMFRC522()
config = configparser.ConfigParser()
config.read('config.ini')

# reading & applying config
version = config['INITIAL']['version']
debug = config['INITIAL'].getboolean('debug')

db_address = config['DATABASE']['address']
db_port = config['DATABASE']['port']
db_username = config['DATABASE']['username']
db_password = config['DATABASE']['password']
db_database = config['DATABASE']['database']

if debug: print("INITIALIZED")
print("This is version " + version + ".")


# FUNCTIONS BEING DEFINED HERE
# FOR READING TAG
def read_nfc():
    print("Hold your NFC-Tag close the reader please.")
    nfc_id, text = reader.read()
    print("The read id is: " + str(nfc_id))
    GPIO.cleanup()
    return str(nfc_id)


# FOR READING THE DATABASE (RETURNING UserID)
def check_database(tag_id):
    cnx = mysql.connector.connect(user=db_username, host=db_address, password=db_password, database=db_database)
    query = "SELECT uid FROM ident WHERE nfc_id = " + tag_id
    cursor = cnx.cursor()
    cursor.execute(query)
    result = cursor.fetchone()
    if debug: print(result[0])
    cursor.close()
    cnx.close()
    return result[0]


# FOR ADDING TAGs TO THE DATABASE
def add_tag_to_db(tag_id):
    print("This tag is not yet known to the database.")
    print("Would you like to create a new entry?")
    answer_to_question = input("type y/N: ")

    if answer_to_question.upper() == "Y":
        uid_to_write = input("Please specify the persons UserID: ")
        print("The UserID you gave is: " + uid_to_write + ".")
        confirm_assignment = input("Do you want to assign this UserID to TAG-" + tag_id + "? Type y/N: ")

        if confirm_assignment.upper() == "Y":
            if debug: print("connecting to database")
            cnx = mysql.connector.connect(user=db_username, host=db_address, password=db_password, database=db_database)
            if debug: print("connected")
            query = "INSERT INTO `ident` (`uid`, `nfc_id`) VALUES ('" + uid_to_write + "', '" + tag_id + "');"
            if debug: print("assigning...")
            cursor = cnx.cursor()

            try:
                cursor.execute(query)
                cnx.commit()
            except mysql.connector.Error as err:
                print("ERROR:")
                print(err)
                sys.exit(1)

            print("SUCCESS")
            cursor.close()
            cnx.close()
            if debug: print("connection closed")

        else:
            print("CANCELED")

    else:
        if debug: print("You can come back any time to setup the RFID-TAG.")


# FOR 'STEMPELN' and DIRECTION ANALYSIS
def stempeln(userid):
    query_get_direction = "SELECT `isHere` FROM `ification` WHERE `uid` = " + userid.str() + ";"
    cnx = mysql.connector.connect(user=db_username, host=db_address, password=db_password, database=db_database)
    cursor = cnx.cursor()
    cursor.execute(query_get_direction)
    currently_here = cursor.fetchone()
    new_direction = 1
    if currently_here[0]:
        new_direction = 0
    query_setnew_direction = "UPDATE `ification` SET `isHere` = " + new_direction.str() + " WHERE `uid` = " + userid.str() + " ;"
    cursor.execute(query_setnew_direction)
    cnx.commit()
    query_log = "INSERT INTO `log` (`datetime`, `uid`, `cameIn`) VALUES ('current_timestamp().000000', '" + userid.str() + "', '" + new_direction.str() + "');"
    cursor.execute(query_log)
    cnx.commit()
    cursor.close()
    cnx.close()


# ALWAYS CLEANUP GPIO AFTER USE!
GPIO.cleanup()
