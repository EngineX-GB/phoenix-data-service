import os
import requests

def list_text_files(directory):
    # List all .txt files in the directory and sort them alphabetically
    return sorted([f for f in os.listdir(directory) if f.endswith('.txt')])

def upload_file(file_path):
    url = "http://localhost:8080/feed"
    
    # Prepare the file to be uploaded
    with open(file_path, 'rb') as file:
        files = {'myFile': file}
        
        # Send a POST request with the file as multipart form-data
        response = requests.post(url, files=files)
        
        # Check the response status
        if response.status_code == 200:
            print(f"Successfully uploaded {file_path}")
        else:
            print(f"Failed to upload {file_path}. Status code: {response.status_code}")

def main(base_directory):
    # Walk through the directory structure
    for root, dirs, files in os.walk(base_directory):
        # Sort the directories and files alphabetically/numerically
        dirs.sort()
        files.sort()

        # Only process .txt files
        text_files = [f for f in files if f.endswith('.txt')]

        if not text_files:
            continue  # Skip if no text files found in the current directory

        # For each text file, upload it
        for text_file in text_files:
            file_path = os.path.join(root, text_file)
            upload_file(file_path)

if __name__ == "__main__":
    # Specify the base directory containing the subfolders (e.g., "2022-01-23", "2022-01-24", etc.)
    base_directory = "C:/Users/Dell/Documents/Development/code/python/phoenix-mobile-connector/feeds"
    
    # Run the main function to recursively upload text files
    main(base_directory)