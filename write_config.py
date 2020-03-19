#!/usr/bin/python3.7
import configparser

config = configparser.ConfigParser()

config['INITIAL'] = {"version": "0.1", "debug": "True"}
config['DATABASE'] = {"address": "127.0.0.1", "port": "3306", "username": "stempeluhr", "password": "Rx8723hm95Wqbnk324zx", "database": "stempeluhr"}

with open('config.ini', 'w') as configfile:
    config.write(configfile)

