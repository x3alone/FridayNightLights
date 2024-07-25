# Docker Commands Reference

- **Build an image**: 
    ```sh
    docker build -t <image-name> .
    ```

- **Run a container**: 
    ```sh
    docker run -p <host-port>:<container-port> --name <container-name> <image-name>
    ```

- **Stop a container**: 
    ```sh
    docker stop <container-name>
    ```

- **Remove a container**: 
    ```sh
    docker rm <container-name>
    ```

- **List all containers**: 
    ```sh
    docker ps
    ```

- **List all images**: 
    ```sh
    docker images
    ```
