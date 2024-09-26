### Setting up the Root Filesystem (rootfs)

To begin, you'll need a root filesystem (`rootfs`). It is recommended to use a stopped Docker container for this. I exported the `rootfs` from a stopped Alpine Docker container. 

### Creating the Configuration File

Next, create a `container.json` configuration file with the following content:

```json
{
    "id": "test-container",
    "command": {
      "args": ["/bin/bash"]
    },
    "rootfs": "/home/rojin/dev/test",
    "namespaces": ["NEWNET", "NEWIPC", "NEWNS", "NEWPID", "NEWUTS"],
    "capabilities": ["SYS_ADMIN", "SYS_RESOURCE"]
}
```

### Notes:
- **Capabilities**: You don’t need to worry about capabilities for now; feel free to omit or ignore them.
- **Command**: Replace the `command` argument (`"/bin/bash"`) with the desired command or script you'd like to execute.

### Executing the Binary

Once you’ve set up the `container.json`, execute the binary as follows:

```bash
sudo ./ghost-project exec --config=container.json
```
