# Build Your Own Container
Based on [Build Your Own Container Using Less than 100 Lines of Go](https://www.infoq.com/articles/build-a-container-golang)

## Read the first argument
 
`parent1` rund `/proc/self/exe` which is a special file containing an
in-memory iage of the current executable. In other words, we run ourselves,
but passing `child` as a first parameter. So right now we are running a 
program that allows us to run another user-requested program suplied in `os.Args[2:]` 

## Adding Namspaces
Add this to the second line of the `parent` method to pass extra flags when
it runs the `child` process
``` 
cmd.SysProcAttr = &syscall.SysProcAttr{
	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
}
```
So your programm will now be running inside the **UTS**, **PID**, and **MNT** namesapces.

## Root Filesystem
Curretly our process is in an isolated set of namespaces and it looks the same
as the host.
This is because we are in a mount namespace, but the initial mounts are
inherited from the creating namespace.

In order to change this we will use the following to swap into a root
filesystem - place these at the begining of `child()`:
```
must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	must(os.MdkirAll("rootfs/oldrootfs", 0700))
	must(syscal.PivotRoot("rootfs", "rootfs/oldrootfs"))
	must(os.Chdir("/"))
```

The last two lines tell the OS to move the current directory at `/` to
`rootfs/oldrootfs`, and to swap the new rootfs directory to `/`.
After the pivot, the directory `/` in the container will refer to the rootfs.
* The bind mount call is needed to satisfy some requirements of the `PivotRoot`
  command - the OS requires the `PivotRoot` command to be used swap two
  filesystems that are not part of the same tree, which bind mounting the
  rootfs to itself achieves.

## Initialising the World of the Container
Now we have a process running in a set of isolated namesapces with a root
filesystem of our choosing.

What we haven't covered is hw to set up cgroups and how to manage the root
filesystem to efficiently download and cache the root filesystem images we
pivotroot-ed into.

Also, the other namesapce still have their default contents.
We should figure out, for example, how to set up networking, swap to the
correct uid before running the processes and so on.


