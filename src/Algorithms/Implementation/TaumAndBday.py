#!/bin/python

import sys

if __name__ == '__main__':
    t = int(raw_input().strip())

    for a0 in xrange(t):
        b, w = raw_input().strip().split(' ')
        b, w = [long(b), long(w)]
        x, y, z = raw_input().strip().split(' ')
        x, y, z = [long(x), long(y), long(z)]

        print "done"
