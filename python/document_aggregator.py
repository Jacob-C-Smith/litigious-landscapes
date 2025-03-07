# Imports
import sys
import os
import csv
import json

ALLOTMENT_NO_COLUMN_INDEX            = 0
ALLOTMENT_NAME_COLUMN_INDEX          = 2
ALLOTMENT_DECISION_DATE_COLUMN_INDEX = 0

# Print usage 
def print_usage_and_abort ( exit_code : int ):

    # Print a usage message
    print(f"Usage: python3 {sys.argv[0]} input_file.csv")

    # Abort 
    sys.exit(exit_code)

# Body
def main( document_directory : str ):

    # Initialized data
    all_docs : dict = { }

    # 
    try:

        # List all documents in the documents directory
        for entry in os.listdir(document_directory):

            # Construct the absolute filee path
            full_path = os.path.join(document_directory, entry)

            # Error catch
            if os.path.isfile(full_path):

                # Open the file for reading
                with open(full_path, "r") as f:

                    # Store the file text
                    text : str = f.read()

                    # Catch empty files
                    if len(text) == 0:
                        continue

                    # Reflect the file into a dictionary
                    doc : dict = json.loads(text)

                    # Aggregate the document
                    all_docs[entry.split('.')[0]] = doc
    
    # Catch file not found
    except FileNotFoundError:
        pass
    
    # Open the aggregate file for writing
    with open("aggregate.json", 'w') as file:

        # Write the aggregated json document
        json.dump(all_docs, file, indent=4)

    # Done
    return

# Entry point
if __name__ == '__main__':

    # Argument check
    #if len(sys.argv) != 2:

    # Error
    #    print_usage()

    main( sys.argv[1] )
    #main( '../documents/' )