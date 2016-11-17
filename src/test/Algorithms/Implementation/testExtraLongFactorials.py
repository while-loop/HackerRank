#!/bin/python

import unittest
import math

from Algorithms.Implementation.ExtraLongFactorials import fact


class TestExtraLongFactorials(unittest.TestCase):
    def test_factorials(self):
        MIN = 1
        MAX = 100

        for i in range(MIN, MAX + 1):
            self.assertEqual(fact(i), math.factorial(i))


if __name__ == '__main__':
    unittest.main()
