#!/bin/bash
timedatectl set-ntp false
data=$(< /etc/default/ntp)
echo $data
echo "exit" > /etc/default/ntp
echo $data >> /etc/default/ntp
timedatectl status
