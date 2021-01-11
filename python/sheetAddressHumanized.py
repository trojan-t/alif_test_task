def sheetAddressHumanized(index):
    alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    alphabetLen = len(alphabet)
    reversedAddress = []

    while index > 0:
        letter = alphabet[(index-1) % alphabetLen]
        reversedAddress.append(letter)
        index = (index-1) // alphabetLen

    return ''.join(reversed(reversedAddress))

print(sheetAddressHumanized(27))