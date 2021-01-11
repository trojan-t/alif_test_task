import math

def EgyptianFractions(top, bottom):
    bottomList = []

    while top != 0:
        x = math.ceil(bottom / top)
        bottomList.append(x)
        top = x * top - bottom
        bottom *= x
    for i in range(len(bottomList)):
        if i != len(bottomList) - 1:
            print("1/{0} +".format(bottomList[i]), end = " ")
        else:
            print(" 1/{0}".format(bottomList[i]), end = " ")

EgyptianFractions(4, 13)