from collections import Counter

N = int(raw_input())
for x in range(N):
    raw_input()
    l = map(int, raw_input().split())
    if Counter(l)[min(l)] % 2 == 0:
        print "Unlucky"
    else:
        print "Lucky"
