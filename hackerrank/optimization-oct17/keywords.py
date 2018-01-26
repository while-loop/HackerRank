#!/bin/python

MAXVAL = 1e9


def minimumLength(text, keys):
    text += " "
    ans = MAXVAL
    key_idx = {}

    word = ""
    for pos in xrange(len(text)):
        # continue until end of word is found
        if text[pos] != ' ':
            word += text[pos]
            continue

        if word in keys:
            key_idx[word] = pos - len(word)

            if len(keys) == len(key_idx):  # we've found all occurrences!
                ans = min(ans, pos - min(key_idx.values()))
        word = ""

    if ans == MAXVAL:
        ans = -1
    return ans


text = raw_input()
keyWords = int(raw_input())
keys = []
for i in xrange(keyWords):
    keys.append(raw_input())
print(minimumLength(text, keys))
