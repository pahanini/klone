# Klone

A tool to clone repositories into a folder structure matching the repository's domain plus username.

For example:

```
git@github.com:pahanini/klone.git to $KLONE_HOME/github.com/pahanini/klone
git@yourdomain.com:user/repo.git to $KLONE_HOME/yourdomain.com/user/repo
```

This tool follows a `domain/username/repo` folder structure, which I use to manage multiple repositories efficiently. 
It eliminates the need for manual directory setup.

```
chmod 600 ~/.ssh/id_rsa
ssh-add ~/.ssh/id_rsa
```
