---
name: is-windows
description: Tell the AI that the current operating system is Windows, and provide PowerShell-based file operations and system checks.
author: FeiNiaoBF
version: 1.0.0
date_added: 2026-04-08
---

# Is Windows Skill

This skill allows the AI to recognize that the current operating system is Windows. It provides relevant information, instructions, and support specific to Windows users, including PowerShell scripts for file operations, checking Windows PATH, and other Windows-specific functionalities.

## Usage

When working on Windows systems, load this skill to ensure the AI understands the environment and can provide appropriate PowerShell commands.

### PowerShell File Operations

Use PowerShell for file and directory operations on Windows:

- **List files in a directory:**
  ```powershell
  Get-ChildItem -Path "C:\Path\To\Directory"
  ```

- **Create a new directory:**
  ```powershell
  New-Item -ItemType Directory -Path "C:\Path\To\NewDirectory"
  ```

- **Copy files:**
  ```powershell
  Copy-Item -Path "C:\Source\File.txt" -Destination "C:\Destination\File.txt"
  ```

- **Move files:**
  ```powershell
  Move-Item -Path "C:\Source\File.txt" -Destination "C:\Destination\File.txt"
  ```

- **Delete files:**
  ```powershell
  Remove-Item -Path "C:\Path\To\File.txt"
  ```

### Checking Windows PATH

To check and modify the Windows PATH environment variable:

- **View current PATH:**
  ```powershell
  $env:PATH
  ```

- **Add a directory to PATH (user level):**
  ```powershell
  $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
  $newPath = $currentPath + ";C:\New\Path"
  [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
  ```

- **Add a directory to PATH (system level, requires admin):**
  ```powershell
  $currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
  $newPath = $currentPath + ";C:\New\Path"
  [Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
  ```

### Other Windows-Specific Functions

- **Check Windows version:**
  ```powershell
  Get-ComputerInfo | Select-Object WindowsProductName, WindowsVersion
  ```

- **Get system information:**
  ```powershell
  systeminfo
  ```

- **Run as administrator (if needed):**
  Use `Start-Process` with `-Verb RunAs` for elevated privileges.

## Best Practices

- Always use absolute paths when possible to avoid issues with current directory.
- Quote paths that contain spaces: `"C:\Program Files\MyApp"`
- Use `Test-Path` to check if files or directories exist before operations.
- For batch operations, consider using pipelines and `ForEach-Object`.

## Examples

- To create a backup script: Combine `Copy-Item` with date formatting.
- To clean up temp files: Use `Get-ChildItem` with filters and `Remove-Item`.
- To set up development environment: Modify PATH and create necessary directories.
