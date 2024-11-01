# PR Pilot ✈️

**PR Pilot** is a terminal-based tool that helps you stay on top of incoming pull requests across multiple GitHub repositories by listing them in order from newest to oldest. Perfect for keeping an eye on your team's PR activity, PR Pilot makes it easy to navigate and open pull requests from a single interface.

## Features

- **Unified PR View:** See open pull requests from all configured repositories, sorted with the newest PRs at the top, allowing you to quickly review recent activity across multiple projects.
- **Navigate PRs:** Use keyboard arrows or `j/k` to move through PRs.
- **Open PRs:** Press `Enter` to open the selected PR in your default browser.
- **Quit:** Press `q` or `Ctrl+C` to exit the application.

## Installation

To install PR Pilot, run the following command in your terminal. This will automatically download and install the latest version:

```bash
curl -sL https://raw.githubusercontent.com/bjess9/pr-pilot/main/install.sh | bash
```

## Getting Started

1. **Configure Repositories**

    Configure the repositories you want to track by running:

    ```bash
    prpilot configure
    ```

    Follow the prompts to enter the list of repositories in `owner/repo` format, separated by commas. Your repository configuration will be saved to `~/.prpilot_config.yaml`. You can edit this file directly if you need to make further changes later.

2. **Run the Application**

    To start the application, use:

    ```bash
    prpilot
    ```

3. **Authentication**

   PR Pilot uses GitHub’s OAuth Device Flow to authenticate. When you run the app for the first time, you’ll be prompted to authenticate:

   - **Follow the Authentication Prompt:**  
     The application will display a URL and a one-time code. Open the URL in your browser, enter the provided code, and authorize the application.

   - **Token Storage:**  
     After authentication, your access token will be stored as follows, depending on your platform:

     - **macOS**: Stored securely in the Keychain.
     - **Windows**: Stored securely in the Windows Credential Manager.
     - **Linux**: Stored securely in the Secret Service API (used by GNOME Keyring or similar).
     - **WSL2**: ⚠️ **Stored as a plain text file in your home directory (`~/.prpilot_token`). This is not a secure storage method**, so take care to protect access to this file if security is a concern.

   With this approach, you won’t need to re-authenticate each time you run PR Pilot. However, on WSL2, **⚠️ the token is stored in plain text** and lacks the added security provided by other platforms' native credential managers.

## Usage

- **Track New PRs:** See incoming PRs in order of newest to oldest, so you can review recent work across repositories.
- **Navigate PRs:** Use `↑/↓` arrow keys or `j/k` to move up and down the list.
- **Open PR:** Press `Enter` to open the selected PR in your browser.
- **Quit:** Press `q` or `Ctrl+C` to exit the application.
