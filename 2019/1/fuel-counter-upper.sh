#!/bin/bash
fuel=0
for module in `cat input`; do
    out=$(($(echo $(($module / 3)) | awk '{print int($1)}')-2))
    fuel=$(($fuel + $out ))
done
echo $fuel
