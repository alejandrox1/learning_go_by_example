1. Read the first argument
 * `parent1` rund `/proc/self/exe` which is a special file containing an
   in-memory iage of the current executable. In other words, we run ourselves,
   but passing `child` as a first parameter. So right now we are running a 
   program that allows us to run another user-requested program suplied in `o.Args[2:]` 

2. Adding Namspaces
 * Add this to the second line of the `parent` method to pass extra flags when
   it runs the `child` process
   ``` 
   cmd.SysProcAttr = &syscall.SysProcAttr{
       Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
   }
   ```
   So your programm will now be running inside the **UTS**, **PID**, and
   **MNT** namesapces.

3. Root Filesystem

