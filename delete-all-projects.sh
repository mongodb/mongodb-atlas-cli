#!/bin/bash

# Script to delete all Atlas projects
# Usage: ./scripts/delete-all-projects.sh [--dry-run] [--atlas-cli-path=/path/to/atlas]

set -e

# Default values
DRY_RUN=false
ATLAS_CLI_PATH="./bin/atlas"
FORCE_DELETE=false

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        --atlas-cli-path=*)
            ATLAS_CLI_PATH="${1#*=}"
            shift
            ;;
        --force)
            FORCE_DELETE=true
            shift
            ;;
        -h|--help)
            echo "Usage: $0 [OPTIONS]"
            echo ""
            echo "Options:"
            echo "  --dry-run              Show what would be deleted without actually deleting"
            echo "  --atlas-cli-path=PATH  Path to atlas CLI binary (default: ./bin/atlas)"
            echo "  --force                Skip individual confirmations (use with caution!)"
            echo "  -h, --help             Show this help message"
            echo ""
            echo "Examples:"
            echo "  $0 --dry-run                    # Preview what would be deleted"
            echo "  $0 --atlas-cli-path=atlas       # Use atlas from PATH"
            echo "  $0 --force                      # Delete all without individual confirmations"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            echo "Use --help for usage information"
            exit 1
            ;;
    esac
done

# Check if atlas CLI exists
if [[ ! -x "$ATLAS_CLI_PATH" ]]; then
    echo "Error: Atlas CLI not found at $ATLAS_CLI_PATH"
    echo "Please ensure the Atlas CLI is built or specify the correct path with --atlas-cli-path"
    exit 1
fi

echo "Using Atlas CLI at: $ATLAS_CLI_PATH"
echo ""

# Get the list of projects
echo "Fetching project list..."
PROJECT_LIST_OUTPUT=$($ATLAS_CLI_PATH projects list 2>/dev/null | grep -v "A new version of atlascli" | grep -v "To upgrade" | grep -v "To disable this alert" | grep -v "^$")

# Extract project IDs (first column, excluding header)
PROJECT_IDS=$(echo "$PROJECT_LIST_OUTPUT" | tail -n +2 | awk '{print $1}' | grep -E '^[0-9a-f]{24}$')

# Count projects
PROJECT_COUNT=$(echo "$PROJECT_IDS" | wc -l | tr -d ' ')

if [[ -z "$PROJECT_IDS" || "$PROJECT_COUNT" -eq 0 ]]; then
    echo "No projects found to delete."
    exit 0
fi

echo "Found $PROJECT_COUNT projects to delete:"
echo ""

# Display projects in a nice format
echo "$PROJECT_LIST_OUTPUT" | head -1  # Header
echo "$PROJECT_LIST_OUTPUT" | tail -n +2
echo ""

if [[ "$DRY_RUN" == "true" ]]; then
    echo "DRY RUN: The following projects would be deleted:"
    for project_id in $PROJECT_IDS; do
        project_name=$(echo "$PROJECT_LIST_OUTPUT" | tail -n +2 | grep "^$project_id" | awk '{print $2}')
        echo "  - $project_id ($project_name)"
    done
    echo ""
    echo "To actually delete these projects, run without --dry-run"
    exit 0
fi

# Final confirmation
if [[ "$FORCE_DELETE" != "true" ]]; then
    echo "⚠️  WARNING: This will delete ALL $PROJECT_COUNT projects listed above!"
    echo "This action cannot be undone."
    echo ""
    read -p "Are you absolutely sure you want to delete all these projects? (type 'DELETE ALL' to confirm): " confirmation
    
    if [[ "$confirmation" != "DELETE ALL" ]]; then
        echo "Operation cancelled."
        exit 0
    fi
fi

echo ""
echo "Starting deletion process..."
echo ""

# Delete each project
DELETED_COUNT=0
FAILED_COUNT=0

for project_id in $PROJECT_IDS; do
    project_name=$(echo "$PROJECT_LIST_OUTPUT" | tail -n +2 | grep "^$project_id" | awk '{print $2}')
    
    echo "Deleting project: $project_id ($project_name)"
    
    if [[ "$FORCE_DELETE" == "true" ]]; then
        # Use expect to automatically answer "Yes" to the confirmation prompt
        if command -v expect >/dev/null 2>&1; then
            expect_script=$(cat << 'EOF'
spawn {*}$argv
expect "Are you sure you want to delete:" { send "Yes\r" }
expect eof
EOF
            )
            if echo "$expect_script" | expect -f - "$ATLAS_CLI_PATH" projects delete "$project_id" >/dev/null 2>&1; then
                echo "  ✅ Successfully deleted $project_id"
                ((DELETED_COUNT++))
            else
                echo "  ❌ Failed to delete $project_id"
                ((FAILED_COUNT++))
            fi
        else
            echo "  ⚠️  expect command not found. Skipping $project_id"
            echo "     Install expect with: brew install expect (macOS) or apt-get install expect (Ubuntu)"
            ((FAILED_COUNT++))
        fi
    else
        # Interactive mode - let user confirm each deletion
        if $ATLAS_CLI_PATH projects delete "$project_id"; then
            echo "  ✅ Successfully deleted $project_id"
            ((DELETED_COUNT++))
        else
            echo "  ❌ Failed to delete $project_id"
            ((FAILED_COUNT++))
        fi
    fi
    
    echo ""
done

echo "Deletion summary:"
echo "  ✅ Successfully deleted: $DELETED_COUNT projects"
if [[ $FAILED_COUNT -gt 0 ]]; then
    echo "  ❌ Failed to delete: $FAILED_COUNT projects"
fi
echo ""

if [[ $FAILED_COUNT -gt 0 ]]; then
    echo "Some projects could not be deleted. Please check the errors above."
    exit 1
else
    echo "All projects have been successfully deleted! 🎉"
fi
