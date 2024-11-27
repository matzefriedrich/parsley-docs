#!/bin/bash

# Array of package names
packages=("bootstrap" "features" "registration" "resolving" "types")  # Add more package names as needed

# Base directory for packages and output
pkg_base_dir="../parsley/pkg"
output_dir="./library/pkg"

for package in "${packages[@]}"; do
    echo "Generating documentation for package: $package"

    # Construct the package path and output file
    package_path="$pkg_base_dir/$package"
    output_file="$output_dir/$package.md"

    # Invoke gomarkdoc with the specified template files
    gomarkdoc "$package_path" \
        --template-file package=./gomarkdoc/package.gotxt \
        --template-file file=./gomarkdoc/file.gotxt > "$output_file"

    # Check if gomarkdoc was successful
    if [[ $? -eq 0 ]]; then
        echo "Documentation for $package generated successfully at $output_file"
    else
        echo "Failed to generate documentation for $package"
    fi
done
