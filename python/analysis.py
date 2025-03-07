import json
import sys
docs : list = None
rc : list = []
with open('aggregate.json', 'r') as f:
    docs = json.load(f)

for d in docs:
    if docs[d]['group'] != sys.argv[1]:
        continue
    for e in docs[d]['resource conditions']:
        for f in docs[d]['resource conditions'][e]:
            for g in docs[d]['resource conditions'][e][f]:
                print(f'{g}')

print(rc)