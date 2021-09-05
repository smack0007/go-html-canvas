#!/bin/bash
echo "Vetting..."

for d in ./src/*/ ; do 
    if test -f "($d)go.mod"; then
        (cd "$d" && go vet);
    fi
done