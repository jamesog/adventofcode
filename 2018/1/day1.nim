import strutils
import system
import tables

proc changeFrequency(freq: int, shift: string): int =
    var shiftInt = strutils.parseInt(shift)
    return freq + shiftInt

proc partOne(input: seq[string]): int =
    var freq = 0
    for shift in input:
        freq = changeFrequency(freq, shift)
    return freq

# This implementation is _really_ slow. Would be good to learn why.
proc partTwo(input: seq[string]): int =
    var freq = 0
    var seenFreqs = initTable[int, int]()
    while true:
        for shift in input:
            freq = changeFrequency(freq, shift)
            if seenFreqs.hasKey(freq):
                seenFreqs[freq] += 1
            else:
                seenFreqs[freq] = 1
            if seenFreqs[freq] == 2:
                return freq

var input: seq[string]
for line in stdin.lines:
    input.add(line)

echo partOne(input)
echo partTwo(input)
