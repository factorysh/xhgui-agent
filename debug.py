import sys

import bson


b = bson.loads(open(sys.argv[1], 'rb').read())
print(b)