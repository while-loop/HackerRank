#!/usr/bin/env python
import sys


def test():
    cases = [
        ([3, 2, 1, 3], 2),
        ([3, 5, 4, 5], 2)
    ]
    for c in cases:
        max = castleTowers(len(c[0]), c[0])
        if c[1] != max:
            print "got:", max, "want:", c[1], c[0]
    sys.exit()


def castleTowers(n, ar):
    counts = 0
    max = -1
    for h in ar:
        if h > max:
            max = h
            counts = 1
        elif h == max:
            counts += 1
    return counts


if __name__ == "__main__":
    n = int(raw_input().strip())
    ar = map(int, raw_input().strip().split(' '))
    result = castleTowers(n, ar)
    print result
