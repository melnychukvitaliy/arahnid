In order to use rasberry server without password you have to run next command:

    cat .ssh/id_rsa.pub | ssh pi@192.168.7.105  'cat >> .ssh/authorized_keys'

If you don't have id_rsa.pub file you have to run:

    ssh-keygen