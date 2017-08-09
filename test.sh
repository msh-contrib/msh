# test

import 'test'

import {
  fs_write,
  fs_read
} from 'fs'

# check whether file is a symlink
main() {
  if [[ -L "$1" ]]; then
    return 0
  fi

  fs_read "./test" err_handler

  return 1
}
