#!/bin/sh
sendip -p ipv4 -is localhost -p udp -us 5070 -ud 5060 -d $1 -v localhost