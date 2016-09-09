package s4a

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func ReadJsonInput(object interface{}) error {
    return ReadFJsonInput(os.Stdin, object)
}

func ReadFJsonInput(r io.Reader, object interface{}) error {
    input, err := ioutil.ReadAll(r)
    if err != nil {
        return err
    }

    err = json.Unmarshal(input, object)
    if err != nil {
        return err
    }

    return nil
}

func WriteJsonOutput(object interface{}) error {
    return WriteFJsonOutput(os.Stdout, object)
}

func WriteFJsonOutput(w io.Writer, object interface{}) error {
    output, err := json.Marshal(object)
    if err != nil {
        return err
    }
    fmt.Fprint(w, string(output))

    return nil
}
