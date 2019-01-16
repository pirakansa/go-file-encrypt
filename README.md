# file-encrypt

## install

```
go get -u github.com/pirakansa/go-file-encrypt
```

## file-encrypt cmd

```
Usage of file-encrypt
  --decode
        do decode
  --encode
        do encode
  --if string
        input file
  --of string
        output file
  --kf string
    	  key file
```

## key file Sample

``` json
{
    "Encrypt": {
        "Algorithm": "AES-256-CFB",
        "Key": "=============256bit=============",
        "InitialVector": "=====128bit====="
    }
}

```
