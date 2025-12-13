# dof
### A simple and easy to use dot files manager

## Usage

```
# Initialize dof
dof init

# Add files
dof add ~/.zshrc ~/.zimfw --name zsh

# Deploy files
dof deploy zsh

# Remove files
dof remove ~/.zshrc --name zsh

# List files
dof list
```

## Features:
- Cross Platform 
- Supports both (copying and symlinking) method
- Static and self-contained

> [!NOTE]
> dof does not interact with VCS (git, ...), you have to manage it yourself.

