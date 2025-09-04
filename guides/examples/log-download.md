# Download and Decompress Logs

You can download and decompress logs by piping the command to gunzip.
To do this you'll need to output to stdout in the following way: 

```bash
atlas logs download <hostname> mongodb.gz --out - | gunzip
```
