# [FrameDB Wiki](https://framedb.github.io/wiki)

FrameDB Wiki is a Logseq graph containing information and references
about vintage eyewear, for collectors.

It is based on [logsex](https://github.com/soyart/logsex),
and will regularly be rebased on top of logsex. Branch master
is actually just logsex.

## Scripts and Git hooks

logsex provides a set of shell scripts to work with Logseq documents.

For example, `remove-collapsed.sh` quickly remove occurrences of `collapsed::`
tag in Markdown documents. This prevents needless changes from our togglings.

We can also enable a pre-commit Git hook that will run `remove-collasped.sh`
on all files easily with:

```shell
git config --local core.hooksPath .githooks/
```

## wiki-cli

FrameDB Wiki also provides a simple Go CLI program `wiki-cli`.

To use wiki-cli, build the executable first (or use scripts `build-cli.sh`):

```sh
cd ./wiki-cli      # cd into /wiki-cli for go.mod
go build ./main.go # Build executable
cp main ../cli     # Copy executable to /cli
cd ..              # cd back into previous pwd
```

Then you can just run it to do stuff:

```sh
# Rename an asset as well as its references with subcommand `rename-asset`
./cli rename 'some_old_name.png' 'new_name.png'

# Or renames from a JSON map file
./cli renames 'some_replacements.json'

# Clean all unreferenced assets
./cli cleanup
```
