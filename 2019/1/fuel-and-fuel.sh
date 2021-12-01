#!/bin/bash
fuel=0
for module in `cat input`; do
    echo $module
    out=$(($(echo $(($module / 3)) | awk '{print int($1)}')-2))
    fuel=$(($fuel + $out ))
    while true; do
      extra_fuel=$(($(echo $(($out / 3)) | awk '{print int($1)}')-2))
      if [ $extra_fuel -gt 0 ]; then
        fuel=$(($fuel + $extra_fuel))
        out=$extra_fuel
      else
        break
      fi
    done
done
echo $fuel
