# [FrameDB Wiki](https://framedb.github.io/wiki)

FrameDB Wiki is a Logseq graph containing information and references
about vintage eyewear, for collectors.

See also: [soyart.github.io/logsex](https://soyart.github.io/logsex)

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

wiki-cli provides tools for some tedious tasks, like renaming asset files.

To use wiki-cli, build the executable first (or use scripts `build-cli.sh`):

```sh
cd ./wiki-cli          # cd into /wiki-cli for go.mod
go build ./cmd/main.go # Build executable /wiki-cli/main
cp main ../cli         # Copy executable to /cli
cd ..
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
