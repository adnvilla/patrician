#!/usr/bin/ruby

require 'tempfile'

filename = "./commit_hash.go"

# the commit hash is the starting point of things affecting the build.
commit_hash = `git rev-parse HEAD`.strip
git_describe = `git describe --tags --dirty=-dirty`.strip

# write the commit_hash.go to a temporary...
file = Tempfile.new("commit_hash.go")
file.write(
<<BEGIN
package main

const COMMIT_HASH = "#{commit_hash}"
const GIT_DESCRIBE = "#{git_describe}"
BEGIN
)
file.close()

# only replace commit_hash.go if the contents changed, so that
# Makefile rules can depend on its timestamp.
if not system("diff #{file.path} #{filename} >/dev/null")
then
  system("cp #{file.path} #{filename}")
  puts "Generated #{filename} with COMMIT_HASH=#{commit_hash} and GIT_DESCRIBE=#{git_describe}"
else
  puts "#{filename} unchanged"
end
