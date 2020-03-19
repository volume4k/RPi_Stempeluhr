#!/usr/bin/python3.7

import sys
import RPi.GPIO as GPIO
import mysql.connector
import configparser
from mfrc522 import SimpleMFRC522

# TODO: create powerloop as operational supervisor
# TODO: cleanup

reader = SimpleMFRC522()
config = configparser.ConfigParser()
config.read('config.ini')
version = config['INITIAL']['version']
debug = False

if debug: print("initializing...")
print("This is version " + version + ".")

# property assignment
db_address = "127.0.0.1"
db_port = "3306"
db_username = "stempeluhr"
db_password = "Rx8723hm95Wqbnk324zx"
db_database = "stempeluhr"


# functions
def read_nfc():
    print("Hold your NFC-Tag close the reader please.")
    nfc_id, text = reader.read()
    print("The read id is: " + str(nfc_id))
    GPIO.cleanup()
    return str(nfc_id)


def check_database(tag_id):
    cnx = mysql.connector.connect(user=db_username, host=db_address, password=db_password, database=db_database)
    query = "SELECT uid FROM ident WHERE nfc_id = " + tag_id
    cursor = cnx.cursor()
    cursor.execute(query)
    result = cursor.fetchone()
    print(result[0])
    cursor.close()
    cnx.close()
    return result[0]


def add_tag_to_db(tag_id):
    print("This tag is not yet known to the database.")
    print("Would you like to create a new entry?")
    answer_to_question = input("type y/N: ")

    if answer_to_question.upper() == "Y":
        uid_to_write = input("Please specify the persons UserID: ")
        print("The UserID you gave is: " + uid_to_write + ".")
        confirm_assignment = input("Do you want to assign this UserID to TAG-" + tag_id + "? Type y/N: ")

        if confirm_assignment.upper() == "Y":
            print("connecting to database")
            cnx = mysql.connector.connect(user=db_username, host=db_address, password=db_password, database=db_database)
            print("connected")
            query = "INSERT INTO `ident` (`uid`, `nfc_id`) VALUES ('" + uid_to_write + "', '" + tag_id + "');"
            print("assigning...")
            cursor = cnx.cursor()

            try:
                cursor.execute(query)
                cnx.commit()
            except mysql.connector.Error as err:
                print(err)
                sys.exit(1)

            print("assigned")
            cursor.close()
            cnx.close()
            print("connection closed")

        else:
            print("operation canceled. please start again.")

    else:
        print("You can come back any time to setup the RFID-TAG.")
