#!/usr/bin/python3.7

import RPi.GPIO as GPIO
from mfrc522 import SimpleMFRC522

import mysql.connector

# TODO: version change on update
version = "0.1"

print("initializing...")
print("This is version " + version + ".")

# property assignment
db_address = "127.0.0.1"
db_port = "3306"
db_username = "stempeluhr"
db_password = "Rx8723hm95Wqbnk324zx"
db_database = "stempeluhr"

reader = SimpleMFRC522()


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
    result = cursor.fetchall()
    print(result)
    cursor.close()
    cnx.close()
    return