cd "$1"
rm "../../../../fragment_log"
NUM_DOCS=$(ls -1 | wc -l)
COUNTER=1
printf " - Found $NUM_DOCS document(s)\n"
ls -1 | 
cut -d" " -f 3- |
cut -d"_" -f 2- |
sed 's/ /\ /g' |
cut -d"." -f 1 |
while IFS= read -r line; do
    touch "../../../../resources/auto_plain_text/$line.txt"
    ls -1 | grep "$line" |
    while IFS= read -r in_file; do
        printf "\r                                                                                "
        printf "\r - [$COUNTER/$NUM_DOCS] \"$in_file\"\r"
        pdf2txt "$(echo $in_file)" | awk '{$1=$1;print}' | uniq > "../../../../resources/auto_plain_text/$line.txt"
        mkdir "../../../../resources/fragments/$line" 1> /dev/null 2> /dev/null
        ../../../../build/ll_analyzer -in "../../../../resources/auto_plain_text/$line.txt" -out "../../../../resources/fragments/$line" >> "../../../../fragment_log"
        pdfinfo "$(echo $in_file)" > "../../../../resources/fragments/$line/info.txt"
    done
    COUNTER=$((COUNTER + 1))
done
printf "\n - Analyzed $NUM_DOCS documents\n\n"