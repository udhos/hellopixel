# hellopixel

## Build on Windows

1. Install Git Bash from https://gitforwindows.org/

2. Install mingw64

- Download x86_64-8.1.0-release-posix-sjlj-rt_v6-rev0.7z from:

https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/8.1.0/threads-posix/sjlj/x86_64-8.1.0-release-posix-sjlj-rt_v6-rev0.7z

- Extract it to c:\mingw64

- Add c:\mingw64\bin to %PATH%

3. Open Git Bash and build as usual

Recipe:

    git clone https://github.com/udhos/hellopixel
    cd hellopixel
    go install ./hellopixel

4. Run 'hellopixel'

You can run it from anywhere (Git Bash, CMD prompt, etc)

    hellopixel

