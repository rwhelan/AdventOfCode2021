
raw_input = open('input', 'r').read()
raw_rows = raw_input.split('\n')
values = [int(i) for i in raw_rows]

increases = 0

for i, v in enumerate(values):
    if i == 0:
        continue

    if values[i-1] < v:
        increases += 1

print(increases)