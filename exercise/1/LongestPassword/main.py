import re

def solution(S):
    re_password = re.compile("[a-zA-Z0-9]+")
    re_letter = re.compile("[a-zA-Z]+")
    re_digit = re.compile("[0-9]+")

    result = -1
    arr = S.split(' ')
    for s in arr:
        letter_num = len(''.join(re_letter.findall(s)))
        digit_num = len(''.join(re_digit.findall(s)))
        if letter_num % 2 == 0 and digit_num % 2 == 1 and re_password.fullmatch(s) != None and result < len(s):
            result = len(s)
    return result