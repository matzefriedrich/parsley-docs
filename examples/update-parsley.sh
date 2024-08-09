#!/bin/bash

DEPENDENCY="github.com/matzefriedrich/parsley"

for project_dir in $(find . -mindepth 1 -maxdepth 1 -type d); do
    echo "Processing $project_dir"
    
    if [ -f "$project_dir/go.mod" ]; then
        echo "Updating dependency $DEPENDENCY in $project_dir"
        
        cd "$project_dir"
        
        go get -u $DEPENDENCY
        go mod tidy
        
        # Find all main.go files within cmd/* subdirectories
        for main_file in $(find ./cmd -mindepth 2 -type f -name main.go); do
            # Get the directory containing the main.go file
            app_dir=$(dirname "$main_file")
            
            echo "Building app in $app_dir"
            
            # Build the application
            if go build -o "$app_dir/app" "$app_dir/main.go"; then
                echo "Build successful for $app_dir"
                rm "$app_dir/app"
            else
                echo "Build failed for $app_dir"
                # Optionally, exit if a build fails
                # exit 1
            fi
        done
        
        cd ..
    fi
done

echo "Dependency update and build process complete."
