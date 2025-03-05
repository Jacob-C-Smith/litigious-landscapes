cd "$1"
ls -1 | 
awk '{print substr($1,1,4) " " substr($1,5,2) " " substr($1,7,2)}' |
sort |
uniq | 
tr ' ' '_'
cd - > /dev/null