# Contributing to cobra-gin-starter

We love your input! ❤️
We want to make contributing to this project as easy and transparent as possible — whether it's:

- Reporting a bug
- Discussing current or future improvements
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## We Develop with GitHub

We use GitHub to host code, track issues, and manage pull requests.

## Pull Requests Are Welcome

Pull requests are the best way to propose changes to the codebase. We actively welcome them!

1. **Fork the repo** and create your branch from `main`.
2. **Install dependencies:**
    ```bash
    go mod tidy
    bun install
    templ generate && bun run build:css
    ```
3. **Run the project locally** to ensure it builds and runs cleanly:
    ```bash
    go run main.go serve
    ```
4. If you've changed APIs or commands, update documentation where appropriate.
5. Ensure tests (if any) pass and run `go fmt ./...` to format your code.
6. Open a pull request!

We’ll review it quickly and provide feedback if needed.

## Contributions

Any contributions you make will be under the **MIT License**.

When you submit code changes, you agree that your contributions will be licensed under the same [MIT License](https://choosealicense.com/licenses/mit/) that covers this project.

If you have questions about licensing, open an issue or contact the maintainers.

## Reporting Bugs

We use [GitHub Issues](https://github.com/6missedcalls/cobra-gin-starter/issues) to track public bugs.
To report a bug, [open a new issue](https://github.com/6missedcalls/cobra-gin-starter/issues/new).

### Writing a Great Bug Report

Clear, detailed bug reports help everyone. Here's what helps most:

- A quick summary and background
- Steps to reproduce (be specific; include commands or code snippets)
- What you expected to happen
- What actually happened
- Notes — possible reasons, environment info, or what you've already tried

## Use a Consistent Coding Style

- Use **2 spaces** for indentation
- Follow Go's standard formatting (`go fmt ./...`)
- Keep HTML and Templ components consistently indented
- Keep Tailwind utility classes readable — for long lists, prefer one per line

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](https://choosealicense.com/licenses/mit/).
