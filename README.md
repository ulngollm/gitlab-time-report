# GitLab Time Report

This project is a command-line tool designed to interact with the GitLab API to generate reports on time spent on issues and merge requests within a specific GitLab project.

## Features

- **Stats Mode**: Calculate the total time spent on issues and merge requests.
- **Report Mode**: Generate a detailed report of time entries.

## Requirements

- GitLab personal access token with appropriate permissions

## Installation

1. Clone the repository:
2. Build the project:
   ```bash
   go build -o gitlab-time-report
   ```

## Usage

The tool requires several parameters to be passed either via command-line arguments or environment variables:

- `--mode` or `-m`: The mode of operation, either `stats` or `report`.
- `--token`: Your GitLab personal access token.
- `--host`: The base URL of the GitLab API.
- `--labels`: Labels to filter issues and merge requests (default is `any`).
- `--project`: The ID of the GitLab project.

Example command:

```bash
./gitlab-time-report --mode stats --token YOUR_TOKEN --host https://gitlab.com --project YOUR_PROJECT_ID
```

## Environment Variables

You can also set the following environment variables instead of passing them as command-line arguments:

- `TOKEN`: Your GitLab personal access token.
- `API_HOST`: The base URL of the GitLab API.
- `LABELS`: Labels to filter issues and merge requests.
- `PROJECT_ID`: The ID of the GitLab project.

## Data Source

The script fetches data from the GitLab API using the provided personal access token and project ID. It retrieves information about time spent on issues and merge requests, filtered by the specified labels.
