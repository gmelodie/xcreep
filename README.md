# xcreep

`xcreep` is an easy wrapper for [`xclip`](https://github.com/astrand/xclip) made in Go.

## Usage

To copy the contents of a file, say `example.txt`, to the clipboard:
```
./xcreep example.txt
```

If you are on a different directory, just use either the absolute or the relative paths to that file:
```
./xcreep /home/root/flag.txt    # absolute path
./xcreep ../root/flag.txt       # relative path (if you are in /home/some-dir)
```

## Why?
Because I was tired of (1) having to look up how to copy to clipboard with xclip and
(2) not being able to `cat file.txt > xclip` or `cat file.txt | xclip` or simply `xclip file.txt`

> But you can just make a simple script to wrap xclip

[I know](https://gist.github.com/gmelodie/0830c03eee0addb44073cc93dc02dd94), but it was fun coding it in Go though
