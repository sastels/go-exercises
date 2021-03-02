import os
import sys

for dir in os.listdir('.'):
    if os.path.isdir(dir) and not dir.startswith('.'):
        retval = os.system("cd {} && go test".format(dir)) != 0
        if retval != 0:
            sys.exit(retval)
