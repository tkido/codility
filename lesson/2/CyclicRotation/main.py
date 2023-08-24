
def solution(A, K):
    # Implement your solution here
    length = len(A)
    if length == 0 or K == 0:
        return A
    if K > length:
        K = K % length
    if K == 0:
        return A
    return A[-K:]+A[:length-K]