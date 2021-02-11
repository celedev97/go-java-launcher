# Go Java Launcher

## Problems?
Have you ever wrote a super cool Java/Kotlin desktop application and packaged it in a **jar** just to find out that most users will not have a JRE installed?

Or have you ever found someone giving you weird bug reports because they have an outdated version of Java while you're developing for Java 11?

If you did, then you just found the solution!

## Solution!
`Go Java Launcher` is a launcher for jar files, that will search for the right JRE version and launch the jar with that version.

Now... what's the point of this if there are more mature project like launch4j?

The point is that `Go Java Launcher` will work even if the JRE is not installed, if it can't find a suitable JRE it will download the one you want from AdoptOpenJDK and it will install it before launching the jar, all pretty fastly and without requiring user interaction.

## Get it

### Release

If you go on releases you will find the builds for windows and linux, currently it is only available for x64 machine, if you want the launcher to support more architectures...

Do a pull request, i'm too lazy to do that ;D

On windows there are a pair of downsides to using the release: 
if you download a release the .exe will have the default icon and no app manifest, feel free to change them with resourcehack to give the launcher the aspect of something that's part of your app.

### Build it

To build this launcher you need to install [Golang](https://golang.org/dl/) and add it to the PATH.

If you build it on windows you will be able to specify icon and manifest file by using [akavel/rsrc](github.com/akavel/rsrc) in this way:

```
rsrc -ico icon.ico
go build -ldflags -H=windowsgui
```

The project already have a dependency on akavel/rsrc so you could install it even by just running `go build` once at first.