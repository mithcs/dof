# dof
### A simple and easy to use dot files manager

## Quick Tour

```
# Initialize dof
dof init

# Add files
dof add ~/.config/helix --name dots
dof add ~/.zshrc ~/.zimrc --name dots
dof add ~/dwm ~/st --name suckless

# Deploy files
dof deploy zsh
dof deploy all

# Remove files
dof remove ~/.zshrc --name dots

# List files
dof list
```

## Features:
- Cross Platform
- Supports both (copying and symlinking) method
- Static and self-contained

## Different approaches

### 1. Modular

- Splitting files among multiple repositories.
- Example: dwm and st build in suckless_builds repo and simple config files in dots repo.

### 2. Monolithic

- All files in single repository.
- Example: .zshrc, .zimrc, dwm and st build &em; all in single folder.

> [!NOTE]
> dof does not interact with VCS (git, ...), you have to manage it yourself.

