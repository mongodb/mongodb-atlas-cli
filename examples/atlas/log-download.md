# Download and Unzip Logs

You can download and unzip logs in one go like

```bash
mongocli atlas logs download <processName> mongodb.gz \
    --out /dev/stdout \
    --force | gunzip
```
