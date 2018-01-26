#!/bin/python
def collatzSequenceSum(T, A, B):
    dp = {}

    def collatzSequenceLen(k):
        if k <= 1:
            return k

        if k in dp:
            return dp[k]

        if k % 2 == 0:
            dp[k] = 1 + collatzSequenceLen(k / 2)
        else:
            dp[k] = 1 + collatzSequenceLen(3 * k + 1)

        return dp[k]

    n = 0
    result = 0
    for _ in xrange(T):
        n = (A * n + B) % 5003
        best_len = 0
        best_num = 0
        for k in xrange(1, n + 1):
            cur_len = collatzSequenceLen(k)
            if cur_len >= best_len:
                best_len = cur_len
                best_num = k
        result += best_num
    return result


if __name__ == "__main__":
    T, A, B = raw_input().strip().split(' ')
    T, A, B = [int(T), int(A), int(B)]
    result = collatzSequenceSum(T, A, B)
    print result
