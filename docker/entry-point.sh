#!/bin/sh

# echo "running entryppoint"

awk -v line=$(awk '/\IPv4 local connections\>/ {a=$0;c=NR;next} END{print c}' /var/lib/postgresql/data/pg_hba.conf) '{if( NR == line + 1) { {orig=$0} $5="md5"; OFS= "";print"# **FOLLOWING ENTRY AUTOMATICALLY MODIFIED. ORIGINAL IS -- ",orig, " - **\n",$orig} else print $0 }' /
var/lib/postgresql/data/pg_hba.conf