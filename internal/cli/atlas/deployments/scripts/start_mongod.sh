if [ ! -f $KEYFILE ]
then
    echo $KEYFILECONTENTS > $KEYFILE
    chmod 400 $KEYFILE
fi

python3 /usr/local/bin/docker-entrypoint.py \
        --transitionToAuth \
        --dbpath $DBPATH \
        --keyFile $KEYFILE \
        --replSet $REPLSETNAME \
        --setParameter "mongotHost=$MONGOTHOST" \
        --setParameter "searchIndexManagementHostAndPort=$MONGOTHOST"