# Heading
printf "\033[94m╭─────────────────────────────────╮\033[0m\n"
printf "\033[94m│ Litigious Landscapes Fragmenter │\033[0m\n"
printf "\033[94m╰─────────────────────────────────╯\033[0m\n\n"

# Cleanup
mkdir -p ../fragments/ &> /dev/null
rm -rf ../fragments/* &> /dev/null

../shell/raw_data_to_plain_text.sh ./

# Final report
printf "\033[96m\033[1m\033[4mFinal report\033[0m\n"
tree ../fragments | tail -n 1 | awk '{printf "Analyzed " $1 " documents; Yielded " $3 " fragments; Time elapsed: "}'