# IndexOf
Get the index of an element in a given array.

## Installation and Usage

1. Add the module to the project using the `go get` command.
    ```sh
    go get github.com/leejose/go-indexof
    ```

2. Import.
    ```go
    import (
        utils "github.com/leejose/go-indexof"
    )
    ```
    > NOTE: To get around naming convention conflicts feel free to change `utils` to another name as you see fit. Ex. `arrs`, `u`, `knnw`

3. Example using `knnw`
    ```go
    arr := []int{1,2,3,4,5,6,7,8}
    idx := knnw.IndexOf(3, arr)
    ```
