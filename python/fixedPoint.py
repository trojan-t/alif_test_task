def FixedPoint(arr):
    for i in range(len(arr)):
        if arr[i] is i:
            return i
    return False

arr = [-6, 0, 2, 40]

print(FixedPoint(arr))