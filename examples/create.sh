#!/bin/bash

go build -o pdfTest main.go

for FORMAT in $(ls examples/inputs/); do
    for COMPRESSION in NONE LOWEST LOW MEDIUM HIGH HIGHEST; do
        for MARGIN in NONE MIN MEDIUM MAX; do
            for MODE in P L; do
                echo "Creating document: $FORMAT, Compression $COMPRESSION, Margin $MARGIN, Orientation $MODE"
                ./pdfTest create examples/inputs/$FORMAT/*.png examples/outputs/$FORMAT$MODE-$COMPRESSION-$MARGIN.pdf -c $COMPRESSION -m $MARGIN -o $MODE
            done
        done    
    done
done
rm pdfTest