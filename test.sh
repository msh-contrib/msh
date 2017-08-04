# test

name=$test

# check whether file is a symlink
fs.is_symlink() {
  if [[ -L "$1" ]]; then
    return 0
  fi

  return 1
}
