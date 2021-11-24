#!/bin/sh
check_match=`cat ./chart/values.yaml | grep "match"`;
if [ -z "$check_match" ];
    then
    check_match=`cat ./chart/values.yaml | grep "ubivius.tk"`;
    if [ -n "$check_match" ];
        then
        main_url_name=`yq e '.ingress.hosts[].host' ./chart/values.yaml | sed "s/dev.//g"`;
        yq eval -i ".ingress.hosts[].host = \"$main_url_name\"" ./chart/values.yaml;
    fi
else
    main_url_name=`yq e '.ingressRoute.match' ./chart/values.yaml | sed "s/dev.//g"`;
    yq eval -i ".ingressRoute.match = \"$main_url_name\"" ./chart/values.yaml;
fi
