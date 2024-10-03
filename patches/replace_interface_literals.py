#!/usr/bin/env python

import os

def replace_string_in_file(file_path, old_string, new_string):
    print(f"Checking file {file_path}")
    # Read the content of the file
    with open(file_path, 'r', encoding='utf-8') as file:
        file_data = file.read()

    # Replace the target string
    new_data = file_data.replace(old_string, new_string)

    # Write the file out again
    with open(file_path, 'w', encoding='utf-8') as file:
        file.write(new_data)

def parse_and_replace_in_directory(directory, old_string, new_string):
    # Walk through the directory
    for foldername, subfolders, filenames in os.walk(directory):
        for filename in filenames:
            file_path = os.path.join(foldername, filename)
            if filename.endswith('.go'):
                replace_string_in_file(file_path, old_string, new_string)
                print(f"Processed {file_path}")

if __name__ == "__main__":
    parse_and_replace_in_directory("pkg/enterprise", "return interface{}{}", "return nil")
