#!/bin/bash

out_folders=(pkg)
golang_out_folder=pkg
app_abs_path=/app

function print_error_and_exit() {
    echo "❌ Error occurred. Printing tree and exiting..."
    tree -aC ./pkg/ ./internal/
    exit 1
}

echo "🧹 Cleaning up..."
rm -rf "${out_folders[@]}"/*
mkdir -p "${out_folders[@]}"

echo "🚀 Compiling the project..."
MAIN_FOLDER="internal"
list_proto_files=()

for folder in $(find "$MAIN_FOLDER" -type d); do
    if [[ $folder != $MAIN_FOLDER ]]; then
        list_proto_files+=($(find "$folder" -type f -name "*.proto"))
    fi
done

echo "📜 List of proto files:"
for file in "${list_proto_files[@]}"; do
    echo $file
done

echo "🌀 Compiling Golang..."
cd "$app_abs_path/$MAIN_FOLDER"
find . -name \*.proto -exec protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {} \;
tree -L 3 || print_error_and_exit

echo "👨🏽‍🔧 Moving Golang files to output folder..."
for file in $(find . -name \*.pb.go); do
    file_name=$(basename "$file")
    current_folder_name=$(dirname "$file" | xargs basename)
    file_name_without_extension=${file_name%.*}
    
    mkdir -p "$app_abs_path/$golang_out_folder/$current_folder_name"
    mv -v "$file" "$app_abs_path/$golang_out_folder/$current_folder_name/$file_name_without_extension.go"
done
tree -L 3 || print_error_and_exit

echo "⭐ Finished compiling the project ⭐"
tree -aC ./pkg/ ./internal/
echo "🌐 Pipeline design: Golang"
