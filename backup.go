package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "golang.org/x/crypto/ssh"
    "github.com/pkg/sftp"
)

func main() {
    // SSH and SFTP configuration
    sshConfig := &ssh.ClientConfig{
        User: "username", // Replace with your SSH username
        Auth: []ssh.AuthMethod{
            ssh.Password("password"), // Replace with your SSH password or use ssh.PublicKeys for key-based auth
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Replace with proper host key verification
    }

    // Establish SSH connection
    client, err := ssh.Dial("tcp", "remote_host:22", sshConfig) // Replace with your remote host and port
    if err != nil {
        log.Fatalf("Failed to dial: %s", err)
    }
    defer client.Close()

    // Create new SFTP client
    sftpClient, err := sftp.NewClient(client)
    if err != nil {
        log.Fatalf("Failed to create SFTP client: %s", err)
    }
    defer sftpClient.Close()

    // Open the source file
    srcFile, err := os.Open("/path/to/source/file") // Replace with the source file path
    if err != nil {
        log.Fatalf("Failed to open source file: %s", err)
    }
    defer srcFile.Close()

    // Create the destination file
    dstFile, err := sftpClient.Create("/path/to/destination/file") // Replace with the destination file path
    if err != nil {
        log.Fatalf("Failed to create destination file: %s", err)
    }
    defer dstFile.Close()

    // Copy the file
    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        log.Fatalf("Failed to copy file: %s", err)
    }

    fmt.Println("File backed up successfully")
}