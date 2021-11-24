#!/bin/bash
[ -a ./chart/values.yaml ] && sed -i -e 's/dev.ubivius.tk/ubivius.tk/g' ./chart/values.yaml