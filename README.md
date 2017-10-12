# DevOpsEval
ヾ(⌐■_■)ノ♪ Simple Evaluation Task for DevOps / Infrastructure / Cloud Engineers

## Write some automation to do the following:

1. Create a new `t2.micro` Amazon EC2 instance in `us-east-1` using the provided credentials and keypair (named `DevOpsEval`).
2. Upload the tarball found at: $URL
3. Extract the tarball's contents somewhere onto the system (use your best judgment as to where this should live).
4. Configure the service to run on system startup.
5. Reboot the instance.

## Constraints

- You can use any tool/language you wish to script the automation (e.g: Bash, Python, Elixir, or a framework like Chef or Ansible.)
- You can make the service run using any method you like (e.g.: systemd, Supervisor, init).
- You must ensure that the service is publicly reachable on TCP port 80 and responds via a web browser.
- You must ensure that the instance is publicly reachable via SSH using the provided keypair.

## Deliverables

- Tarball of a local Git repository containing your script.  A commit history that shows your progress as you go is strongly preferred.
- The public IP address of the cloud instance with the service up and responding on port 80 that is also reachable via SSH using the provided keypair.  We should be able to visit this IP in our browsers and see the service we sent you running.
