if [ ! -f $KEYFILEPATH ]
then
    echo $KEYFILECONTENTS > $KEYFILEPATH
    chmod 400 $KEYFILEPATH
fi

/etc/mongot-localdev/mongot \
                     --data-dir $DATADIR \
                     --mongodHostAndPort $MONGODHOST \
                     --keyFile $KEYFILEPATH