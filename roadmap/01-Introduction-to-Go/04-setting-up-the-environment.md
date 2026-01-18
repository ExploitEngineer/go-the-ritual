# Setting Up the Go Environment

This guide covers installing Go on Linux, macOS, and Windows.

---

## Linux

1. **Remove any previous Go installation**

```bash
sudo rm -rf /usr/local/go
```

1. **Extract the downloaded archive**

```bash
tar -C /usr/local -xzf go1.25.6.linux-amd64.tar.gz
```

> ⚠️ Do not untar into an existing `/usr/local/go` directory. This may break your Go installation.

1. **Add Go to your PATH**

Add the following line to `$HOME/.profile` for a user-specific installation, or `/etc/profile` for system-wide installation:

```bash
export PATH=$PATH:/usr/local/go/bin
```

To apply changes immediately:

```bash
source $HOME/.profile
```

1. **Verify the installation**

```bash
go version
```

Confirm it prints the installed version of Go.

---

## macOS

1. **Install Go using the downloaded package**

- Open the `.pkg` file and follow the prompts.
- The package installs Go to `/usr/local/go`.
- Ensure `/usr/local/go/bin` is in your PATH. Restart Terminal sessions if needed.

1. **Verify the installation**

```bash
go version
```

Confirm it prints the installed version of Go.

---

## Windows

1. **Install Go using the downloaded MSI**

- Open the MSI file and follow the installation prompts.
- By default, Go installs to `C:\Program Files\Go` or `C:\Program Files (x86)\Go`.
- Close and reopen any command prompts for environment changes to take effect.

1. **Verify the installation**

- Open Command Prompt (`cmd`).
- Run:

```cmd
go version
```

Confirm it prints the installed version of Go.
