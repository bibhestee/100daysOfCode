#!/usr/bin/python
# -*- coding: utf-8 -*-

import csv
import re
import operator

# Initialize the dicts to save data

per_user = {}
error = {}

# Create the info and error pattern

info_pattern = r"ticky: INFO ([\w+\s]*)"
err_pattern = r"ticky: ERROR ([\w+\s']*)"
usr_pattern = r"[\W]+\(([\w.]*)\)$"

# Open the log file and parse the logs

with open('syslog.log') as file:
    logs = file.readlines()
    for log in logs:
        user = re.search(usr_pattern, log).group(1)
        info = re.search(info_pattern, log)
        err = re.search(err_pattern, log)
        if user not in per_user:
            per_user[user] = [0, 0]
        if info != None:
            msg = info.group(1)

      # Increment the number of info in per_user if log is info

            per_user[user][0] += 1
        elif err != None:
            msg = err.group(1)

      # Increment the number of error in per_user if log is error

            per_user[user][1] += 1

      # Add message to erro dictionary

            if msg in error:
                error[msg] += 1
                continue
            error[msg] = 1
errors = sorted(error.items(), key=operator.itemgetter(1), reverse=True)
errors.insert(0, ('Error', 'Count'))

# Store the error as csv

with open('error_message.csv', 'w') as file:
    f = csv.writer(file)
    for err in errors:
        f.writerow(err)

# Store the user as csv

per_user = sorted(per_user.items(), key=operator.itemgetter(0))
per_user.insert(0, ('Username', 'INFO', 'ERROR'))
with open('user_statistics.csv', 'w') as file:
    f = csv.writer(file)
    f.writerow([per_user[0][0], per_user[0][1], per_user[0][2]])
    for usr in per_user[1:]:
        f.writerow([usr[0], usr[1][0], usr[1][1]])
