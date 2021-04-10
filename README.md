# tds

tds is a simple TUI for disk usage inspired by kdirstat/qdirstat/windirstat.

# Install

```
go get github.com/shric/tds
```

# Usage

```
tds
```

# Current features

- Shows an expandable tree view of files from the current directory
- Shows the disk usage of that file (the entire tree's usage if it's a directory)
- Sorts the files in the current directory by usage, descending

# Planned features

- Specify sort order
- Allow deletion of directories
- Allow refresh (e.g. after you've made a change output the program)
- Show a partial view while gathering stats (currently it shows nothing until all stats gathered)
- Improve formatting (e.g. toggle between size formats)

# No plans to implement

- Flame graphs like qdirstat/kdirstat
- Breakdown by file type
