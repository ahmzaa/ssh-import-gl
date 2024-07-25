# ssh-import-gl

## Overview
ssh-import-gl is a Go-based tool designed to import SSH keys from GitLab. This
repository contains the source code, allowing users to seamlessly import SSH
keys into their .ssh/authorized_keys file from GitLab.

## Features
- Imports users' public ssh keys from Gitlab.
- Options to either append or to overwrite the existing keys in the .ssh/authorized_keys file.

## Motivation
Many Linux/Unix systems provide commands to import SSH keys from GitHub, but
there is no native importer for GitLab. After moving from GitHub to GitLab
following Microsoft's acquisition of GitHub, I needed a way to keep my SSH keys
updated without relying on GitHub. This tool aims to fill that gap by allowing
users to import SSH keys directly from GitLab.

## Benefits
- Allows for a complete migration from GitHub to GitLab without losing key management functionality.
- Centralizes the storage of public keys, eliminating the need to update keys in multiple locations.

## Drawbacks
- Currently, no known drawbacks.

## Implementation
- **Programming Language:** Golang
- **API:** Gitlab API

## To-Do List
- [x] Gather list of keys from gitlab
- [ ] Import keys into the authorized_keys file
- [ ] Check if keys already exist in the file

## Notes
- Requires a GitLab account with public keys added.

# Getting Started
## Prerequisites
- Go installed on your machine.
- A Gitlab account with public SSH keys added.

## Installation
1. Clone the repository
git clone https://gitlab.com/ahmza/ssh-import-gl

2. Navigate to the project directory
cd ssh-import-gl

3. Build the project:
go build -o ssh-import-gl

## Usage
Run the ssh-import-gl tool with the necessary options:
ssh-import-gl ahmza -a

Options:
- -a: append
- -o: overwrite
