#!/bin/python

import sys


def fact(n):
    ret = 1
    for i in range(2, n+1):
        ret *= i
    return ret


if __name__ == '__main__':
    n = int(raw_input().strip())
    assert 1 <= n <= 100

    print fact(n)